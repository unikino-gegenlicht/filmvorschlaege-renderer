package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pterm/pterm"
)

var dependencies []string = []string{"weasyprint"}

func init() {
	printHeader()
	checkDependencies()
}

func checkDependencies() {
	spinner, _ := pterm.DefaultSpinner.Start("Checking dependencies")
	time.Sleep(1000 * time.Millisecond)
	for _, dependency := range dependencies {
		if !commandExists(dependency) {
			spinner.Fail(fmt.Sprintf(`"%[1]s" is not installed. please install "%[1]s"`, dependency))
			os.Exit(1)
		}
		time.Sleep(500 * time.Millisecond)
	}
	spinner.Success("Required dependencies found")
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
