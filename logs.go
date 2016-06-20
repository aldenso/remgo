package main

import (
	"fmt"
	"os"
	"time"
)

// NewLog to create a log file for stdout or stderr
func NewLog(logdir string, filename string, extension string, message []byte) {
	if logdir != "" {
		if _, err := os.Stat(logdir); err != nil {
			if os.IsNotExist(err) {
				if err := os.Mkdir(logdir, 0750); err != nil {
					fmt.Printf("Can't create logs directory\n%v\n", err)
				}
			}
		}
		file, err := os.Create(logdir + "/" + filename + "_" + time.Now().Format(time.RFC3339) + "." + extension)
		if err != nil {
			fmt.Println("Error creating log file", err)
		}
		defer file.Close()

		if _, err := file.Write(message); err != nil {
			fmt.Printf("Can't write message\n%v\n", err)
		}
	} else {
		file, err := os.Create(filename + "_" + time.Now().Format(time.RFC3339) + "." + extension)
		if err != nil {
			fmt.Println("Error creating log file", err)
		}
		defer file.Close()

		if _, err := file.Write(message); err != nil {
			fmt.Printf("Can't write message\n%v\n", err)
		}
	}

}
