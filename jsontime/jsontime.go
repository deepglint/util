package jsontime

import "time"

const LongForm = "2006-01-02T15:04:05.000MST"

type JsonTime struct {
	time.Time
	F string
}

func (j JsonTime) format() string {
	return j.Time.Format(j.F)
}

func (j JsonTime) MarshalText() ([]byte, error) {
	return []byte(j.format()), nil
}

func (j JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + j.format() + `"`), nil
}
