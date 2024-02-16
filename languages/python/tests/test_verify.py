from unittest.mock import patch
from goodbot import is_good_bot, is_verified_ip, is_user_agent_match, get_domain_name

bots_data = [
    {
        "name": "TestBot",
        "User Agent": "TestBot",
        "Method": "dnsReverseForward",
        "Valid domains": ["test.domain.com"]
    }
]


def test_get_domain_name():
    assert get_domain_name("8.8.8.8") == "dns.google"
    assert get_domain_name("127.0.0.1") == "No domain name found"


def test_is_user_agent_match():
    assert is_user_agent_match("TestBot", None, "TestBot") is True
    assert is_user_agent_match("OtherBot", None, "TestBot") is False
    assert is_user_agent_match("TestBot/1.0", None, r"TestBot\/\d\.\d") is True


@patch('socket.gethostbyaddr')
def test_get_domain_name_with_mock(mock_gethost):
    mock_gethost.return_value = ("mocked.domain.com", [], ["8.8.8.8"])
    assert get_domain_name("8.8.8.8") == "mocked.domain.com"


@patch('goodbot.get_domain_name')
def test_is_verified_ip_dns(mock_get_domain_name):
    mock_get_domain_name.return_value = "test.domain.com"
    assert is_verified_ip("192.0.2.1", ["test.domain.com"], method='dnsReverseForward') is True
    mock_get_domain_name.assert_called_once_with("192.0.2.1")

    mock_get_domain_name.return_value = "No domain name found"
    assert is_verified_ip("192.0.2.2", ["test.domain.com"], method='dnsReverseForward') is False


def test_is_verified_ip_uaCidrMatch():
    assert is_verified_ip("192.0.2.1", ["192.0.2.0/24"], method='uaCidrMatch') is True
    assert is_verified_ip("192.0.3.1", ["192.0.2.0/24"], method='uaCidrMatch') is False
    

@patch('goodbot.is_verified_ip')
@patch('goodbot.is_user_agent_match')
def test_is_good_bot(mock_is_user_agent_match, mock_is_verified_ip):
    mock_is_user_agent_match.return_value = True
    mock_is_verified_ip.return_value = True

    is_bot, bot_name = is_good_bot("TestBot", "192.0.2.1", bots_data)
    assert is_bot is True
    assert bot_name == "TestBot"

    mock_is_user_agent_match.assert_called_once_with("TestBot", None, "TestBot")
    mock_is_verified_ip.assert_called_once_with("192.0.2.1", ["test.domain.com"], 'dnsReverseForward')

    # Test for a negative case
    mock_is_user_agent_match.return_value = False
    is_bot, _ = is_good_bot("BadBot", "192.0.2.1", bots_data)
    assert is_bot is False
