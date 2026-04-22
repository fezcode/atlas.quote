# Atlas Quote

![Banner Image](./banner-image.png)

**atlas.quote** is a beautiful cowsay-like quote generator for the terminal with a wise owl and a rainbow color mode. Part of the **Atlas Suite**, it provides inspiring quotes straight to stdout inside a neat unicode text bubble.

```text
 ╭────────────────────────────────────╮
 │ What we do now echoes in eternity. │
 │                                    │
 │                  — Marcus Aurelius │
 ╰─┬──────────────────────────────────╯
   │
   ╰─>  /\_/\
       ((@v@))
       ():::()
        VV-VV
```

## ✨ Features

- 🦉 **The Wise Owl:** Quotes are beautifully wrapped in an auto-sizing unicode text bubble with an ASCII owl.
- 🌈 **Rainbow Mode:** Output the entire bubble and quote in alternating ANSI rainbow colors.
- ✍️ **Custom Quotes:** Provide your own custom messages and authors.
- ⚡ **Dynamic Fallbacks:** Fetches daily quotes from the ZenQuotes API with hardcoded offline fallbacks.
- 📦 **Cross-Platform:** Binaries available for Windows, Linux, and macOS.

## 🚀 Installation

Requires [gobake](https://github.com/fezcode/gobake).

```bash
git clone https://github.com/fezcode/atlas.quote
cd atlas.quote
gobake build
```
Binaries will be placed in the `build/` directory.

## ⌨️ Usage

Simply run the binary to get a random quote:
```bash
./atlas.quote
```

### Rainbow Mode
Output the quote in alternating rainbow colors:
```bash
./atlas.quote -c
# or
./atlas.quote --color
```

### Custom Quotes
You can provide your own quote and author. If you omit the author, it defaults to "Anonymous":
```bash
./atlas.quote -m "Look what I made!" -s "The Developer" -c
```

## 🛠️ Help & Flags

```text
atlas.quote [flags]

Flags:
  -h, --help       Show this help message
  -v, --version    Show version number
  -c, --color      Output the quote in rainbow colors
  -m, --message    Provide a custom quote message
  -s, --said-by    Provide a custom author for the quote (used with -m)
```

## 📄 License
MIT License
