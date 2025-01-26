package main

import "server_siem/service"

func main() {
	program := service.InitProgram("/config/config.ini")
	program.Work()
}
