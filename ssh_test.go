package main

import (
	"strings"
	"testing"
)

// this test assumes you can login to localhost with root
func TestDialer(t *testing.T) {
	sshconfig := &Input{
		Command: "whoami",
		IP:      "127.0.0.1",
		Port:    22,
		User:    "root",
	}
	outexpected := "root"
	outresponse := &Resp{}
	outresponse = Dialer(sshconfig)
	if outresponse.Error != nil {
		t.Errorf("Can't connect to localhost: %v", outresponse.Error)
	}
	stringoutput := strings.TrimSpace(string(outresponse.Output))
	if stringoutput != outexpected {
		t.Errorf("Text mismatch, expected '%s', got '%s'.", outexpected, stringoutput)
	}
	sshconfig = &Input{
		Command: "whoami",
		IP:      "127.0.0.1",
		Port:    8888,
		User:    "root",
	}
	outresponse = Dialer(sshconfig)
	if outresponse.Error == nil {
		t.Errorf("error got 'nil' value for expected error")
	}
}
