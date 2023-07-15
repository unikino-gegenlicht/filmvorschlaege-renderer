package main

import "time"

type ResponseFile struct {
	Responses []Response `json:"responses"`
}

type Response struct {
	Id                  string `json:"id"`
	Submitdate          string `json:"submitdate"`
	Lastpage            string `json:"lastpage"`
	Startlanguage       string `json:"startlanguage"`
	Seed                string `json:"seed"`
	Name                string `json:"name"`
	Film1Ttitel         string `json:"vorschlag1daten[m1title]"`
	Film1Originalttitel string `json:"vorschlag1daten[m1originalTitle]"`
	Film1Sprache        string `json:"vorschlag1daten[m1Language]"`
	Film1Land           string `json:"vorschlag1daten[m1country]"`
	Film1Genre          string `json:"vorschlag1daten[m1Genre]"`
	Film1Trailer        string `json:"vorschlag1daten[m1trailer]"`
	Film1IMDB           string `json:"vorschlag1daten[m1imdbid]"`
	Film1TMDB           string `json:"vorschlag1daten[m1tmdbid]"`
	Film1Release        string `json:"m1releaseYear"`
	Film1Bescheibung    string `json:"m1description"`
	Film2Titel          string `json:"m2daten[m2title]"`
	Film2OriginalTitel  string `json:"m2daten[m2originalTitle]"`
	Film2Sprache        string `json:"m2daten[m2language]"`
	Film2Land           string `json:"m2daten[m2coutries]"`
	Film2Genre          string `json:"m2daten[m2genre]"`
	Film2Trailer        string `json:"m2daten[m2trailer]"`
	Film2IMDB           string `json:"m2daten[m2imdbID]"`
	Film2TMDB           string `json:"m2daten[m2tmdbID]"`
	Film2Release        string `json:"m2releaseDate"`
	Film2Description    string `json:"m2description"`
}

func (r Response) ExtractMovies() []*Movie {
	var movies []*Movie
	if r.Film1Ttitel != "" && r.Film1Originalttitel != "" {
		t, _ := time.Parse("2006-01-02", r.Film1Release)
		m := Movie{
			ProposedBy:    r.Name,
			Title:         r.Film1Ttitel,
			OriginalTitle: r.Film1Originalttitel,
			Language:      r.Film1Sprache,
			Country:       r.Film1Land,
			Genre:         r.Film1Genre,
			IMDBId:        r.Film1IMDB,
			TMDBId:        r.Film1TMDB,
			TrailerLink:   r.Film1Trailer,
			ReleaseDate:   t,
			Description:   r.Film1Bescheibung,
		}
		movies = append(movies, &m)
	} else {
		movies = append(movies, nil)
	}
	if r.Film2Titel != "" && r.Film2OriginalTitel != "" {
		t, _ := time.Parse("2006-01-02", r.Film2Release)
		m := Movie{
			ProposedBy:    r.Name,
			Title:         r.Film2Titel,
			OriginalTitle: r.Film2OriginalTitel,
			Language:      r.Film2Sprache,
			Country:       r.Film2Land,
			Genre:         r.Film2Genre,
			IMDBId:        r.Film2IMDB,
			TMDBId:        r.Film2TMDB,
			TrailerLink:   r.Film2Trailer,
			ReleaseDate:   t,
			Description:   r.Film2Description,
		}
		movies = append(movies, &m)
	} else {
		movies = append(movies, nil)
	}
	return movies
}

type Movie struct {
	ProposedBy    string    `json:"proposedBy"`
	Title         string    `json:"title"`
	OriginalTitle string    `json:"originalTitle"`
	Language      string    `json:"language"`
	Country       string    `json:"country"`
	Genre         string    `json:"genre"`
	IMDBId        string    `json:"imdbId"`
	TMDBId        string    `json:"tmdbId"`
	ReleaseDate   time.Time `json:"releaseDate"`
	Description   string    `json:"description"`
	TrailerLink   string    `json:"trailerLink"`
}
