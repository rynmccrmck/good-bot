// Package goodbot provides utilities for network operations such as domain name resolution,
// ASN lookup, and bot detection mechanisms based on various criteria including IP verification,
// User-Agent matching, and more. It utilizes external libraries for enhanced functionality like
// IP to ASN mapping and CIDR checks.
package goodbot

import (
	"context"
	"net"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/ammario/ipisp/v2"
	internal "github.com/rynmccrmck/good-bot/internal"
	cidr "github.com/yl2chen/cidranger"
)

// NetworkUtils defines an interface for network-related utilities including
// domain name resolution and ASN lookup for a given IP address.
type NetworkUtils interface {
	GetDomainName(ipAddress string) string
	GetASN(ipAddress string) (string, error)
}

// defaultNetworkUtils implements the NetworkUtils interface with basic network
// utilities.
type defaultNetworkUtils struct{}

// GetDomainName resolves the domain name for the given IP address. It returns
// the first hostname found or a default message if no domain name is found.
func (n defaultNetworkUtils) GetDomainName(ipAddress string) string {
	hosts, err := net.LookupAddr(ipAddress)
	if err != nil {
		return "No domain name found"
	}
	return hosts[0]
}

// GetASN looks up the ASN information for the given IP address using the
// iptoasn external library. It returns the ASN number as a string.
func (n defaultNetworkUtils) GetASN(ipAddress string) (string, error) {
	resp, err := ipisp.LookupIP(context.Background(), net.ParseIP(ipAddress))
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(resp.ASN)), nil
}

// BotService provides methods for bot detection using network utilities.
type BotService struct {
	networkUtils NetworkUtils
}

// NewBotService creates a new instance of BotService with the provided
// NetworkUtils implementation.
func NewBotService(nu NetworkUtils) *BotService {
	return &BotService{
		networkUtils: nu,
	}
}

// isVerifiedIP checks if the given IP matches the domain's resolved IPs.
func isVerifiedIP(nu NetworkUtils, ip string, sources []string, method string) bool {
	switch method {
	case "dnsReverseForward":
		hostname := nu.GetDomainName(ip)
		for _, source := range sources {
			if strings.HasSuffix(hostname, source) {
				return true
			}
		}
	case "uaCidrMatch":
		ranger := cidr.NewPCTrieRanger()
		for _, source := range sources {
			_, network, _ := net.ParseCIDR(source)
			ranger.Insert(cidr.NewBasicRangerEntry(*network))
		}
		ipAddress := net.ParseIP(ip)
		contains, _ := ranger.Contains(ipAddress)
		if contains {
			return true
		}
	case "uaAsnMatch":
		asn, _ := nu.GetASN(ip)
		for _, source := range sources {
			if asn == source {
				return true
			}
		}
	}
	return false
}

// isUserAgentMatch checks if the user agent matches the pattern.
func IsUserAgentMatch(userAgent, uaPattern string) bool {
	caseInsensitivePattern := "(?i)" + uaPattern
	matched, err := regexp.MatchString(caseInsensitivePattern, userAgent)
	if err != nil {
		return false
	}
	return matched
}

type BotStatus int

const (
	BotStatusUnknown  BotStatus = iota // Bot is not recognized
	BotStatusFriendly                  // Bot is recognized as friendly
)

type BotCheckResult struct {
	BotStatus BotStatus
	BotName   string
}

// CheckBotStatus determines the status of a bot based on the given user agent
// and IP address. It utilizes internal and external checks to classify bots.
func (bs *BotService) CheckBotStatus(ctx context.Context, userAgent, ipAddress string) (BotCheckResult, error) {
	botsData := internal.GetBots().Bots
	resultChan := make(chan BotCheckResult)
	errChan := make(chan error)
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, bot := range botsData {
		wg.Add(1)
		go func(bot internal.BotEntry) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				if IsUserAgentMatch(userAgent, bot.UserAgentPattern) {
					if isVerifiedIP(bs.networkUtils, ipAddress, bot.ValidDomains, bot.Method) {
						select {
						case resultChan <- BotCheckResult{BotStatusFriendly, bot.Name}:
						case <-ctx.Done():
						}
						cancel()
						return
					}
				}
			}
		}(bot)
	}

	go func() {
		wg.Wait()
		close(resultChan)
		close(errChan)
	}()

	for {
		select {
		case result := <-resultChan:
			if result.BotStatus == BotStatusFriendly {
				return result, nil
			}
		case <-ctx.Done():
			break
		}
	}

	return BotCheckResult{BotStatusUnknown, ""}, nil
}

var defaultService = NewBotService(defaultNetworkUtils{})

// CheckBotStatus is a convenience function that uses a default BotService
// instance to check the bot status for the given user agent and IP address.
func CheckBotStatus(userAgent, ipAddress string) (BotCheckResult, error) {
	ctx := context.Background()
	res, err := defaultService.CheckBotStatus(ctx, userAgent, ipAddress)
	return res, err
}
