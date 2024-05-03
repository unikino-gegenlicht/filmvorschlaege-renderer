package types

type SMBool bool

const trueString = `"Ja"`
const falseString = `"Nein"`
const notApplicable = `"N/A"`

func (b *SMBool) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case trueString:
		*b = true
		break
	case falseString, notApplicable:
		*b = false
	default:
		*b = false
	}
	return nil
}
