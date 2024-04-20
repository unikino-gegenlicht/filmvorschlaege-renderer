package types

import "time"

type Movie struct {
	Title              string
	Country            string
	Release            time.Time
	Distributor        string
	Description        string
	IMDBUrl            string
	AudioLanguage      string
	DistributorWarning bool
}
