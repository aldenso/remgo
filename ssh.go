package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

var usr string

// Dialer function to create ssh connection.
func Dialer(I *Input) *Resp {
	usr, err := user.Current()
	output := &Resp{}
	keyfile := (usr.HomeDir + "/.ssh/id_rsa")
	key, err := ioutil.ReadFile(keyfile)
	if err != nil {
		log.Printf("unable to read private key: %v\n", err)
		output.Error = err
		return output
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Printf("unable to parse private key: %v\n", err)
		output.Error = err
		return output
	}

	config := &ssh.ClientConfig{
		User: I.User,
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
		},
		Timeout: time.Duration(SSHTimeout) * time.Second,
	}
	//Create dial
	client, err := ssh.Dial("tcp", I.IP+":"+strconv.Itoa(I.Port), config)
	if err != nil {
		fmt.Printf("Can't Dial\n")
		output.Error = err
		return output
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		fmt.Printf("Failed to create session: %s\n", err.Error())
		output.Error = err
		return output
	}
	defer session.Close()

	if len(I.SFTP) != 0 {
		for _, v := range I.SFTP {
			Lines1()
			fmt.Printf("SFTP: %s ", v)
			sftp, err := sftp.NewClient(client)
			if err != nil {
				fmt.Printf("Error creating sftp client: %v\n", err)
			}
			defer sftp.Close()
			switch {
			case strings.Split(v, ",")[2] == "PUT":
				originfile, err := os.Open(strings.Split(v, ",")[0])
				if err != nil {
					fmt.Printf("Can't read origin file for sftp: %v\n", err)
					break
				}
				defer originfile.Close()
				destinyfile, err := sftp.Create(strings.Split(v, ",")[1])
				if err != nil {
					fmt.Printf("Can't create destiny file for sftp: %v\n", err)
					break
				}
				defer destinyfile.Close()
				if _, err := io.Copy(destinyfile, originfile); err != nil {
					fmt.Printf("Can't write in destiny file: %v\n", err)
					break
				}
				fmt.Println("SUCCESS")
			case strings.Split(v, ",")[2] == "GET":
				originfile, err := sftp.Open(strings.Split(v, ",")[0])
				if err != nil {
					fmt.Printf("Can't open remote file for sftp: %v\n", err)
					break
				}
				defer originfile.Close()
				destinyfile, err := os.Create(strings.Split(v, ",")[1])
				if err != nil {
					fmt.Printf("Can't create local file for sftp: %v\n", err)
					break
				}
				defer destinyfile.Close()
				_, err = io.Copy(destinyfile, originfile)
				if err != nil {
					fmt.Printf("Can't copy files for sftp: %v\n", err)
					break
				}
				fmt.Println("SUCCESS")
			default:
				fmt.Printf("Action '%s' for sftp not valid\n", strings.Split(v, ",")[2])
			}
			Lines1()
		}
	}

	var stdout []byte
	stdout, err = session.CombinedOutput(I.Command)
	output.Output = stdout
	output.Error = err
	return output
}
