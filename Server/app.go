package main

import (
	"log"
	"net"
	"os"
	"sync"
)

func Run() {
	f, err := os.OpenFile(s.logfile, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatalf("Cannot open file: %s", s.logfile)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err.Error())
		}
	}()
	conn, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		log.Println(err.Error())
		return
	}
	var connId int64 = 0
	var mu sync.Mutex
	for {
		cl, err := conn.Accept()

		if err != nil {
			log.Println(err.Error())
			continue
		}

		go HandleConnection(cl, connId, f, &mu)
		connId++
	}
}
