package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	"strconv"
	"time"

	"golang.org/x/crypto/ssh"
)

var usr string

func Dialer(I *Input) *Resp {
	usr, err := user.Current()
	output := &Resp{}
	keyfile := (usr.HomeDir + "/.ssh/id_rsa")
	key, err := ioutil.ReadFile(keyfile)
	if err != nil {
		log.Printf("unable to read private key: %v", err)
		output.Error = err
		return output
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Printf("unable to parse private key: %v", err)
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
		fmt.Printf("Failed to create session: " + err.Error())
		output.Error = err
		return output
	}
	defer session.Close()

	//var b bytes.Buffer
	//session.Stdout = &b
	var stdout []byte
	if stdout, err = session.CombinedOutput(I.Command); err != nil {
		fmt.Printf("Failed to run: " + err.Error())
	}
	//fmt.Println(b.String())
	output.Output = stdout
	output.Error = nil
	return output
}
