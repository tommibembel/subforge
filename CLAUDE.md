# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

**SubForge** is a Go desktop GUI application that converts SRT subtitle files into DOCX, XLSX, TXT, and flowing-text formats. Built with the [Fyne](https://fyne.io/) cross-platform GUI toolkit.

## Commands

```bash
go run .          # Run the GUI application
go build          # Build the binary (output: subforge / subforge.exe)
go test ./...     # Run tests
go test -v -run TestName ./internal/...  # Run a specific test
```

Cross-platform builds (for future distribution):
```bash
GOOS=windows GOARCH=amd64 go build
fyne package -os linux   # Native app bundle
```

## Architecture

```
main.go           → ui.Run() (entry point, nothing else)
ui/app.go         → Fyne GUI: file dialogs, format selector, orchestrates conversion
internal/parser/srt.go   → Parses SRT into []Subtitle structs
internal/writer/
  txt.go          → WriteTXT(), WriteTXTFlowing()
  docx.go         → WriteDOCX()
  xlsx.go         → WriteXLSX()
```

**Core data structure** (`internal/parser/srt.go`):
```go
type Subtitle struct {
    Index    int
    Start    string   // HH:MM:SS,mmm
    End      string
    Duration string   // calculated
    Text     string   // may contain <i> tags
}
```

The UI layer calls the parser to get `[]Subtitle`, then dispatches to the appropriate writer based on the selected format. Writers handle their own file creation and formatting—DOCX preserves `<i>` italic tags, TXT/flowing text strips them.

## Key Dependencies

- `fyne.io/fyne/v2` — GUI framework
- `github.com/gomutex/godocx` — DOCX generation
- `github.com/xuri/excelize/v2` — XLSX generation

## Planned Work

See [TODO.md](TODO.md) for the full roadmap (written in German). High-priority items include unit tests for the parser, GUI refinements (layout, app icon, settings persistence), cross-platform distribution via GitHub Actions, and improved DOCX column-width control.
