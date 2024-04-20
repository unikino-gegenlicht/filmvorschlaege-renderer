package types

import (
	"fmt"
	"time"
)

type SMTime time.Time

func (t *SMTime) UnmarshalJSON(data []byte) error {
	fmt.Println(string(data))
	realTime, err := time.Parse(`"2006-01-02 15:04:05"`, string(data))
	if err != nil {
		return err
	}
	*t = SMTime(realTime)
	return nil
}
