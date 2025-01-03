package types

import "time"

type Response struct {
	// Name contains the name of the person proposing the movie
	Name string `json:"name"`

	SpecialProgramIncluded SurveyMonkeyBoolean `json:"proposeSpecialProg"`

	M1Title           string `json:"feat1title"`
	M1AudioLanguage   string `json:"feat1audioLang"`
	M1CountryOfOrigin string `json:"feat1country"`
	M1Distributor     string `json:"feat1distributor"`
	M1ReleaseDate     Time   `json:"feat1release"`
	M1IMDBUri         string `json:"feat1imdb"`
	M1Description     string `json:"feat1description"`

	M2Title           string `json:"feat2title"`
	M2AudioLanguage   string `json:"feat2audioLang"`
	M2CountryOfOrigin string `json:"feat2country"`
	M2Distributor     string `json:"feat2distributor"`
	M2ReleaseDate     Time   `json:"feat2release"`
	M2IMDBUri         string `json:"feat2imdb"`
	M2Description     string `json:"feat2description"`

	DTitle           string `json:"docTitle"`
	DAudioLanguage   string `json:"docAudioLang"`
	DCountryOfOrigin string `json:"docCountry"`
	DDistributor     string `json:"docDistributor"`
	DReleaseDate     Time   `json:"docRelease"`
	DIMDBUri         string `json:"docImdb"`
	DDescription     string `json:"docDescription"`

	SName        *string `json:"specialProgName"`
	SDescription *string `json:"specialProgDesc"`
	SFrequency   *string `json:"specialProgFrequency"`
}

func (r Response) FeatureMovies() (movies []Movie) {
	movies = append(movies, Movie{
		Title:           r.M1Title,
		AudioLanguage:   r.M1AudioLanguage,
		CountryOfOrigin: r.M1CountryOfOrigin,
		Distributor:     r.M1Distributor,
		ReleaseYear:     time.Time(r.M1ReleaseDate),
		Description:     r.M1Description,
		IMDBUrl:         r.M1IMDBUri,
	})
	movies = append(movies, Movie{
		Title:           r.M2Title,
		AudioLanguage:   r.M2AudioLanguage,
		CountryOfOrigin: r.M2CountryOfOrigin,
		Distributor:     r.M2Distributor,
		ReleaseYear:     time.Time(r.M2ReleaseDate),
		Description:     r.M2Description,
		IMDBUrl:         r.M2IMDBUri,
	})
	return
}

func (r Response) Documentary() (movies []Movie) {
	movies = append(movies, Movie{
		Title:           r.DTitle,
		AudioLanguage:   r.DAudioLanguage,
		CountryOfOrigin: r.DCountryOfOrigin,
		Distributor:     r.DDistributor,
		ReleaseYear:     time.Time(r.DReleaseDate),
		Description:     r.DDescription,
		IMDBUrl:         r.DIMDBUri,
	})
	return
}

func (r Response) Special() (specials []Special) {
	if r.SpecialProgramIncluded {
		return []Special{{
			Title:      *r.SName,
			ProposedBy: r.Name,
			Frequency:  *r.SFrequency,
		}}
	}
	return
}
