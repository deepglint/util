package pad

import (
	"fmt"
	"strings"
)

//TODO convert these into a
/*
 * leftPad and rightPad just repoeat the padStr the indicated
 * number of times
 *
 */
func parseString(arg interface{}) string {
	switch arg.(type) {
	case string:
		return arg.(string)
	case int, int16, int32, int64:
		return fmt.Sprintf("%d", arg)
	case uint, uint16, uint32, uint64:
		return fmt.Sprintf("%x", arg)
	default:
		return fmt.Sprint("unknown")
	}
}

func LeftPad(s interface{}, padStr string, pLen int) string {
	return strings.Repeat(padStr, pLen) + parseString(s)
}
func RightPad(s interface{}, padStr string, pLen int) string {
	return parseString(s) + strings.Repeat(padStr, pLen)
}

/* the Pad2Len functions are generally assumed to be padded with short sequences of strings
 * in many cases with a single character sequence
 *
 * so we assume we can build the string out as if the char seq is 1 char and then
 * just substr the string if it is longer than needed
 *
 * this means we are wasting some cpu and memory work
 * but this always get us to want we want it to be
 *
 * in short not optimized to for massive string work
 *
 * If the overallLen is shorter than the original string length
 * the string will be shortened to this length (substr)
 *
 */

func RightPad2Len(s interface{}, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = parseString(s) + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}
func LeftPad2Len(s interface{}, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + parseString(s)
	return retStr[(len(retStr) - overallLen):]
}
