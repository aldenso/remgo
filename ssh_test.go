package main

import (
	"strings"
	"testing"

	"github.com/aldenso/remgo/remgofs"
)

// this test assumes you can login to localhost with root
func TestDialer(t *testing.T) {
	newfs := remgofs.InitMemFs()
	filetoput := "/tmp/file1_put.txt"
	newfs.Create(filetoput)
	newfs.Chmod("/badfs", 0000)
	sshconfig := &Input{
		Command: "whoami",
		IP:      "127.0.0.1",
		Port:    22,
		User:    "root",
		SFTP: []string{
			"/etc/passwd,/tmp/file2_get.txt,GET",
			"/tmp/file1_put.txt,/tmp/file1_put.txt,PUT",
			//file2_put doesn't exists
			"/tmp/file2_put.txt,/tmp/file2_put.txt,PUT",
			//file2_get doesn't exists
			"/tmp/file2_get.txt,/tmp/file2_geet.txt,GET",
			//bad sftp request
			"/tmp/file2_get.txt,/tmp/file2_geet.txt,PULL",
		},
	}
	outexpected := "root"
	outresponse := &Resp{}
	outresponse = Dialer(newfs, sshconfig)
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
	outresponse = Dialer(newfs, sshconfig)
	if outresponse.Error == nil {
		t.Errorf("error got 'nil' value for expected error")
	}
}
