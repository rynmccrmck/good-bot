package goodbot

import (
	"testing"
)

// Mock data and helper functions for testing
var mockBotsData = []map[string]interface{}{
	{
		"name":             "TestBot",
		"UserAgentPattern": "^TestBot",
		"Method":           "dnsReverseForward",
		"ValidDomains":     []string{"test.domain.com"},
	},
	// Add more mock bot data as needed for testing
}

// Test getDomainName function
func TestGetDomainName(t *testing.T) {
	ipAddress := "8.8.8.8"
	expectedHostname := "dns.google."

	hostname := getDomainName(ipAddress)
	if hostname != expectedHostname {
		t.Errorf("Expected hostname %s, got %s", expectedHostname, hostname)
	}

	// Add more test cases as needed
}

// Test IsUserAgentMatch function
func TestIsUserAgentMatch(t *testing.T) {
	userAgent := "TestBot/1.0"
	uaPattern := "^TestBot"

	if !IsUserAgentMatch(userAgent, uaPattern) {
		t.Errorf("User agent %s should match pattern %s", userAgent, uaPattern)
	}

	// Test a negative case
	userAgent = "AnotherBot/1.0"
	if IsUserAgentMatch(userAgent, uaPattern) {
		t.Errorf("User agent %s should not match pattern %s", userAgent, uaPattern)
	}
}

// For testing functions like getASN and isVerifiedIP, you might consider interfaces for network-related functions to easily mock them.

// Since getASN and isVerifiedIP involve external network calls, consider abstracting these calls behind interfaces and inject mock implementations for testing.
