package main

import (
	"strings"
	"testing"
)

// this test assumes you can login to localhost with root
func TestDialer(t *testing.T) {
	sshconfig := &Input{
		Command: "echo $HOME",
		IP:      "127.0.0.1",
		Port:    22,
		User:    "root",
	}
	outexpected := "/root"
	outresponse := &Resp{}
	outresponse = Dialer(sshconfig)
	if outresponse.Error != nil {
		t.Errorf("Can't connect to localhost: %v", outresponse.Error)
	}
	stringoutput := strings.TrimSpace(string(outresponse.Output))
	if stringoutput != outexpected {
		t.Errorf("Text mismatch, expected '%s', got '%s'.", outexpected, stringoutput)
	}
}
