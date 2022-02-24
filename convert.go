package slices

import (
	"strconv"
	"strings"
)

func ParseInt(s string) (int, error) {
	n, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	return int(n), err
}

func ParseFloat(s string) (float64, error) {
	return strconv.ParseFloat(strings.TrimSpace(s), 64)
}

func ParseBool(s string) (bool, error) {
	return strconv.ParseBool(strings.TrimSpace(s))
}

func Convert[S any, D any](in []S, parseFnc func(t S) (D, error)) ([]D, error) {
	ts := make([]D, len(in))
	for i, st := range in {
		t, err := parseFnc(st)
		if err != nil {
			return nil, err
		}
		ts[i] = t
	}
	return ts, nil
}
