package time

import (
	"time"
)

func WeekFromDate(day int, month int, year int) int {
	if month == 1 || month == 2 {
		month += 12
		year--
	}
	return (day+2*month+3*(month+1)/5+year+year/4-year/100+year/400)%7 + 1
}

func MonthFromSystime(month time.Month) int {
	switch month {
	case time.January:
		return 1
	case time.February:
		return 2
	case time.March:
		return 3
	case time.April:
		return 4
	case time.May:
		return 5
	case time.June:
		return 6
	case time.July:
		return 7
	case time.August:
		return 8
	case time.September:
		return 9
	case time.October:
		return 10
	case time.November:
		return 11
	case time.December:
		return 12
	}
	return -1
}

func DayFromWeekSearch(month int, year int, week int, num int) int {
	hit := 0
	for i := 1; i <= DaysInMonth(month, year); i++ {
		if WeekFromDate(i, month, year) == week {
			hit++
			if hit == num {
				return i
			}
		}
	}
	return -1
}

func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func DaysInMonth(month int, year int) int {
	switch month {
	case 1:
		return 31
	case 2:
		if IsLeapYear(year) {
			return 29
		} else {
			return 28
		}
	case 3:
		return 31
	case 4:
		return 30
	case 5:
		return 31
	case 6:
		return 30
	case 7:
		return 31
	case 8:
		return 31
	case 9:
		return 30
	case 10:
		return 31
	case 11:
		return 30
	case 12:
		return 31
	}
	return -1
}
