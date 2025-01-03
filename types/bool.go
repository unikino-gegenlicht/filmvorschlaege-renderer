package types

import (
	"strings"
)

type SurveyMonkeyBoolean bool

func (b *SurveyMonkeyBoolean) UnmarshalJSON(src []byte) error {
	switch strings.TrimSpace(string(src)) {
	case `"Ja"`:
		*b = true
	case `"Nein"`, `"N/A"`:
		*b = false
	default:
		*b = false
	}
	return nil
}
