package writer

import (
	"fmt"

	"github.com/tommibembel/subforge/internal/parser"
	"github.com/xuri/excelize/v2"
)

func WriteXLSX(path string, subtitles []parser.Subtitle) error {
	f := excelize.NewFile()
	defer f.Close()

	sheet := "Sheet1"

	headerStyle, err := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true},
		Alignment: &excelize.Alignment{Vertical: "top"},
	})
	if err != nil {
		return err
	}

	wrapStyle, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Vertical: "top", WrapText: true},
	})
	if err != nil {
		return err
	}

	headers := []string{"#", "Start Time", "End Time", "Duration", "Text"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
		f.SetCellStyle(sheet, cell, cell, headerStyle)
	}

	for row, sub := range subtitles {
		excelRow := row + 2
		values := []string{
			intToString(sub.Index),
			sub.Start,
			sub.End,
			sub.Duration,
			stripTags(sub.Text),
		}
		for col, v := range values {
			cell, _ := excelize.CoordinatesToCellName(col+1, excelRow)
			f.SetCellValue(sheet, cell, v)
			f.SetCellStyle(sheet, cell, cell, wrapStyle)
		}
	}

	f.SetColWidth(sheet, "A", "A", 6)
	f.SetColWidth(sheet, "B", "D", 16)
	f.SetColWidth(sheet, "E", "E", 50)

	return f.SaveAs(path)
}

func intToString(n int) string {
	return fmt.Sprintf("%d", n)
}
