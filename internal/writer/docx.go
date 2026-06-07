package writer

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gomutex/godocx/docx"
	"github.com/gomutex/godocx/wml/stypes"
	"github.com/tommibembel/subforge/internal/parser"
)

var italicRe = regexp.MustCompile(`<i>(.*?)</i>`)

func WriteDOCX(path string, subtitles []parser.Subtitle) error {
	doc := docx.NewRootDoc()

	table := doc.AddTable()
	table.Style("TableGrid")

	headers := []string{"#", "Start Time", "End Time", "Duration", "Text"}
	headerRow := table.AddRow()
	for _, h := range headers {
		cell := headerRow.AddCell()
		para := cell.AddEmptyPara()
		run := para.AddText(h)
		run.Bold(true)
		run.Color("FFFFFF")
		run.Shading(stypes.ShdSolid, "auto", "000000")
	}

	for _, sub := range subtitles {
		row := table.AddRow()

		for _, val := range []string{fmt.Sprintf("%d", sub.Index), sub.Start, sub.End, sub.Duration} {
			cell := row.AddCell()
			cell.AddParagraph(val)
		}

		textCell := row.AddCell()
		para := textCell.AddEmptyPara()
		addTextWithItalics(para, sub.Text)
	}

	return doc.SaveTo(path)
}

// addTextWithItalics splits subtitle text on <i>...</i> tags and writes
// normal and italic runs accordingly.
func addTextWithItalics(para *docx.Paragraph, text string) {
	text = strings.TrimSuffix(text, "\n")

	if !strings.Contains(text, "<i>") {
		para.AddText(text)
		return
	}

	pointer := 0
	for _, m := range italicRe.FindAllStringSubmatchIndex(text, -1) {
		before := text[pointer:m[0]]
		italicText := text[m[2]:m[3]]
		pointer = m[1]

		if before != "" {
			para.AddText(before)
		}
		run := para.AddText(italicText)
		run.Italic(true)
	}

	if pointer < len(text) {
		para.AddText(text[pointer:])
	}
}
