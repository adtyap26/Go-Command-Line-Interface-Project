module CLIAPPS/weatherCLI

go 1.19

replace CLIAPPS/data => ./data

require (
	CLIAPPS/data v0.0.0-00010101000000-000000000000
	github.com/Delta456/box-cli-maker/v2 v2.3.0
	github.com/fatih/color v1.13.0
	github.com/janeczku/go-spinner v0.0.0-20150530144529-cf8ef1d64394
)

require (
	github.com/gookit/color v1.5.2 // indirect
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/mattn/go-colorable v0.1.9 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect
	golang.org/x/crypto v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/term v0.4.0 // indirect
)
