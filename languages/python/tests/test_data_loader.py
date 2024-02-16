import pytest
from goodbot.data_loader import load_bot_data

mock_json_data = """
{
    "bots": [
        {
            "name": "MockBot",
            "User Agent": "MockBot",
            "Method": "dnsReverseForward",
            "Valid domains": ["mock.domain.com"],
            "Sources": ["https://mock.source.com"]
        }
    ]
}
"""

def test_load_bot_data(monkeypatch):
    def mock_open(*args, **kwargs):
        return mock_json_data

    monkeypatch.setattr("builtins.open", lambda *args, **kwargs: mock_open())

    data = load_bot_data()
    assert data['bots'][0]['name'] == 'MockBot'
    assert data['bots'][0]['User Agent'] == 'MockBot'
    assert data['bots'][0]['Method'] == 'dnsReverseForward'
    assert data['bots'][0]['Valid domains'] == ['mock.domain.com']
    assert data['bots'][0]['Sources'] == ['https://mock.source.com']
