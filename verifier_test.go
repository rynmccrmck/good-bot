package goodbot_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	goodbot "github.com/rynmccrmck/good-bot"
	"github.com/rynmccrmck/good-bot/mocks"
)

func TestIsUserAgentMatch(t *testing.T) {
	userAgent := "TestBot/1.0"
	uaPattern := "^TestBot"

	if !goodbot.IsUserAgentMatch(userAgent, uaPattern) {
		t.Errorf("User agent %s should match pattern %s", userAgent, uaPattern)
	}

	// Test a negative case
	userAgent = "AnotherBot/1.0"
	if goodbot.IsUserAgentMatch(userAgent, uaPattern) {
		t.Errorf("User agent %s should not match pattern %s", userAgent, uaPattern)
	}
}

func TestCheckBotIdentity(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockNetworkUtils := mocks.NewMockNetworkUtils(mockCtrl)
	botService := goodbot.NewBotService(mockNetworkUtils)

	mockNetworkUtils.EXPECT().GetHosts("66.249.66.1").Return([]string{"crawl-66-249-66-1.googlebot.com"}).AnyTimes()
	mockNetworkUtils.EXPECT().DoesHostnameResolveBackToIP("66.249.66.1", "crawl-66-249-66-1.googlebot.com").Return(true).AnyTimes()
	mockNetworkUtils.EXPECT().GetHosts("127.0.0.1").Return([]string{"localhost"}).AnyTimes()
	mockNetworkUtils.EXPECT().DoesHostnameResolveBackToIP("127.0.0.1", "localhost").Return(true).AnyTimes()
	mockNetworkUtils.EXPECT().GetASN("66.249.66.2").Return("32934", nil).AnyTimes()
	mockNetworkUtils.EXPECT().GetASN("66.249.66.3").Return("12345", nil).AnyTimes()

	mockNetworkUtils.EXPECT().GetHosts("66.249.66.4").Return([]string{"crawl-66-249-66-1.googlebot.com"}).AnyTimes()
	mockNetworkUtils.EXPECT().DoesHostnameResolveBackToIP("66.249.66.4", "crawl-66-249-66-1.googlebot.com").Return(false).AnyTimes()

	tests := []struct {
		name              string
		userAgent         string
		ipAddress         string
		expectedBotStatus goodbot.BotStatus
		expectedBotName   string
	}{
		{
			name:              "Googlebot",
			userAgent:         "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
			ipAddress:         "66.249.66.1",
			expectedBotStatus: goodbot.BotStatusFriendly,
			expectedBotName:   "Googlebot",
		},
		{
			name:              "Googlebot Wrong IP",
			userAgent:         "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
			ipAddress:         "127.0.0.1",
			expectedBotStatus: goodbot.BotStatusPotentialImposter,
			expectedBotName:   "Googlebot",
		},
		{
			name:              "Googlebot Spoofed reverse DNS",
			userAgent:         "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
			ipAddress:         "66.249.66.4",
			expectedBotStatus: goodbot.BotStatusPotentialImposter,
			expectedBotName:   "Googlebot",
		},
		{
			name:              "Not A known Useragent",
			userAgent:         "Mozilla/5.0 (compatible; nothing-weve-seen-before)",
			ipAddress:         "127.0.0.1",
			expectedBotStatus: goodbot.BotStatusUnknown,
			expectedBotName:   "",
		},
		{
			name:              "Facebook external hit",
			userAgent:         "facebookexternalhit/1.1 (+http://www.facebook.com/externalhit_uatext.php)",
			ipAddress:         "66.249.66.2",
			expectedBotStatus: goodbot.BotStatusFriendly,
			expectedBotName:   "Facebook external hit",
		},
		{
			name:              "FacebookBot UA Wrong ASN",
			userAgent:         "facebookexternalhit/2.0",
			ipAddress:         "66.249.66.3",
			expectedBotStatus: goodbot.BotStatusPotentialImposter,
			expectedBotName:   "Facebook external hit",
		},
		{
			name:              "Unknown UA and IP",
			userAgent:         "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N)",
			ipAddress:         "192.168.1.1",
			expectedBotStatus: goodbot.BotStatusUnknown,
			expectedBotName:   "",
		},
		{
			name:              "Grapeshot CIDR match",
			userAgent:         "Mozilla/5.0 (compatible; GrapeshotCrawler/2.0; +http://www.grapeshot.co.uk/crawler.php)",
			ipAddress:         "132.145.9.5",
			expectedBotStatus: goodbot.BotStatusFriendly,
			expectedBotName:   "Grapeshot",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			botResult, _ := botService.CheckBotStatus(ctx, tc.userAgent, tc.ipAddress)

			if botResult.BotStatus != tc.expectedBotStatus || botResult.BotName != tc.expectedBotName {
				t.Errorf("CheckBotIdentity(%q, %q) = (%v, %q), want (%v, %q)",
					tc.userAgent, tc.ipAddress, botResult.BotStatus, botResult.BotName, tc.expectedBotStatus, tc.expectedBotName)
			}
		})
	}
}

func TestAdhocTest(t *testing.T) {
	result, _ := goodbot.CheckBotStatus("facebookexternalhit/1.1 (+http://www.facebook.com/externalhit_uatext.php)",
		"173.252.127.15")
	// fmt.Printf("Is Good Bot: %v, Bot Name: %s\n", result.BotName, result.BotStatus)
	// Output: Is Good Bot: true, Bot Name: Facebook external hit
	if result.BotStatus != goodbot.BotStatusFriendly || result.BotName != "Facebook external hit" {
		t.Errorf("Problem with Facebook external hit %v", result.BotName)
	}

}
