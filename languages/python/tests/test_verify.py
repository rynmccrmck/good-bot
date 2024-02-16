from unittest.mock import patch
from goodbot.verify import (is_good_bot, is_verified_ip, is_user_agent_match, 
    get_domain_name)




def test_get_domain_name():
    assert get_domain_name("127.0.0.1") == "localhost"


def test_is_user_agent_match():
    assert is_user_agent_match("TestBot v0.1.1", "TestBot") is True
    assert is_user_agent_match("OtherBot v0.8", "TestBot") is False
    assert is_user_agent_match("TestBot/1.0", r"TestBot\/\d\.\d") is True


@patch('socket.gethostbyaddr')
def test_get_domain_name_with_mock(mock_gethost):
    mock_gethost.return_value = ("mocked.domain.com", [], ["8.8.8.8"])
    assert get_domain_name("8.8.8.8") == "mocked.domain.com"


@patch('goodbot.verify.get_domain_name')
def test_is_verified_ip_dns(mock_get_domain_name):
    mock_get_domain_name.return_value = "test.domain.com"
    assert is_verified_ip("192.0.2.1", ["test.domain.com"], method='dnsReverseForward') is True
    mock_get_domain_name.assert_called_once_with("192.0.2.1")

    mock_get_domain_name.return_value = "No domain name found"
    assert is_verified_ip("192.0.2.2", ["test.domain.com"], method='dnsReverseForward') is False


def test_is_verified_ip_uaCidrMatch():
    assert is_verified_ip("192.0.2.1", ["192.0.2.0/24"], method='uaCidrMatch') is True
    assert is_verified_ip("192.0.3.1", ["192.0.2.0/24"], method='uaCidrMatch') is False
    
bots_data = [
    {
        "name": "TestBot",
        "UserAgentPattern": "TestBot",
        "Method": "dnsReverseForward",
        "ValidDomains": ["test.domain.com"]
    }
]

@patch('goodbot.verify.is_verified_ip')
@patch('goodbot.verify.is_user_agent_match')
def test_is_good_bot(mock_is_user_agent_match, mock_is_verified_ip):
    mock_is_user_agent_match.return_value = True
    mock_is_verified_ip.return_value = True

    is_bot, bot_name = is_good_bot("TestBot v0.1.1", "192.0.2.1", bots_data)
    assert is_bot is True
    assert bot_name == "TestBot"

    mock_is_user_agent_match.assert_called_once_with("TestBot v0.1.1", "TestBot")
    mock_is_verified_ip.assert_called_once_with("192.0.2.1", ["test.domain.com"], 'dnsReverseForward')

    # Test for a negative case
    mock_is_user_agent_match.return_value = False
    is_bot, _ = is_good_bot("BadBot", "192.0.2.1", bots_data)
    assert is_bot is False
