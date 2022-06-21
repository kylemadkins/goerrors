package main

import "testing"

func TestParseTimeStr(t *testing.T) {
	table := []struct {
		time     string
		ok       bool
		expected Time
	}{
		{"19:00:12", true, Time{19, 0, 12}},
		{"1:3:44", true, Time{1, 3, 44}},
		{"bad", false, Time{0, 0, 0}},
		{"1:-3:44", false, Time{0, 0, 0}},
		{"0:59:59", true, Time{0, 59, 59}},
		{"", false, Time{0, 0, 0}},
		{"11:22", false, Time{0, 0, 0}},
		{"aa:bb:cc", false, Time{0, 0, 0}},
		{"5:23:", false, Time{0, 0, 0}},
		{"04:62:10", false, Time{0, 0, 0}},
		{"04:62:10", false, Time{0, 0, 0}},
		{"99:05:32", false, Time{0, 0, 0}},
		{"07:45:83", false, Time{0, 0, 0}},
	}

	for _, data := range table {
		time, err := parseTimeStr(data.time)
		if data.ok && err != nil {
			t.Errorf("Error should be nil for %v: %v", data.time, err)
		}
		if time != data.expected {
			t.Errorf("Unexpected value for %v. Expected %v. Received %v", data.time, data.expected, time)
		}
	}
}
