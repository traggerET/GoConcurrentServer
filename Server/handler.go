package main

import (
	"bytes"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
)

func HandleConnection(conn net.Conn, connId int64, f *os.File, mu *sync.Mutex) {
	log.Printf("connection id %d established\n", connId)
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(conn)

	connLog := []byte(strconv.FormatInt(connId, 10) + ":")
	for {
		buff := make([]byte, 4096)
		_, err := conn.Read(buff)

		if err != nil && err != io.EOF {
			log.Println(err.Error())
			return
		}
		if err == io.EOF {
			return
		}

		buff = bytes.Trim(buff, "\x00")
		toWrite := append(connLog, buff...)

		mu.Lock()
		_, err = f.Write(toWrite)
		mu.Unlock()

		if err != nil {
			log.Println(err.Error())
			return
		}
	}
}
