# Good Bot

[![codecov](https://codecov.io/gh/rynmccrmck/good-bot/graph/badge.svg?token=4X6S00774G)](https://codecov.io/gh/rynmccrmck/good-bot)

[![Known Vulnerabilities](https://snyk.io/test/github/rynmccrmck/good-bot/badge.svg)](https://snyk.io/test/github/rynmccrmck/good-bot)

<img src="https://github.com/rynmccrmck/good-bot/assets/5178938/ece7cfe2-4369-4da6-a440-915cbc48368c" alt="drawing" width="200"/>

**Good Bot** is an open-source Go library designed to enhance web application security and user experience by distinguishing beneficial automated agents, or "good bots", from potentially harmful traffic. In the digital ecosystem where bots play a crucial role—from search engine indexing to social media insights and link previews—it's essential to identify and welcome these friendly bots. Good Bot equips Go developers with the tools to recognize these agents accurately, ensuring your analytics remain accurate and your services optimized.

### Features

- **Accurate Bot Recognition**: Utilizes a comprehensive database of user-agent strings, IP addresses, DNS verification methods, and more to identify good bots with high precision.
- **High Performance**: With embedded data, Good Bot starts quickly and operates efficiently, requiring no external dependencies.
- **Flexibility and Customization**: The bot database is easily extendable and supports customization to align with various application requirements. Contributions are highly encouraged!

### Getting Started

To use Good Bot in your Go project, simply add it as a dependency:

```bash
go get github.com/rynmccmrmck/goodbot
```

#### Basic Usage

Here's a quick example of how to use Good Bot to detect whether a request comes from a known good bot:

```go
package main

import (
    "fmt"
    "github.com/rynmccmrmck/goodbot"
)

func main() {
    userAgent := "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"
    ipAddress := "66.249.66.1"

    result := goodbot.CheckBotStatus(userAgent, ipAddress)
    if result.Status == goodbot.BotStatusFriendly {
        fmt.Printf("Friendly bot detected: %s\n", result.BotName)
    } else {
        fmt.Println("This bot is not recognized as friendly.")
    }
}
```

### How It Works

Good Bot meticulously analyzes HTTP request headers, verifying user-agent strings and IP addresses against a curated list of known friendly bots. By focusing on valid domain names, CIDR ranges, ASNs, and specific user-agent patterns, Good Bot can accurately classify a bot's intentions, distinguishing between those that enhance your web ecosystem and those that do not.

### Contributing

Contributions to Good Bot are welcome! Whether it's enhancing detection logic, reporting bugs, or improving documentation, your input helps make Good Bot better for everyone.

### License

Good Bot is released under the MIT License. See the [LICENSE](LICENSE.txt) file for more information.

### Support

For support, please open an issue on our GitHub repository.
