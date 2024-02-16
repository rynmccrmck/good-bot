#### Python

```bash
pip install good-bot
```

#### Basic Usage

**Python Example:**

```python
from goodbot import is_good_bot
bots_data = load_bot_data()

user_agent = "<USER_AGENT_STRING>"
ip_address = "<IP_ADDRESS>"

is_bot, bot_name = is_good_bot(user_agent, ip_address, bots_data)
if is_bot:
    print(f"Good bot detected: {bot_name}")
else:
    print("No good bot detected.")
```