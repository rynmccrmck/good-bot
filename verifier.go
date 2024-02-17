package goodbot

import (
	"net"
	"regexp"
	"strconv"
	"strings"

	iptoasn "github.com/jamesog/iptoasn"
	internal "github.com/rynmccrmck/goodbot/internal"
	cidr "github.com/yl2chen/cidranger"
)

type NetworkUtils interface {
	GetDomainName(ipAddress string) string
	GetASN(ipAddress string) (string, error)
}

type defaultNetworkUtils struct{}

func (n defaultNetworkUtils) GetDomainName(ipAddress string) string {
	hosts, err := net.LookupAddr(ipAddress)
	if err != nil {
		return "No domain name found"
	}
	return hosts[0]
}

func (n defaultNetworkUtils) GetASN(ipAddress string) (string, error) {
	ip, err := iptoasn.LookupIP(ipAddress)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(ip.ASNum)), nil
}

type BotService struct {
	networkUtils NetworkUtils
}

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

func (bs *BotService) CheckBotStatus(userAgent, ipAddress string) BotCheckResult {
	botsData := internal.GetBots().Bots
	for _, bot := range botsData {

		uaPattern := bot.UserAgentPattern
		if IsUserAgentMatch(userAgent, uaPattern) {
			sources := bot.ValidDomains
			method := bot.Method
			if isVerifiedIP(bs.networkUtils, ipAddress, sources, method) {
				return BotCheckResult{BotStatusFriendly, bot.Name}
			}
		}
	}
	return BotCheckResult{BotStatusUnknown, ""}
}

var defaultService = NewBotService(defaultNetworkUtils{})

func CheckBotStatus(userAgent, ipAddress string) BotCheckResult {
	return defaultService.CheckBotStatus(userAgent, ipAddress)
}
