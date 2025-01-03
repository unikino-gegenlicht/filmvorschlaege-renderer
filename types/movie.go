package types

import (
	"bufio"
	"os"
	"strings"
	"time"
)

type Movie struct {
	Title           string
	AudioLanguage   string
	CountryOfOrigin string
	Distributor     string
	ReleaseYear     time.Time
	Description     string
	IMDBUrl         string
}

func (m Movie) DistributorWarning(blacklistFile string) bool {
	f, err := os.Open(blacklistFile)
	if err != nil {
		return true
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		blacklistedDistributor := scanner.Text()
		if strings.Contains(strings.TrimSpace(m.Distributor), blacklistedDistributor) {
			return true
		}
	}
	return false
}
