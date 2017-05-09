package jsontime

import (
	"fmt"
	"strconv"
	"time"
)

func UnParse(thetime time.Time, timeformat string) interface{} {
	var formatedtime interface{}
	if timeformat == "string" || timeformat == "" {
		formatedtime = JsonTime{thetime, LongForm}
	} else if timeformat == "ns" {
		formatedtime = thetime.UnixNano()
	} else if timeformat == "ms" {
		formatedtime = thetime.UnixNano() / 1e6
	} else if timeformat == "s" {
		formatedtime = thetime.Unix()
	} else {
		formatedtime = JsonTime{thetime, timeformat}
	}
	return formatedtime
}

func Parse(timestr string) (*time.Time, error) {
	var thetime time.Time
	var timeint int64
	timeint, err := strconv.ParseInt(timestr, 10, 64)
	if err != nil {
		//try to use string time format
		const longForm = "2006-01-02T15:04:05.000MST"
		thetime, err = time.Parse(longForm, timestr)
		if err != nil {
			const longForm = "2006-01-02T15:04:05.000Z0700"
			thetime, err = time.Parse(longForm, timestr)
			if err != nil {
				const longForm = "2006-01-02T15:04:05.000Z07:00"
				thetime, err = time.Parse(longForm, timestr)
				if err != nil {
					const longForm = "2006-01-02T15:04:05MST"
					thetime, err = time.Parse(longForm, timestr)
					if err != nil {
						const longForm = "2006-01-02T15:04:05Z0700"
						thetime, err = time.Parse(longForm, timestr)
						if err != nil {
							const longForm = "2006-01-02T15:04:05Z07:00"
							thetime, err = time.Parse(longForm, timestr)
							if err != nil {
								const longForm = "2006-01-02T15:04:05.000ZMST"
								thetime, err = time.Parse(longForm, timestr)
								if err != nil {
									const longForm = "2006-01-02T15:04:05ZMST"
									thetime, err = time.Parse(longForm, timestr)
									if err != nil {
										const longForm = "2006-01-12 15:04:05"
										thetime, err = time.Parse(longForm, timestr)
										if err != nil {
											thetime, err = time.Parse(time.UnixDate, timestr)
											if err != nil {
												return nil, fmt.Errorf("time format error")
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		//use unix time
		if timeint < 1e11 {
			thetime = time.Unix(timeint, 0)
		} else if timeint < 1e14 {
			thetime = time.Unix(0, timeint*1e6)
		} else {
			thetime = time.Unix(0, timeint)
		}
	}
	thetime = thetime.UTC()
	return &thetime, nil
}
