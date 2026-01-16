// Package internal contains the internal implementation of the fetcher.
package internal

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var (
	// Styles for the devotional output
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("212")).
			MarginBottom(1)

	verseStyle = lipgloss.NewStyle().
			Italic(true).
			Foreground(lipgloss.Color("39")).
			PaddingLeft(2).
			BorderLeft(true).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("39")).
			MarginBottom(1)

	bodyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252")).
			MarginBottom(1)

	readingStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("243")).
			Italic(true).
			MarginTop(1)

	// Regex to detect Bible reading references (e.g., "Genesis 36-38; Matthew 10:21-42")
	bibleRefRegex = regexp.MustCompile(`^[A-Z][a-z]+\s+\d+[-:\d]*;\s+[A-Z][a-z]+\s+\d+`)
)

func getTerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || width <= 0 {
		return 80 // default width
	}
	if width > 100 {
		return 100 // cap at 100 for readability
	}
	return width
}

func FetchDevotional(cmd *cobra.Command, args []string) error {
	width := getTerminalWidth()

	// Apply width to styles
	styledBody := bodyStyle.Width(width)
	styledVerse := verseStyle.Width(width - 3) // account for border (1) and padding (2)
	styledReading := readingStyle.Width(width)
	styledTitle := titleStyle.Width(width)

	// Fetch the devotional from the website
	url := "https://utmost.org/updated/today"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Parse the HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	out := cmd.OutOrStdout()

	// Extract and display the title
	title := strings.TrimSpace(doc.Find(".elementor-heading-title").First().Text())
	if title != "" {
		fmt.Fprintln(out, styledTitle.Render(title))
	}

	// Extract the scripture verse reference
	verse := strings.TrimSpace(doc.Find(".devotional-verse").First().Text())
	if verse == "" {
		// Fallback: look for the verse in a text-editor widget
		doc.Find(".elementor-widget-text-editor p").Each(func(i int, p *goquery.Selection) {
			text := strings.TrimSpace(p.Text())
			if verse == "" && text != "" {
				verse = text
			}
		})
	}
	if verse != "" {
		fmt.Fprintln(out, styledVerse.Render(verse))
	}

	// Extract the main devotional content
	doc.Find(".elementor-widget-theme-post-content p").Each(func(i int, p *goquery.Selection) {
		text := strings.TrimSpace(p.Text())
		if text != "" {
			if bibleRefRegex.MatchString(text) {
				fmt.Fprintln(out, styledReading.Render(text))
			} else {
				fmt.Fprintln(out, styledBody.Render(text))
			}
		}
	})

	return nil
}
