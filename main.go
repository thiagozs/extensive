package main

import (
	"extensive-number/server"
	"flag"
	"fmt"
	"os"
)

var (
	port   = flag.String("p", ":8080", "Port for expose the server for requests")
	debug  = flag.Bool("d", false, "Turn on debug log on server")
	appver = flag.Bool("v", false, "prints current roxy version")
)

func main() {
	flag.Parse()

	version := "1.0.0"

	if *appver {
		fmt.Println("Version : ", version)
		os.Exit(0)
	}

	server := server.New(*port, *debug)

	server.RegisterRoutes()

	if err := server.Run(); err != nil {
		fmt.Println("Error on server, got : ", err.Error())
		os.Exit(1)
	}

}
