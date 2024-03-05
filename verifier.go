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
	GetHosts(ipAddress string) []string
	GetASN(ipAddress string) (string, error)
	DoesHostnameResolveBackToIP(ipAddress, hostname string) bool
}

// defaultNetworkUtils implements the NetworkUtils interface with basic network
// utilities.
type defaultNetworkUtils struct{}

// GetDomainName resolves the domain name for the given IP address. It returns
// the first hostname found or a default message if no domain name is found.
func (n defaultNetworkUtils) GetHosts(ipAddress string) []string {
	hosts, err := net.LookupAddr(ipAddress)
	if err != nil {
		return nil
	}
	return hosts
}

func (n defaultNetworkUtils) DoesHostnameResolveBackToIP(ipAddress, hostname string) bool {
	ips, err := net.LookupIP(hostname)
	if err != nil {
		return false
	}
	for _, ip := range ips {
		if ip.String() == ipAddress {
			return true
		}
	}
	return false
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
func isVerifiedIP(nu NetworkUtils, ip string, sources []string, method string) (bool, BotStatus) {
	switch method {

	case "uaOnly":
		return true, BotStatusPotentiallyFriendly
	case "dnsReverseForward":
		hosts := nu.GetHosts(ip)
		for _, host := range hosts {
			for _, source := range sources {
				if strings.HasSuffix(host, source) {
					if nu.DoesHostnameResolveBackToIP(ip, host) {
						return true, BotStatusFriendly
					} else {
						return false, BotStatusPotentialImposter
					}
				}
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
			return true, BotStatusFriendly
		}
	case "uaAsnMatch":
		asn, _ := nu.GetASN(ip)
		for _, source := range sources {
			if asn == source {
				return true, BotStatusFriendly
			}
		}
	}
	return false, BotStatusUnknown
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
	BotStatusUnknown             BotStatus = iota // Bot is not recognized
	BotStatusFriendly                             // Bot is recognized as friendly
	BotStatusPotentiallyFriendly                  // Bot is recognized as potentially friendly
	BotStatusPotentialImposter                    // Bot is recognized as potentially unfriendly
	BotStatusUnfriendly                           // Bot is recognized as unfriendly
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
					if ok, status := isVerifiedIP(bs.networkUtils, ipAddress, bot.ValidDomains, bot.Method); ok {
						select {
						case resultChan <- BotCheckResult{status, bot.Name}:
						case <-ctx.Done():
						}
					} else {
						select {
						case resultChan <- BotCheckResult{BotStatusPotentialImposter, bot.Name}:
						case <-ctx.Done():
						}
					}
					cancel()
					return
				}
			}
		}(bot)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for {
		select {
		case result, ok := <-resultChan:
			if !ok {
				return BotCheckResult{BotStatusUnknown, ""}, nil
			}
			return result, nil
		case <-ctx.Done():
			return BotCheckResult{BotStatusUnknown, ""}, nil
		}
	}
}

var defaultService = NewBotService(defaultNetworkUtils{})

// CheckBotStatus is a convenience function that uses a default BotService
// instance to check the bot status for the given user agent and IP address.
func CheckBotStatus(userAgent, ipAddress string) (BotCheckResult, error) {
	ctx := context.Background()
	res, err := defaultService.CheckBotStatus(ctx, userAgent, ipAddress)
	return res, err
}
