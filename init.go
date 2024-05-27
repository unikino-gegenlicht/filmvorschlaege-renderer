package main

import (
	"fmt"

	"github.com/pterm/pterm"
)

func init() {
	printHeader()
}

func printHeader() {
	p := pterm.DefaultCenter
	p.Println("Movie Proposal Renderer")
	msg := fmt.Sprintf("Version: %s\n", Commit)
	msg += fmt.Sprintf("Commit Date: %s\n", CommitTime)
	msg += "\n© Unikino GEGENLICHT (gegenlicht.net) — Licensed under the MIT License"
	p.WithCenterEachLineSeparately().Println(msg)
}
