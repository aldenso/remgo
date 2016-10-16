package main

import (
	"testing"

	"github.com/aldenso/remgo/remgofs"
)

func Test_CreateTemplate(t *testing.T) {
	newfs := remgofs.InitMemFs()
	_, err := CreateTemplate(newfs)
	if err != nil {
		t.Errorf("Error creating tomlfile: %v", err)
	}
	msg, err := CreateTemplate(newfs)
	if err != nil {
		t.Errorf("Error creating tomlfile: %v", err)
	}
	if msg != msgexists {
		t.Errorf("expected '%s', got '%s'", msgexists, msg)
	}
}
