package main

import (
	"github.com/pterm/pterm"
)

func init() {
	printHeader()
}

func printHeader() {
	headerPrinter := pterm.HeaderPrinter{
		TextStyle:       pterm.NewStyle(pterm.FgBlack),
		BackgroundStyle: pterm.NewStyle(pterm.BgLightYellow),
		FullWidth:       true,
		Margin:          0,
	}
	headerPrinter.Println("Movie Proposals Renderer")
	headerPrinter.Println("© 2024 Jan Eike Suchard — Licensed under the MIT License")
}
