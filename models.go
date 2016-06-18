package main

import "fmt"

type Tomlconfig struct {
	Title   string
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
		###################################
		###################################
		#.----.-----.--------.-----.-----.#
		#|   _|  -__|        |  _  |  _  |#
		#|__| |_____|__|__|__|___  |_____|#
		#####################|_____|#######
		###################################
		`
	fmt.Println(banner)
}
