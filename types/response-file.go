package types

import "time"

type SummerResponseFile struct {
	Responses []response `json:"responses"`
}

func (receiver SummerResponseFile) GetFeatureMovies() (featureMovies map[string][]Movie) {
	featureMovies = make(map[string][]Movie)
	for _, response := range receiver.Responses {
		var movies []Movie
		movies = append(movies, Movie{
			Title:              response.Feat1Title,
			Country:            response.Feat1Country,
			Release:            time.Time(response.Feat1Release),
			Distributor:        response.Feat1Distributor,
			Description:        response.Feat1Description,
			IMDBUrl:            response.Feat1IMDBUrl,
			AudioLanguage:      response.Feat1Country,
			DistributorWarning: bool(response.Feat1DistributorWarning),
		})
		movies = append(movies, Movie{
			Title:              response.Feat2Title,
			Country:            response.Feat2Country,
			Release:            time.Time(response.Feat2Release),
			Distributor:        response.Feat2Distributor,
			Description:        response.Feat2Description,
			IMDBUrl:            response.Feat2IMDBUrl,
			AudioLanguage:      response.Feat2Country,
			DistributorWarning: bool(response.Feat2DistributorWarning),
		})
		featureMovies[response.Name] = movies
	}
	return featureMovies
}

func (receiver SummerResponseFile) GetDocumentations() (documentations map[string][]Movie) {
	documentations = make(map[string][]Movie)
	for _, response := range receiver.Responses {
		var movies []Movie
		movies = append(movies, Movie{
			Title:              response.DocTitle,
			Country:            response.DocCountry,
			Release:            time.Time(response.DocRelease),
			Distributor:        response.DocDistributor,
			Description:        response.DocDescription,
			IMDBUrl:            response.DocIMDBUrl,
			AudioLanguage:      response.DocCountry,
			DistributorWarning: bool(response.DocDistributorWarning),
		})
		documentations[response.Name] = movies
	}
	return documentations
}

type WinterResponseFile struct {
	Responses []winterResponse `json:"responses"`
}

func (receiver WinterResponseFile) GetFeatureMovies() (featureMovies map[string][]Movie) {
	featureMovies = make(map[string][]Movie)
	for _, response := range receiver.Responses {
		var movies []Movie
		movies = append(movies, Movie{
			Title:              response.Feat1Title,
			Country:            response.Feat1Country,
			Release:            time.Time(response.Feat1Release),
			Distributor:        response.Feat1Distributor,
			Description:        response.Feat1Description,
			IMDBUrl:            response.Feat1IMDBUrl,
			AudioLanguage:      response.Feat1Country,
			DistributorWarning: bool(response.Feat1DistributorWarning),
		})
		movies = append(movies, Movie{
			Title:              response.Feat2Title,
			Country:            response.Feat2Country,
			Release:            time.Time(response.Feat2Release),
			Distributor:        response.Feat2Distributor,
			Description:        response.Feat2Description,
			IMDBUrl:            response.Feat2IMDBUrl,
			AudioLanguage:      response.Feat2Country,
			DistributorWarning: bool(response.Feat2DistributorWarning),
		})
		featureMovies[response.Name] = movies
	}
	return featureMovies
}

func (receiver WinterResponseFile) GetDocumentations() (documentations map[string][]Movie) {
	documentations = make(map[string][]Movie)
	for _, response := range receiver.Responses {
		var movies []Movie
		movies = append(movies, Movie{
			Title:              response.DocTitle,
			Country:            response.DocCountry,
			Release:            time.Time(response.DocRelease),
			Distributor:        response.DocDistributor,
			Description:        response.DocDescription,
			IMDBUrl:            response.DocIMDBUrl,
			AudioLanguage:      response.DocCountry,
			DistributorWarning: bool(response.DocDistributorWarning),
		})
		documentations[response.Name] = movies
	}
	return documentations
}

func (receiver WinterResponseFile) GetHalloweenMovies() (halloweenMovies map[string][]Movie) {
	halloweenMovies = make(map[string][]Movie)
	for _, response := range receiver.Responses {
		var movies []Movie
		movies = append(movies, Movie{
			Title:              response.HalloweenTitle,
			Country:            response.HalloweenCountry,
			Release:            time.Time(response.HalloweenRelease),
			Distributor:        response.HalloweenDistributor,
			Description:        response.HalloweenDescription,
			IMDBUrl:            response.HalloweenIMDBUrl,
			AudioLanguage:      response.HalloweenCountry,
			DistributorWarning: bool(response.HalloweenDistributorWarning),
		})
		halloweenMovies[response.Name] = movies
	}
	return halloweenMovies
}

func (receiver WinterResponseFile) GetXmasMovies() (xmasMovies map[string][]Movie) {
	xmasMovies = make(map[string][]Movie)
	for _, response := range receiver.Responses {
		var movies []Movie
		movies = append(movies, Movie{
			Title:              response.ChristmasTitle,
			Country:            response.ChristmasCountry,
			Release:            time.Time(response.ChristmasRelease),
			Distributor:        response.ChristmasDistributor,
			Description:        response.ChristmasDescription,
			IMDBUrl:            response.ChristmasIMDBUrl,
			AudioLanguage:      response.ChristmasCountry,
			DistributorWarning: bool(response.ChristmasDistributorWarning),
		})
		xmasMovies[response.Name] = movies
	}
	return xmasMovies
}

type response struct {
	Name                    string `json:"name"`
	Feat1Title              string `json:"feat1title"`
	Feat1AudioLang          string `json:"feat1audioLang"`
	Feat1Country            string `json:"feat1country"`
	Feat1Distributor        string `json:"feat1distributor"`
	Feat1DistributorWarning SMBool `json:"feat1DistributorWarn"`
	Feat1IMDBUrl            string `json:"feat1imdb"`
	Feat1Release            SMTime `json:"feat1release"`
	Feat1Description        string `json:"feat1description"`
	Feat2Title              string `json:"feat2title"`
	Feat2AudioLang          string `json:"feat2audioLang"`
	Feat2Country            string `json:"feat2country"`
	Feat2Distributor        string `json:"feat2distributor"`
	Feat2DistributorWarning SMBool `json:"feat2DistributorWarn"`
	Feat2IMDBUrl            string `json:"feat2imdb"`
	Feat2Release            SMTime `json:"feat2release"`
	Feat2Description        string `json:"feat2description"`
	DocTitle                string `json:"docTitle"`
	DocAudioLang            string `json:"docAudioLang"`
	DocCountry              string `json:"docCountry"`
	DocDistributor          string `json:"docDistributor"`
	DocDistributorWarning   SMBool `json:"docDistributorWarn"`
	DocIMDBUrl              string `json:"docImdb"`
	DocRelease              SMTime `json:"docRelease"`
	DocDescription          string `json:"docDescription"`
}

type winterResponse struct {
	response
	HalloweenTitle              string `json:"halloweenTitle"`
	HalloweenAudioLang          string `json:"halloweenAudioLang"`
	HalloweenCountry            string `json:"halloweenCountry"`
	HalloweenDistributor        string `json:"halloweenDistributor"`
	HalloweenDistributorWarning SMBool `json:"halloweenDistribWarn"`
	HalloweenIMDBUrl            string `json:"halloweenImdb"`
	HalloweenRelease            SMTime `json:"halloweenRelease"`
	HalloweenDescription        string `json:"halloweenDescription"`
	ChristmasTitle              string `json:"xmasTitle"`
	ChristmasAudioLang          string `json:"xmasAudioLang"`
	ChristmasCountry            string `json:"xmasCountry"`
	ChristmasDistributor        string `json:"xmasDistributor"`
	ChristmasDistributorWarning SMBool `json:"xmasDistributorWarn"`
	ChristmasIMDBUrl            string `json:"xmasImdb"`
	ChristmasRelease            SMTime `json:"xmasRelease"`
	ChristmasDescription        string `json:"xmasDescription"`
}
