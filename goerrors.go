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
	hour   int
	minute int
	second int
}

const (
	Hour = iota
	Minute
	Second
)

type TimeParseError struct {
	msg string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("TimeParseError: %v", t.msg)
}

func isValidRange(v int, component int) bool {
	if v < 0 {
		return false
	}
	switch component {
	case Hour:
		return v < 24
	case Minute, Second:
		return v < 60
	}
	return false
}

func parseTimeStr(s string) (Time, error) {
	var nv [3]int
	sv := strings.Split(s, ":")
	if len(sv) != 3 {
		return Time{}, &TimeParseError{fmt.Sprintf("Unexpected format. Received %v. Expected 00:00:00", s)}
	}
	for i := 0; i < len(sv); i++ {
		v, err := strconv.Atoi(sv[i])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Could not convert value \"%v\" to int", sv[i])}
		}
		if !isValidRange(v, i) {
			return Time{}, &TimeParseError{fmt.Sprintf("Value %v is not in valid range", v)}
		}
		nv[i] = v
	}
	return Time{nv[0], nv[1], nv[2]}, nil
}

func main() {
	fmt.Println(parseTimeStr("14:07:11"))
	fmt.Println(parseTimeStr("14:07:11:11"))
	fmt.Println(parseTimeStr("a:b:c"))
	fmt.Println(parseTimeStr("25:11:09"))
}
