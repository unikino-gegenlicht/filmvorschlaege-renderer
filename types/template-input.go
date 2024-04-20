package types

type SummerTemplateInput struct {
	Stylesheet    string
	FeatureMovies map[string][]Movie
	Documentaries map[string][]Movie
}

type WinterTemplateInput struct {
	SummerTemplateInput
	HalloweenMovies map[string][]Movie
	XMasMovies      map[string][]Movie
}
