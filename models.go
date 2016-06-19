package main

import (
	"fmt"
	"os"
)

type Tomlconfig struct {
	Title   string
	LogDir  string
	Servers map[string]Server
	Tasks   map[string]Task
}

type Server struct {
	IPs []string
}

type Task struct {
	Command string
	Role    string
	User    string
	Log     bool
}

type Resp struct {
	Output []byte
	Error  error
}

type Input struct {
	Command string
	IP      string
	Port    int
	User    string
}

func Lines1() {
	fmt.Println("###############################################################")
}

func Lines2() {
	fmt.Println("===============================================================")
}

func Banner() {
	banner :=
		`
		#####################################
		#####################################
		##.----.-----.--------.-----.-----.##
		##|   _|  -__|        |  _  |  _  |##
		##|__| |_____|__|__|__|___  |_____|##
		######################|_____|########
		#####################################
		`
	fmt.Println(banner)
}

func CreateTemplate() {
	template := `title = "Example of remgo Configuration"
logdir = "/tmp/logs"

[servers]
[servers.Internal]
IPs = ["192.168.125.100", "localhost" ]
[servers.External]
IPs = ["192.168.50.100", "server1.github.com"]

[tasks]
[tasks.CheckHostname]
role = "Internal"
command = "hostname"
log = true

[tasks.CheckDir]
role = "External"
command = "ls -l | tail -3"

[tasks.WhoamIandIP]
user = "username"
role = "Internal"
command = "whoami; ip addr show"
log = true
`
	tomlfile := "remgo.toml"
	if _, err := os.Stat(tomlfile); err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(tomlfile)
			if err != nil {
				fmt.Println("Error creating remgo.toml file", err)
				os.Exit(1)
			}
			defer file.Close()
			if _, err := file.Write([]byte(template)); err != nil {
				fmt.Printf("Can't write message\n%v\n", err)
				os.Exit(1)
			}
		}
	} else {
		fmt.Println("remgo.toml already exist.")
		os.Exit(1)
	}
	os.Exit(0)
}
