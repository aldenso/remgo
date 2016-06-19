package main

import (
	"fmt"
	"os"
	"os/user"

	"flag"

	"github.com/BurntSushi/toml"
)

var (
	tomlfile   = "remgo.toml"
	template   bool
	SSHTimeout int
)

func init() {
	flag.BoolVar(&template, "template", false, "Create an example remgo.toml file")
	flag.IntVar(&SSHTimeout, "timeout", 5, "Set ssh timeout in seconds.")
}

func ReadTomlFile(tomlfile string) (*Tomlconfig, error) {
	var config *Tomlconfig
	if _, err := toml.DecodeFile(tomlfile, &config); err != nil {
		return nil, err
	}
	return config, nil
}

func main() {
	flag.Parse()
	if template {
		CreateTemplate()
	}
	Banner()
	config, err := ReadTomlFile(tomlfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Running %s\n", config.Title)

	var username string
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Can't get username")
	}
	username = usr.Username

	for taskKey, taskval := range config.Tasks {
		Lines1()
		fmt.Printf("Task: %s\n", taskKey)
		Lines1()
		if taskval.User != "" {
			username = taskval.User
		}
		for k, v := range config.Servers {
			if k == taskval.Role {
				fmt.Printf("Servers Role: %s\n", k)
				Lines2()
				for _, ip := range v.IPs {
					Lines2()
					fmt.Printf("IP: %s\n", ip)
					Lines2()
					output := &Resp{}
					input := &Input{
						taskval.Command,
						ip,
						22,
						username,
					}
					output = Dialer(input)
					if output.Error != nil {
						fmt.Printf("--- FAILED ---\n%v\n", output.Error)
						if taskval.Log {
							filename := taskKey + "_" + ip
							NewLog(config.LogDir, filename, "err", []byte(output.Error.Error()))
						}
					} else {
						fmt.Printf("+++ SUCCESS +++\n%s\n", string(output.Output))
						if taskval.Log {
							filename := taskKey + "_" + ip
							NewLog(config.LogDir, filename, "log", output.Output)
						}
					}
				}
			}
		}
	}
}
