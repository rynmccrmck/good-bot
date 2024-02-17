package goodbot_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	goodbot "github.com/rynmccmrmck/goodbot"
	"github.com/rynmccmrmck/goodbot/mocks"
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

	mockNetworkUtils.EXPECT().GetDomainName("66.249.66.1").Return("crawl-66-249-66-1.googlebot.com").AnyTimes()
	mockNetworkUtils.EXPECT().GetDomainName("127.0.0.1").Return("localhost").AnyTimes()
	mockNetworkUtils.EXPECT().GetASN("66.249.66.2").Return("32934", nil).AnyTimes()
	mockNetworkUtils.EXPECT().GetASN("66.249.66.3").Return("12345", nil).AnyTimes()

	tests := []struct {
		name            string
		userAgent       string
		ipAddress       string
		expectedIsBot   bool
		expectedBotName string
	}{
		{
			name:            "Googlebot",
			userAgent:       "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
			ipAddress:       "66.249.66.1",
			expectedIsBot:   true,
			expectedBotName: "Googlebot",
		},
		{
			name:            "Googlebot Wrong IP",
			userAgent:       "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
			ipAddress:       "127.0.0.1",
			expectedIsBot:   false,
			expectedBotName: "",
		},
		{
			name:            "Facebook external hit",
			userAgent:       "facebookexternalhit/2.0",
			ipAddress:       "66.249.66.2",
			expectedIsBot:   true,
			expectedBotName: "Facebook external hit",
		},
		{
			name:            "FacebookBot UA Wrong ASN",
			userAgent:       "facebookexternalhit/2.0",
			ipAddress:       "66.249.66.3",
			expectedIsBot:   false,
			expectedBotName: "",
		},
		{
			name:            "Unknown UA and IP",
			userAgent:       "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N)",
			ipAddress:       "192.168.1.1",
			expectedIsBot:   false,
			expectedBotName: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			botResult := botService.IsGoodBot(tc.userAgent, tc.ipAddress)

			if botResult.IsGoodBot != tc.expectedIsBot || botResult.BotName != tc.expectedBotName {
				t.Errorf("CheckBotIdentity(%q, %q) = (%v, %q), want (%v, %q)",
					tc.userAgent, tc.ipAddress, botResult.IsGoodBot, botResult.BotName, tc.expectedIsBot, tc.expectedBotName)
			}
		})
	}
}
