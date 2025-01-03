package main

import (
	"fmt"
	"os"

	"github.com/pterm/pterm"
)

var blacklistFileName string

func init() {
	printHeader()

	file, err := os.CreateTemp("", "distributors*")
	defer file.Close()
	if err != nil {
		pterm.DefaultLogger.Fatal("unable to extract distributor list")
		os.Exit(1)
	}

	file.WriteString(distributorBlacklist)

	blacklistFileName = file.Name()

}

func printHeader() {
	p := pterm.DefaultCenter
	p.Println("Movie Proposal Renderer")
	msg := fmt.Sprintf("Version: %s\n", Commit)
	msg += fmt.Sprintf("Commit Date: %s\n", CommitTime)
	msg += "\n© Unikino GEGENLICHT (gegenlicht.net) — Licensed under the MIT License"
	p.WithCenterEachLineSeparately().Println(msg)
}
