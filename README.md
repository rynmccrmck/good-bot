# Good Bot

[![codecov](https://codecov.io/gh/rynmccrmck/good-bot/graph/badge.svg?token=4X6S00774G)](https://codecov.io/gh/rynmccrmck/good-bot)

<img src="https://github.com/rynmccrmck/good-bot/assets/5178938/ece7cfe2-4369-4da6-a440-915cbc48368c" alt="drawing" width="200"/>

**Good Bot** is an open-source Go library designed to enhance web application security and user experience by distinguishing beneficial automated agents, or "good bots", from potentially harmful traffic. In the digital ecosystem where bots play a crucial role—from search engine indexing to social media insights and link previews—it's essential to identify and welcome these friendly bots. Good Bot equips Go developers with the tools to recognize these agents accurately, ensuring your analytics remain accurate and your services optimized.

### Features

- **Accurate Bot Recognition**: Utilizes a comprehensive database of user-agent strings, IP addresses, DNS verification methods, and more to identify good bots with high precision.
- **High Performance**: With embedded data, Good Bot starts quickly and operates efficiently, requiring no external dependencies.
- **Flexibility and Customization**: The bot database is easily extendable and supports customization to align with various application requirements. Contributions are highly encouraged!

### How It Works

Good Bot meticulously analyzes HTTP request headers, verifying user-agent strings and IP addresses against a curated list of known friendly bots. By focusing on valid domain names, CIDR ranges, ASNs, and specific user-agent patterns, Good Bot can accurately classify a bot's intentions, distinguishing between those that enhance your web ecosystem and those that do not.

### Getting Started

To use Good Bot in your Go project, simply add it as a dependency:

```bash
go get github.com/rynmccrmck/good-bot
```

#### Basic Usage

Here's a quick example of how to use Good Bot to detect whether a request comes from a known good bot:

```go
package main

import (
    "fmt"
    goodbot "github.com/rynmccrmck/good-bot"
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

#### Bulk Verifier Tool

Goodbot includes a command-line tool, bulkVerifier, designed to process a CSV file and identify whether each entry (based on user agent and IP address) corresponds to a known good bot. This tool adds two columns to the output CSV: is_good_bot (true/false) and bot_name (the name of the bot if identified).

##### Installation
Ensure you have Go installed on your system. Clone the repository and navigate to the bulkVerifier directory:

```sh
git clone https://github.com/rynmccmrmck/good-bot.git
cd good-bot/cmd/bulkVerifier
```

Build the tool with Go:

```sh
go build -o bulkVerifier
```

##### Usage

After building the tool, you can run it directly from the command line:

```sh
./bulkVerifier <input.csv> <output.csv>
```

- `<input.csv>`: Path to the input CSV file containing the data to be processed. The CSV should have headers, with the first two columns being `user_agent` and `ip_address`.
- `<output.csv>`: Path where the output CSV file will be saved. This file will include the original data plus two additional columns: `is_good_bot` and `bot_name`.

##### Input File Format

The input CSV file should be formatted with at least two columns: `user_agent` and `ip_address`. Here's an example:

```csv
user_agent,ip_address
Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html),66.249.66.1
```

##### Output File Format

The output CSV file will include the same data as the input file, with two additional columns indicating whether each entry is a known good bot and the bot's name if identified:

```csv
user_agent,ip_address,bot_status,bot_name
Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html),66.249.66.1,friendly,Googlebot
```

##### Example

To process an input file named `requests.csv` and save the output to `results.csv`, use the following command:

```sh
./bulkVerifier requests.csv results.csv
```

This will analyze each row in `requests.csv`, determine if the user agent and IP address match a known good bot, and append the results to `results.csv`.

### Contributing

Contributions to Good Bot are welcome! Whether it's enhancing detection logic, reporting bugs, or improving documentation, your input helps make Good Bot better for everyone.

### License

Good Bot is released under the MIT License. See the [LICENSE](LICENSE.txt) file for more information.

### Support

For support, please open an issue on our GitHub repository.
