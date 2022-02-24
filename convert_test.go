package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConvertStringsToInts(t *testing.T) {
	tests := []struct {
		in     []string
		fail   bool
		expect []int
	}{
		{
			in:     []string{"1", "2", "3"},
			fail:   false,
			expect: []int{1, 2, 3},
		},
		{
			in:   []string{"1", "2", "xt42"},
			fail: true,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			res, err := Convert(test.in, ParseInt)
			if err != nil {
				if !test.fail {
					t.Fatalf("expect NOT to fail, but got %v", err)
				}
			} else {
				if !reflect.DeepEqual(test.expect, res) {
					t.Fatalf("want %v, have %v", test.expect, res)
				}
			}
		})
	}
}

func TestConvertStringsToFloats(t *testing.T) {
	tests := []struct {
		in     []string
		fail   bool
		expect []float64
	}{
		{
			in:     []string{"1.2", "2.3", "3.4"},
			fail:   false,
			expect: []float64{1.2, 2.3, 3.4},
		},
		{
			in:   []string{"1", "2", "xt42"},
			fail: true,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			res, err := Convert(test.in, ParseFloat)
			if err != nil {
				if !test.fail {
					t.Fatalf("expect NOT to fail, but got %v", err)
				}
			} else {
				if !reflect.DeepEqual(test.expect, res) {
					t.Fatalf("want %v, have %v", test.expect, res)
				}
			}
		})
	}
}

func TestConvertStringsToBools(t *testing.T) {
	tests := []struct {
		in     []string
		fail   bool
		expect []bool
	}{
		{
			in:     []string{"true", "false", "true"},
			fail:   false,
			expect: []bool{true, false, true},
		},
		{
			in:   []string{"true", "false", "42"},
			fail: true,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			res, err := Convert(test.in, ParseBool)
			if err != nil {
				if !test.fail {
					t.Fatalf("expect NOT to fail, but got %v", err)
				}
			} else {
				if !reflect.DeepEqual(test.expect, res) {
					t.Fatalf("want %v, have %v", test.expect, res)
				}
			}
		})
	}
}
