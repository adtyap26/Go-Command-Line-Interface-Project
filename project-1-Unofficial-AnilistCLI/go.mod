module CLIAPPS/anilist_cli

go 1.19

require (
	CLIAPPS/search v0.0.0-00010101000000-000000000000
	github.com/fatih/color v1.13.0
	github.com/janeczku/go-spinner v0.0.0-20150530144529-cf8ef1d64394
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d
)

require (
	github.com/gosuri/uitable v0.0.4 // indirect
	github.com/mattn/go-colorable v0.1.9 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/rivo/uniseg v0.4.2 // indirect
	golang.org/x/crypto v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/term v0.4.0 // indirect
)

replace CLIAPPS/search => ./search

replace CLIAPPS/box => ./box

replace CLIAPPS/topten => ./topten

replace CLIAPPS/layout => ./layout
