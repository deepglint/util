package jsontime

import (
	"encoding/json"
	"testing"
	"time"
)

func TestFormatedTime(t *testing.T) {
	timestrs := []string{
		"1440151362",
		"1440151362000",
		"1440151362000000000",
		"2015-08-20T08:04:05.000+0800",
		"2015-08-20T08:04:05.000+08:00",
		"2015-08-20T08:04:05.000CST",
		"2015-08-20T00:04:05.000UTC",
		"2015-08-20T00:04:05.000Z",
		"2015-08-20T08:04:05+0800",
		"2015-08-20T08:04:05+08:00",
		"2015-08-20T08:04:05CST",
		"2015-08-20T00:04:05UTC",
		"2015-08-20T00:04:05Z",
		"2015-08-20T00:04:05.000ZUTC",
		"2015-08-20T00:04:05ZUTC",
		"2015-08-20T08:04:05.000ZCST", //wrong case
		"2015-08-20T08:04:05ZCST",     //wrong case
	}
	for _, timestr := range timestrs {
		thetime, err := Parse(timestr)
		if err != nil {
			t.Logf("error: %v", err)
		} else {
			t.Logf("%s as format %s", thetime.UTC().Format("2006-01-02T15:04:05.000-0700"), timestr)
		}
	}
	jsontimevalue := UnParse(time.Now(), "string")
	value, _ := json.Marshal(jsontimevalue)
	t.Logf("string format: %v", string(value))
	jsontimevalue = UnParse(time.Now(), "")
	value, _ = json.Marshal(jsontimevalue)
	t.Logf("nil format: %v", string(value))
	jsontimevalue = UnParse(time.Now(), "2006-01-02T15:04:05.000Z07:00")
	value, _ = json.Marshal(jsontimevalue)
	t.Logf("custumozed format: %v", string(value))
	t.Logf("ns format: %v", UnParse(time.Now(), "ns"))
	t.Logf("ms format: %v", UnParse(time.Now(), "ms"))
	t.Logf("s format: %v", UnParse(time.Now(), "s"))

}
