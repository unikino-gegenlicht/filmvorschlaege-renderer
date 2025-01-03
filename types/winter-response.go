package types

import "time"

type WinterResponse struct {
	Response

	HTitle           string `json:"halloweenTitle"`
	HAudioLanguage   string `json:"halloweenAudioLang"`
	HCountryOfOrigin string `json:"halloweenCountry"`
	HDistributor     string `json:"halloweenDistributor"`
	HReleaseDate     Time   `json:"halloweenRelease"`
	HIMDBUri         string `json:"halloweenImdb"`
	HDescription     string `json:"halloweenDescription"`

	XTitle           string `json:"xmasTitle"`
	XAudioLanguage   string `json:"xmasAudioLang"`
	XCountryOfOrigin string `json:"xmasCountry"`
	XDistributor     string `json:"xmasDistributor"`
	XReleaseDate     Time   `json:"xmasRelease"`
	XIMDBUri         string `json:"xmasImdb"`
	XDescription     string `json:"xmasDescription"`
}

func (r WinterResponse) HalloweenMovie() (movies []Movie) {
	movies = append(movies, Movie{
		Title:           r.HTitle,
		AudioLanguage:   r.HAudioLanguage,
		CountryOfOrigin: r.HCountryOfOrigin,
		Distributor:     r.HDistributor,
		ReleaseYear:     time.Time(r.HReleaseDate),
		Description:     r.HDescription,
		IMDBUrl:         r.HIMDBUri,
	})
	return
}

func (r WinterResponse) ChristmasMovie() (movies []Movie) {
	movies = append(movies, Movie{
		Title:           r.XTitle,
		AudioLanguage:   r.XAudioLanguage,
		CountryOfOrigin: r.XCountryOfOrigin,
		Distributor:     r.XDistributor,
		ReleaseYear:     time.Time(r.XReleaseDate),
		Description:     r.XDescription,
		IMDBUrl:         r.XIMDBUri,
	})
	return
}
