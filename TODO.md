# SubForge TODO

## GUI
- [ ] Layout verfeinern (Sidebar für Einstellungen analog zu srt_converter)
- [ ] Einstellungen persistieren (Ausgabeformat, Theme, Skalierung) via config-Datei
- [ ] "Open after convert"-Option
- [ ] Statusanzeige farblich hervorheben (grün/rot)
- [ ] App-Icon einbinden

## Funktionalität
- [ ] DOCX: Zell-Shading für Header-Zeile (aktuell nur Run-Shading, nicht volle Zellbreite)
- [ ] DOCX: Spaltenbreiten setzen (Spalte 1 schmal, Spalte 5 breit)
- [ ] Ausgabeformat "flow" im Dateidialog korrekt als .txt anzeigen
- [ ] Fehlerbehandlung in der GUI (Parse-Fehler, Schreibfehler anzeigen)

## Distribution
- [ ] Cross-Compilation für Windows (.exe): `GOOS=windows GOARCH=amd64 go build`
- [ ] Cross-Compilation für macOS (.app-Bundle): `GOOS=darwin go build`
- [ ] GitHub Actions Workflow für automatische Builds (Windows, macOS, Linux)
- [ ] fyne package für native App-Bundles prüfen (`fyne package`)

## Tests
- [ ] Unit-Tests für parser/srt.go (Randfälle: fehlende Leerzeile am Ende, CRLF, BOM)
- [ ] Unit-Tests für writer/txt.go und writer/docx.go
