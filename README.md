# Good Bot

<img src="https://github.com/rynmccrmck/good-bot/assets/5178938/ece7cfe2-4369-4da6-a440-915cbc48368c" alt="drawing" width="200"/>

Good Bot is an open-source, multi-language library designed to enhance web application security and user experience by accurately identifying beneficial automated agents, commonly known as "good bots". These agents perform various essential tasks, including search engine indexing, social media insights, link previews, and more. With bot traffic constituting a significant portion of online interactions, distinguishing between helpful and potentially harmful bots is vital for maintaining performance, security, and accurate analytics. Good Bot provides developers and system administrators with a reliable toolset for this purpose, across multiple programming languages.

### Features

- **Multi-Language Support**: Available for Python, GoLang, and JavaScript, making Good Bot adaptable to a wide range of projects and development environments.
- **Accurate Detection**: Leverages an extensive database of user-agent strings, IP addresses, and DNS verification methods to recognize good bots with high accuracy.
- **Extensibility and Customization**: Designed for easy updates to the bot database and allows custom configurations to meet diverse application needs.

### Getting Started

Select the appropriate installation instructions for your development environment:


#### Basic Usage

**Python Example:**

```bash
pip install good-bot
```

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

**Refer to the documentation for GoLang and JavaScript usage examples.**

- [Go](./languages/go/README.md)
- [Python](./languages/python/README.md)
- [JavaScript](./languages/javascript/README.md) - Coming soon!
- [Ruby](./languages/ruby/README.md) - Coming soon!
- [Java](./languages/java/README.md) - Coming soon!

### How It Works

Good Bot analyzes HTTP request headers for signatures known to belong to good bots, including user-agent strings and IP addresses, and optionally employs DNS reverse lookup for additional verification. This multifaceted approach ensures a balance between accuracy and efficiency.

### Configuration

Good Bot allows extensive configuration options to tailor bot detection capabilities to your specific requirements, including method adjustments, bot data source specifications, and performance optimization settings.

### Contributing

We welcome contributions to Good Bot! Whether it's through adding support for additional languages, enhancing existing features, submitting bug reports, or improving documentation, your help is invaluable in making Good Bot more effective for everyone.

### License

Good Bot is released under the MIT License. See the LICENSE file for more information.

### Support

For support, please open an issue on our GitHub repository. Premium support options are available; contact us directly for more details.
