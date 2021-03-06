package main

import (
	"fmt"
	"os"
	"os/user"
	"sync"

	"flag"

	"github.com/BurntSushi/toml"
	"github.com/aldenso/remgo/remgofs"
)

var (
	template bool
	tomlfile string
	// SSHTimeout is passed to ssh conf to avoid a hang connection from not
	// responding servers.
	SSHTimeout int
	// Fs afero fs to help later with testing
	Fs = remgofs.InitOSFs()
)

func init() {
	flag.BoolVar(&template, "template", false, "Create an example remgo.toml file.")
	flag.IntVar(&SSHTimeout, "timeout", 5, "Set ssh timeout in seconds.")
	flag.StringVar(&tomlfile, "t", "remgo.toml", "Specify a config file.")
}

func readTomlFile(tomlfile string) (*Tomlconfig, error) {
	var config *Tomlconfig
	if _, err := toml.DecodeFile(tomlfile, &config); err != nil {
		return nil, err
	}
	return config, nil
}

func main() {
	flag.Parse()
	if template {
		msg, err := CreateTemplate(Fs)
		if err != nil {
			fmt.Printf("Error creating tomfile: %v", err)
		} else {
			fmt.Println(msg)
			os.Exit(0)
		}
	}
	Banner()
	config, err := readTomlFile(tomlfile)
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
				var wg sync.WaitGroup
				for _, ip := range v.IPs {
					wg.Add(1)
					go func() {
						defer wg.Done()
						Lines2()
						fmt.Printf("IP: %s\n", ip)
						Lines2()
						output := &Resp{}
						input := &Input{
							taskval.Command,
							taskval.SFTP,
							ip,
							22,
							username,
						}
						output = Dialer(Fs, input)
						if output.Error != nil {
							fmt.Printf("--- FAILED ---\n%s\n%s\n", output.Error, string(output.Output))
							if taskval.Log {
								filename := taskKey + "_" + ip
								msg := fmt.Sprintf("%s%s", output.Output, output.Error)
								NewLog(config.LogDir, filename, "err", []byte(msg))
								//NewLog(config.LogDir, filename, "err", []byte(output.Error.Error()))
							}
						} else {
							fmt.Printf("+++ SUCCESS +++\n%s\n", string(output.Output))
							if taskval.Log {
								filename := taskKey + "_" + ip
								NewLog(config.LogDir, filename, "log", output.Output)
							}
						}
					}()
					wg.Wait()
				}
			}
		}
	}
}
