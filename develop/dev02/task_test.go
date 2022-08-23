package main

import (
	"errors"
	"testing"
)


func TestUnpackingSuccess(t *testing.T) {
	testCases := map[string]string{
		"a4bc2d5e":  "aaaabccddddde",
		"abcd":      "abcd",
		"":          "",
/*		"qwe\\4\\5": "qwe45",
		"qwe\\45":   "qwe44444",
		"qwe\\\\5":  "qwe\\\\\\\\\\",*/
	}

	for test, control := range testCases {
		res, err := Unpacking(test)
		if err != nil {
			t.Error(err)
		}
		if res != control {
			t.Errorf("test case (%s): control (%s) != result (%s)", test, control, res)
		}
	}
}

func TestUnpackingError(t *testing.T) {
	controlErr := errors.New("invalid string")
	unpacking, err := Unpacking("45")
	if unpacking != "" && !errors.Is(err, controlErr) {
		t.Errorf("control (\"\") != result (%s) or err %v != %v", unpacking, err, controlErr)
	}
}
