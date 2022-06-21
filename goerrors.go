/*
Summary:
Create a function that can parse time strings into component values.

Requirements:
The function must parse a string into a struct containing:
- Hour, minute, and second integer components

If parsing fails, then a descriptive error must be returned

Write some unit tests to check your work

Notes:
Example time string: 14:07:33
Use the `strings` package from stdlib to get time components
Use the `strconv` package from stdlib to convert strings to ints
Use the `errors` package to generate errors
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour   uint64
	minute uint64
	second uint64
}

type TimeParseError struct {
	msg string
}

func (t *TimeParseError) Error() string {
	return "TimeParseError: " + t.msg
}

func parseTimeStr(s string) (Time, error) {
	var tv [3]uint64
	tslice := strings.Split(s, ":")
	if len(tslice) != 3 {
		return Time{0, 0, 0}, &TimeParseError{"Unexpected format. Expected 00:00:00"}
	}
	for i := 0; i < len(tslice); i++ {
		v, err := strconv.ParseUint(tslice[i], 10, 32)
		if err != nil {
			return Time{0, 0, 0}, &TimeParseError{"Could not convert value \"" + tslice[i] + "\" to uint64"}
		}
		tv[i] = v
	}
	return Time{tv[0], tv[1], tv[2]}, nil
}

func main() {
	fmt.Println(parseTimeStr("14:07:11"))
	fmt.Println(parseTimeStr("14:07:11:11"))
	fmt.Println(parseTimeStr("a:b:c"))
}
