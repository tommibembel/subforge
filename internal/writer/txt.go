package writer

import (
	"os"
	"regexp"
	"strings"

	"github.com/tommibembel/subforge/internal/parser"
)

var tagRe = regexp.MustCompile(`<[^>]+>`)

func stripTags(s string) string {
	return tagRe.ReplaceAllString(s, "")
}

// WriteTXT writes each subtitle as plain text, separated by blank lines.
func WriteTXT(path string, subtitles []parser.Subtitle) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, sub := range subtitles {
		clean := stripTags(sub.Text)
		if _, err := f.WriteString(clean + "\n\n"); err != nil {
			return err
		}
	}
	return nil
}

// WriteTXTFlowing joins all subtitle texts into a single flowing paragraph.
func WriteTXTFlowing(path string, subtitles []parser.Subtitle) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	parts := make([]string, 0, len(subtitles))
	for _, sub := range subtitles {
		clean := stripTags(sub.Text)
		clean = strings.ReplaceAll(clean, "\n", " ")
		parts = append(parts, clean)
	}

	_, err = f.WriteString(strings.Join(parts, " "))
	return err
}
