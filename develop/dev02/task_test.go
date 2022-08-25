package main

import (
	"errors"
	"testing"
)


func TestSuccess(t *testing.T) {
	testCases := map[string]string{
		"a4bc2d5e":  "aaaabccddddde",
		"abcd":      "abcd",
		"":          "",
		"qwe\\4\\5": "qwe45",
		"qwe\\45":   "qwe44444",
		"qwe\\\\5":  "qwe\\\\\\\\\\",
	}

	for test, expected := range testCases {
		res, err := Unpacking(test)
		if err != nil {
			t.Error(err)
		}
		if res != expected {
			t.Errorf("expected (%s) != result (%s)", expected, res)
		}
	}
}

func TestError(t *testing.T) {
	unpacking, err := Unpacking("45")
	if unpacking != "" && !errors.Is(err, errors.New("Invalid string")) {
		t.Errorf("expected (\"\") != result (%s) or err %v != %v", unpacking, err, errors.New("Invalid string"))
	}
}
