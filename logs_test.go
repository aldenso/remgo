package main

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestNewLog(t *testing.T) {
	logdir, filename, extension, message := "/tmp/logs", "test", "log", []byte("hello")
	NewLog(logdir, filename, extension, message)
	file := logdir + "/" + filename + "_" + time.Now().Format(time.RFC3339) + "." + extension
	if _, err := os.Stat(logdir); err != nil {
		if os.IsNotExist(err) {
			t.Errorf("file '%s' not created", file)
		}
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		t.Errorf("Can't read file: %v", err)
	}
	if string(data) != string(message) {
		t.Errorf("mismatch, expected '%s', got '%s'.", string(message), string(data))
	}
	if err := os.Remove(file); err != nil {
		t.Errorf("Can't remove '%s', got error:%v", file, err)
	}
}
