package main

import "testing"

func Test_readTomlFile(t *testing.T) {
	if _, err := readTomlFile(tomlfile); err != nil {
		t.Errorf("error: %v", err)
	}
	nonTomlfile := "non-tomfile"
	if _, err := readTomlFile(nonTomlfile); err == nil {
		t.Errorf("expected not nil, got 'nil'\n")
	}
}
