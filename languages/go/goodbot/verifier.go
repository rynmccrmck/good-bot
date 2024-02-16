package goodbot

import (
	"net"
	"regexp"
	"strconv"

	iptoasn "github.com/jamesog/iptoasn"
	cidr "github.com/yl2chen/cidranger"
)

// getDomainName attempts to find the domain name associated with an IP address.
func getDomainName(ipAddress string) string {
	hosts, err := net.LookupAddr(ipAddress)
	if err != nil {
		return "No domain name found"
	}
	return hosts[0]
}

// getASN returns the Autonomous System Number (ASN) of the given IP address.
func getASN(ipAddress string) (string, error) {
	ip, err := iptoasn.LookupIP(ipAddress)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(ip.ASNum)), nil
}

// isVerifiedIP checks if the given IP matches the domain's resolved IPs.
func isVerifiedIP(ip string, sources []string, method string) bool {
	switch method {
	case "dnsReverseForward":
		hostname := getDomainName(ip)
		for _, source := range sources {
			if hostname == source {
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
		asn, _ := getASN(ip)
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
	matched, err := regexp.MatchString(uaPattern, userAgent)
	if err != nil {
		return false
	}
	return matched
}

func IsGoodBot(userAgent, ipAddress string, botsData []map[string]interface{}) (bool, string) {
	for _, bot := range botsData {
		uaPattern := bot["UserAgentPattern"].(string)
		if IsUserAgentMatch(userAgent, uaPattern) {
			sources := bot["Valid domains"].([]string)
			method := bot["Method"].(string)
			if isVerifiedIP(ipAddress, sources, method) {
				return true, bot["name"].(string)
			}
		}
	}
	return false, ""
}
