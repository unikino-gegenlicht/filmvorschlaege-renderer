package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/pterm/pterm"

	"github.com/unikino-gegenlicht/movie-proposal-renderer/types"
)

var availableSemesters = []string{"Wintersemester", "Sommersemester"}

func main() {
	// read the stylesheet
	stylesheetBytes, err := content.ReadFile("style.css")
	if err != nil {
		pterm.DefaultLogger.Fatal("unable to read stylesheet from embeds")
	}

	semester, _ := pterm.DefaultInteractiveSelect.WithOptions(availableSemesters).Show("Bitte wähle ein Semester aus")
	responseFilePath, _ := pterm.DefaultInteractiveTextInput.Show("Bitte gebe den Pfad zu der Antwortdatei ein")
	responseFile, err := os.Open(strings.TrimSpace(responseFilePath))
	if err != nil {
		pterm.DefaultLogger.Fatal("unable to open response file")

	}
	outputFile, err := os.Create("output.html")
	if err != nil {
		pterm.DefaultLogger.Fatal("unable to create temporary output file")
	}
	switch semester {
	case "Wintersemester":
		var responses types.WinterResponseFile
		err = json.NewDecoder(responseFile).Decode(&responses)
		if err != nil {
			pterm.DefaultLogger.Fatal("unable to read response file", pterm.DefaultLogger.ArgsFromMap(map[string]any{
				"error": err.Error(),
			}))
		}

		tmpl, err := template.ParseFS(content, "templates/winter.html")
		if err != nil {
			pterm.DefaultLogger.Fatal("unable to open parse embedded template file", pterm.DefaultLogger.ArgsFromMap(map[string]any{
				"error": err.Error(),
			}))
		}
		input := types.WinterTemplateInput{
			SummerTemplateInput: types.SummerTemplateInput{
				Stylesheet:    fmt.Sprintf("%s", stylesheetBytes),
				FeatureMovies: responses.GetFeatureMovies(),
				Documentaries: responses.GetDocumentations(),
			},
			HalloweenMovies: responses.GetHalloweenMovies(),
			XMasMovies:      responses.GetXmasMovies(),
		}

		err = tmpl.Execute(outputFile, input)
		if err != nil {
			pterm.DefaultLogger.Fatal("unable to execute template", pterm.DefaultLogger.ArgsFromMap(map[string]any{
				"error": err.Error(),
			}))
		}
		break
	case "Sommersemester":
		var responses types.SummerResponseFile
		err = json.NewDecoder(responseFile).Decode(&responses)
		if err != nil {
			pterm.DefaultLogger.Fatal("unable to read response file", pterm.DefaultLogger.ArgsFromMap(map[string]any{
				"error": err.Error(),
			}))
		}

		tmpl, err := template.ParseFS(content, "templates/sommer.html")
		if err != nil {
			pterm.DefaultLogger.Fatal("unable to open parse embedded template file", pterm.DefaultLogger.ArgsFromMap(map[string]any{
				"error": err.Error(),
			}))
		}
		input := types.SummerTemplateInput{
			Stylesheet:    fmt.Sprintf("%s", stylesheetBytes),
			FeatureMovies: responses.GetFeatureMovies(),
			Documentaries: responses.GetDocumentations(),
		}

		err = tmpl.Execute(outputFile, input)
		if err != nil {
			pterm.DefaultLogger.Fatal("unable to execute template", pterm.DefaultLogger.ArgsFromMap(map[string]any{
				"error": err.Error(),
			}))
		}
		break
	default:
		pterm.Println(pterm.Red("Unknown semester selected"))
		os.Exit(3)
	}
	err = outputFile.Close()
	if err != nil {
		pterm.DefaultLogger.Fatal("unable to close output file", pterm.DefaultLogger.ArgsFromMap(map[string]any{
			"error": err.Error(),
		}))
	}
}
