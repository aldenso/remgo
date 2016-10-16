package main

import (
	"fmt"
	"os"

	"github.com/spf13/afero"
)

const (
	msgexists  = "remgo.toml already exist."
	msgcreated = "remgo.toml created."
)

// Tomlconfig struct to read toml file components.
type Tomlconfig struct {
	Title   string
	LogDir  string
	Servers map[string]Server
	Tasks   map[string]Task
}

// Server struct to list a group of servers IPs, names or FDQN.
type Server struct {
	IPs []string
}

// Task struct to read command, role, user (optional) and set a log (optional)
// for stdout or stderr.
type Task struct {
	Command string
	SFTP    []string
	Role    string
	User    string
	Log     bool
}

// Resp struct to return stdout and stderr from ssh connection.
type Resp struct {
	Output []byte
	Error  error
}

// Input struct to set ssh connection parameters.
type Input struct {
	Command string
	SFTP    []string
	IP      string
	Port    int
	User    string
}

// Lines1 just a # line to separate the CLI output and make it more readable.
func Lines1() {
	fmt.Println("###############################################################")
}

// Lines2 just a = line to separate the CLI output and make it more readable.
func Lines2() {
	fmt.Println("===============================================================")
}

// Banner a nice banner to be displayed when you run de remgo program.
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

// CreateTemplate fucntion to create a base remgo.toml file
func CreateTemplate(fs afero.Fs) (string, error) {
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
sftp = [
# origin, destiny, action(PUT or GET), don't put spaces after commas
"/tmp/file1.txt,/tmp/file1_put.txt,PUT",
"/tmp/file2.txt,/tmp/file2_get.txt,GET"
]

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
	if _, err := fs.Stat(tomlfile); err != nil {
		if os.IsNotExist(err) {
			file, err := fs.Create(tomlfile)
			if err != nil {
				return "", err
			}
			defer file.Close()
			if _, err := file.Write([]byte(template)); err != nil {
				return "", err
			}
		}
	} else {
		return msgexists, nil
	}
	return msgcreated, nil
}
