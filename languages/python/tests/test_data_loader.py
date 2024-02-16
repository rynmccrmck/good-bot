from unittest.mock import patch, mock_open
from goodbot.data_loader import load_bot_data

mock_json_data = """
{
    "bots": [
        {
            "name": "MockBot",
            "UserAgent": "MockBot",
            "Method": "dnsReverseForward",
            "ValidDomains": ["mock.domain.com"],
            "Sources": ["https://mock.source.com"]
        }
    ]
}
"""

@patch("builtins.open", new_callable=mock_open, read_data=mock_json_data)
def test_load_bot_data(monkeypatch):

    data = load_bot_data()
    assert data['bots'][0]['name'] == 'MockBot'
    assert data['bots'][0]['UserAgent'] == 'MockBot'
    assert data['bots'][0]['Method'] == 'dnsReverseForward'
    assert data['bots'][0]['ValidDomains'] == ['mock.domain.com']
    assert data['bots'][0]['Sources'] == ['https://mock.source.com']
