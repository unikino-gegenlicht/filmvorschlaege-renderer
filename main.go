package main

import (
	"encoding/json"
	"os"
	"strings"
	"text/template"

	"github.com/pterm/pterm"

	"github.com/unikino-gegenlicht/movie-proposal-renderer/templates"
	"github.com/unikino-gegenlicht/movie-proposal-renderer/types"
)

var availableSemesters = []string{"Wintersemester", "Sommersemester"}

func main() {
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
		var responses struct {
			Responses []types.WinterResponse `json:"responses"`
		}
		err = json.NewDecoder(responseFile).Decode(&responses)
		if err != nil {
			pterm.DefaultLogger.Fatal("unable to read response file", pterm.DefaultLogger.ArgsFromMap(map[string]any{
				"error": err.Error(),
			}))
		}

		var featureMovies, documentaries, halloweenMovies, christmasMovies map[string][]types.Movie
		featureMovies = make(map[string][]types.Movie)
		documentaries = make(map[string][]types.Movie)
		halloweenMovies = make(map[string][]types.Movie)
		christmasMovies = make(map[string][]types.Movie)
		specials := make([]types.Special, 0)

		for _, response := range responses.Responses {
			featureMovies[response.Name] = response.FeatureMovies()
			documentaries[response.Name] = response.Documentary()
			halloweenMovies[response.Name] = response.HalloweenMovie()
			christmasMovies[response.Name] = response.ChristmasMovie()
			specials = append(specials, response.Special()...)
		}

		tmpl, err := template.ParseFS(content, "templates/winter.html")
		if err != nil {
			pterm.DefaultLogger.Fatal("unable to open parse embedded template file", pterm.DefaultLogger.ArgsFromMap(map[string]any{
				"error": err.Error(),
			}))
		}

		input := templates.WinterInput{
			Input: templates.Input{
				Stylesheet:               stylesheet,
				DistributorBlacklistFile: blacklistFileName,
				FeatureMovies:            featureMovies,
				Documentaries:            documentaries,
				Specials:                 specials,
			},
			HalloweenMovies: halloweenMovies,
			ChristmasMovies: christmasMovies,
		}

		err = tmpl.Execute(outputFile, input)
		if err != nil {
			pterm.DefaultLogger.Fatal("unable to execute template", pterm.DefaultLogger.ArgsFromMap(map[string]any{
				"error": err.Error(),
			}))
		}
	case "Sommersemester":
		var responses struct {
			Responses []types.Response `json:"responses"`
		}
		err = json.NewDecoder(responseFile).Decode(&responses)
		if err != nil {
			pterm.DefaultLogger.Fatal("unable to read response file", pterm.DefaultLogger.ArgsFromMap(map[string]any{
				"error": err.Error(),
			}))
		}

		var featureMovies, documentaries map[string][]types.Movie
		featureMovies = make(map[string][]types.Movie)
		documentaries = make(map[string][]types.Movie)
		specials := []types.Special{}

		for _, response := range responses.Responses {
			featureMovies[response.Name] = response.FeatureMovies()
			documentaries[response.Name] = response.Documentary()
			specials = append(specials, response.Special()...)
		}
		tmpl, err := template.ParseFS(content, "templates/sommer.html")
		if err != nil {
			pterm.DefaultLogger.Fatal("unable to open parse embedded template file", pterm.DefaultLogger.ArgsFromMap(map[string]any{
				"error": err.Error(),
			}))
		}
		input := templates.Input{
			Stylesheet:               stylesheet,
			DistributorBlacklistFile: blacklistFileName,
			FeatureMovies:            featureMovies,
			Documentaries:            documentaries,
			Specials:                 specials,
		}

		err = tmpl.Execute(outputFile, input)
		if err != nil {
			pterm.DefaultLogger.Fatal("unable to execute template", pterm.DefaultLogger.ArgsFromMap(map[string]any{
				"error": err.Error(),
			}))
		}
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
