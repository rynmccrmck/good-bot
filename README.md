# Good Bot

<img src="https://github.com/rynmccrmck/good-bot/assets/5178938/ece7cfe2-4369-4da6-a440-915cbc48368c" alt="drawing" width="200"/>

Good Bot is an open-source Go library designed to enhance web application security and user experience by accurately identifying beneficial automated agents, commonly known as "good bots". These agents perform various essential tasks, including search engine indexing, social media insights, link previews, and more. With bot traffic constituting a significant portion of online interactions, distinguishing between helpful and potentially harmful bots is vital for maintaining performance, security, and accurate analytics. Good Bot provides Go developers with a reliable toolset for this purpose.

### Features

- **Accurate Detection**: Leverages an extensive database of user-agent strings, IP addresses, and DNS verification methods to recognize good bots with high accuracy.
- **Efficiency**: Embedded data ensures quick startup and high-performance operation without external dependencies.
- **Customization**: Allows for easy updates to the bot database and supports custom configurations to meet diverse application needs. Contributions encouraged!

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
    "github.com/rynmccmrmck/goodbot/pkg/goodbot"
)

func main() {
    userAgent := "<USER_AGENT_STRING>"
    ipAddress := "<IP_ADDRESS>"

    isBot, botName := goodbot.IsGoodBot(userAgent, ipAddress)
    if isBot {
        fmt.Printf("Good bot detected: %s\n", botName)
    } else {
        fmt.Println("No good bot detected.")
    }
}
```

### How It Works

Good Bot analyzes HTTP request headers for signatures known to belong to good bots, including user-agent strings and IP addresses. Valid domains, CIDR ranges and ASNs and useragent combinations yeild a "good bot" result.

### Contributing

We welcome contributions to Good Bot! Enhancing detection capabilities, submitting bug reports, or improving documentation, your help is invaluable in making Good Bot more effective for everyone.

### License

Good Bot is released under the MIT License. See the [LICENSE](LICENSE.txt) file for more information.

### Support

For support, please open an issue on our GitHub repository. 
