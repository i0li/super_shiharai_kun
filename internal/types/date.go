package types

import (
	"strings"
	"time"
)

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(b []byte) error {
	dateStr := strings.Trim(string(b), `"`)
	dateTime, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return err
	}
	d.Time = dateTime
	return nil
}
