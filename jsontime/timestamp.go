package jsontime

import (
	"fmt"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func NewTimestamp(tm time.Time) *Timestamp {
	t := Timestamp(tm)
	return &t
}

type Timestamp time.Time

func (t *Timestamp) Time() time.Time {
	return time.Time(*t)
}

func (t *Timestamp) Format(layout string) string {
	return time.Time(*t).Format(layout)
}

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	ts := time.Time(*t).UnixNano() / int64(time.Millisecond)
	stamp := fmt.Sprint(ts)

	return []byte(stamp), nil
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	timestr := string(b)
	timeint, err := strconv.ParseInt(timestr, 10, 64)
	//try to use int time format
	if err != nil {
		//try to use string time format
		//const longForm = "2006-01-02T15:04:05.000ZMST"
		//parsedtime, err := time.Parse(longForm, string(b[1:len(b)-1]))
		parsedtime, err := Parse(string(b[1 : len(b)-1]))
		if err != nil {
			return err
		}
		*t = Timestamp(*parsedtime)
	} else {
		//unix timestamp
		*t = Timestamp(time.Unix(0, int64(timeint)*int64(time.Millisecond)))
	}
	//
	return nil
}

func (t *Timestamp) GetBSON() (interface{}, error) {
	if time.Time(*t).IsZero() {
		return nil, nil
	}

	return time.Time(*t), nil
}

func (t *Timestamp) SetBSON(raw bson.Raw) error {
	var tm time.Time

	if err := raw.Unmarshal(&tm); err != nil {
		return err
	}

	*t = Timestamp(tm)

	return nil
}

func (t *Timestamp) String() string {
	return time.Time(*t).String()
}

func (t *Timestamp) UTC() *Timestamp {
	_t := time.Time(*t)
	re := Timestamp(_t.UTC())
	return &re
}

func (t *Timestamp) Unix() int64 {
	return time.Time(*t).Unix()
}

func (t *Timestamp) UnixMilli() int64 {
	return time.Time(*t).UnixNano() / int64(time.Millisecond)
}

func (t *Timestamp) UnixNano() int64 {
	return time.Time(*t).UnixNano()
}
