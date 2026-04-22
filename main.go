package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

var Version = "dev" // Overwritten by gobake build

type Quote struct {
	Text   string `json:"q"`
	Author string `json:"a"`
}

var fallbackQuotes = []Quote{
	{"The only way to do great work is to love what you do.", "Steve Jobs"},
	{"Life is what happens when you're busy making other plans.", "John Lennon"},
	{"Get busy living or get busy dying.", "Stephen King"},
	{"You only live once, but if you do it right, once is enough.", "Mae West"},
	{"In three words I can sum up everything I've learned about life: it goes on.", "Robert Frost"},
	{"To be yourself in a world that is constantly trying to make you something else is the greatest accomplishment.", "Ralph Waldo Emerson"},
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func fetchQuote() Quote {
	// Try ZenQuotes API
	resp, err := http.Get("https://zenquotes.io/api/random")
	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			body, err := io.ReadAll(resp.Body)
			if err == nil {
				var quotes []Quote
				if err := json.Unmarshal(body, &quotes); err == nil && len(quotes) > 0 {
					return quotes[0]
				}
			}
		}
	}
	// Fallback to random hardcoded quote
	return fallbackQuotes[rand.Intn(len(fallbackQuotes))]
}

func rainbowPrint(text string) {
	colors := []int{31, 33, 32, 36, 34, 35} // Red, Yellow, Green, Cyan, Blue, Magenta
	colorIndex := 0

	for _, char := range text {
		// Do not colorize whitespace characters to avoid weird background issues in some terminals
		if char == '\n' || char == ' ' || char == '\r' || char == '\t' {
			fmt.Print(string(char))
			continue
		}
		fmt.Printf("\033[%dm%c\033[0m", colors[colorIndex%len(colors)], char)
		colorIndex++
	}
}

func wrapText(text string, width int) []string {
	var lines []string
	words := strings.Fields(text)
	if len(words) == 0 {
		return lines
	}

	currentLine := words[0]
	for _, word := range words[1:] {
		if len(currentLine)+1+len(word) <= width {
			currentLine += " " + word
		} else {
			lines = append(lines, currentLine)
			currentLine = word
		}
	}
	lines = append(lines, currentLine)
	return lines
}

func buildBubble(text, author string) string {
	maxWidth := 50
	lines := wrapText(text, maxWidth)

	authorLine := fmt.Sprintf("— %s", author)
	if len(authorLine) > maxWidth {
		maxWidth = len(authorLine)
	}
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	var sb strings.Builder
	sb.WriteString(" " + strings.Repeat("_", maxWidth+2) + "\n")

	for i, line := range lines {
		padding := strings.Repeat(" ", maxWidth-len(line))
		if len(lines) == 1 {
			sb.WriteString(fmt.Sprintf("< %s%s >\n", line, padding))
		} else if i == 0 {
			sb.WriteString(fmt.Sprintf("/ %s%s \\\n", line, padding))
		} else if i == len(lines)-1 {
			sb.WriteString(fmt.Sprintf("\\ %s%s /\n", line, padding))
		} else {
			sb.WriteString(fmt.Sprintf("| %s%s |\n", line, padding))
		}
	}

	// Add an empty line before the author
	sb.WriteString(fmt.Sprintf("| %s |\n", strings.Repeat(" ", maxWidth)))

	// Add author aligned to the right
	authorPadding := strings.Repeat(" ", maxWidth-len(authorLine))
	sb.WriteString(fmt.Sprintf("| %s%s |\n", authorPadding, authorLine))

	sb.WriteString(" " + strings.Repeat("-", maxWidth+2) + "\n")
	sb.WriteString("        \\   ^__^\n")
	sb.WriteString("         \\  (oo)\\_______\n")
	sb.WriteString("            (__)\\       )\\/\\\n")
	sb.WriteString("                ||----w |\n")
	sb.WriteString("                ||     ||\n")

	return sb.String()
}

func printHelp() {
	fmt.Printf("atlas.quote v%s - A cowsay-like quote generator\n\n", Version)
	fmt.Println("Usage:")
	fmt.Println("  atlas.quote [flags]")
	fmt.Println("\nFlags:")
	fmt.Println("  -h, --help       Show this help message")
	fmt.Println("  -v, --version    Show version number")
	fmt.Println("  -c, --color      Output the quote in rainbow colors")
}

func main() {
	colorMode := false

	for _, arg := range os.Args[1:] {
		if arg == "-h" || arg == "--help" {
			printHelp()
			return
		}
		if arg == "-v" || arg == "--version" {
			fmt.Printf("atlas.quote v%s\n", Version)
			return
		}
		if arg == "-c" || arg == "--color" {
			colorMode = true
		}
	}

	q := fetchQuote()
	bubble := buildBubble(q.Text, q.Author)

	if colorMode {
		rainbowPrint(bubble)
	} else {
		fmt.Print(bubble)
	}
}
