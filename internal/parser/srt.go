package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Subtitle struct {
	Index    int
	Start    string
	End      string
	Duration string
	Text     string
}

var timeRe = regexp.MustCompile(`\d{2}:\d{2}:\d{2},\d{3} --> \d{2}:\d{2}:\d{2},\d{3}`)

func ParseSRT(path string) ([]Subtitle, error) {
	if !strings.HasSuffix(path, ".srt") {
		return nil, fmt.Errorf("input file must be an .srt file")
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var results []Subtitle
	var key int
	var textLines []string
	var start, end time.Time
	hasKey := false
	hasTime := false

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if idx, err := strconv.Atoi(strings.TrimSpace(line)); err == nil {
			key = idx
			hasKey = true
			continue
		}

		if timeRe.MatchString(line) {
			parts := strings.Split(line, " --> ")
			start, err = parseTimecode(strings.TrimSpace(parts[0]))
			if err != nil {
				return nil, fmt.Errorf("invalid timecode: %w", err)
			}
			end, err = parseTimecode(strings.TrimSpace(parts[1]))
			if err != nil {
				return nil, fmt.Errorf("invalid timecode: %w", err)
			}
			hasTime = true
			continue
		}

		if line == "" {
			if hasKey && hasTime && len(textLines) > 0 {
				dur := end.Sub(start)
				results = append(results, Subtitle{
					Index:    key,
					Start:    formatTimecode(start),
					End:      formatTimecode(end),
					Duration: formatDuration(dur),
					Text:     strings.Join(textLines, "\n"),
				})
				textLines = nil
				hasKey = false
				hasTime = false
			}
			continue
		}

		if hasKey && hasTime {
			textLines = append(textLines, line)
		}
	}

	// last entry may have no trailing blank line
	if hasKey && hasTime && len(textLines) > 0 {
		dur := end.Sub(start)
		results = append(results, Subtitle{
			Index:    key,
			Start:    formatTimecode(start),
			End:      formatTimecode(end),
			Duration: formatDuration(dur),
			Text:     strings.Join(textLines, "\n"),
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func parseTimecode(s string) (time.Time, error) {
	return time.Parse("15:04:05,000", s)
}

func formatTimecode(t time.Time) string {
	return t.Format("15:04:05,000")
}

func formatDuration(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	ms := int(d.Milliseconds()) % 1000
	return fmt.Sprintf("%02d:%02d:%02d,%03d", h, m, s, ms)
}
