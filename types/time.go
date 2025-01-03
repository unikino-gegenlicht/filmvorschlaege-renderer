package types

import "time"

type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) error {
	realTime, err := time.Parse(`"2006-01-02 15:04:05"`, string(data))
	if err != nil {
		return err
	}
	*t = Time(realTime)
	return nil
}
