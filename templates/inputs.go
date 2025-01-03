package templates

import "github.com/unikino-gegenlicht/movie-proposal-renderer/types"

type Input struct {
	DistributorBlacklistFile string
	Stylesheet               string
	FeatureMovies            map[string][]types.Movie
	Documentaries            map[string][]types.Movie
	Specials                 []types.Special
}

type WinterInput struct {
	Input
	HalloweenMovies map[string][]types.Movie
	ChristmasMovies map[string][]types.Movie
}
