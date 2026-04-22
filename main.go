package main

import (
	"bufio"
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

func displayQuote(q Quote) {
	fmt.Printf("\n\"%s\"\n  — %s\n\n", q.Text, q.Author)
}

func printHelp() {
	fmt.Printf("atlas.quote v%s - An interactive and non-interactive quote generator\n\n", Version)
	fmt.Println("Usage:")
	fmt.Println("  atlas.quote [flags]")
	fmt.Println("\nFlags:")
	fmt.Println("  -h, --help       Show this help message")
	fmt.Println("  -v, --version    Show version number")
	fmt.Println("  -i, --interactive Start in interactive mode")
}

func main() {
	interactive := false

	for _, arg := range os.Args[1:] {
		if arg == "-h" || arg == "--help" {
			printHelp()
			return
		}
		if arg == "-v" || arg == "--version" {
			fmt.Printf("atlas.quote v%s\n", Version)
			return
		}
		if arg == "-i" || arg == "--interactive" {
			interactive = true
		}
	}

	if !interactive {
		q := fetchQuote()
		displayQuote(q)
		return
	}

	// Interactive Mode
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Welcome to atlas.quote v%s (Interactive Mode)\n", Version)
	fmt.Println("Press Enter to get a new quote, or type 'q' to quit.")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "q" || strings.ToLower(input) == "quit" {
			fmt.Println("Goodbye!")
			break
		}

		q := fetchQuote()
		displayQuote(q)
	}
}
