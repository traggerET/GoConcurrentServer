package main

import (
	"flag"
)


type Server struct {
	port    string
	logfile string
}

var s Server

func init() {
	s = Server{
		port:    *flag.String("p", "8081", "TCP listening port"),
		logfile: *flag.String("f", "history.msg", "File for messages"),
	}
}

func main() {
	flag.Parse()
	Run()
}
