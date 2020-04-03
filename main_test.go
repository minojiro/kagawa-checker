package main

import (
	"testing"
)

func TestCheckKagawa(t *testing.T) {
	// https://www.pref.kagawa.lg.jp
	actual := checkKagawa("101.102.218.38")
	expext := true
	if actual != expext {
		t.Errorf("actual: %v, expext: %v", actual, expext)
	}
}

func TestCheckNotKagawa(t *testing.T) {
	// https://www.metro.tokyo.lg.jp
	actual := checkKagawa("27.110.42.78")
	expext := false
	if actual != expext {
		t.Errorf("actual: %v, expext: %v", actual, expext)
	}
}
