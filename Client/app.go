package main

import (
	"bufio"
	"net"
	"os"
	"strings"
)

func Run() {
	socket := strings.Join([]string{params.host, params.port}, ":")
	conn, err := net.Dial("tcp", socket)
	if err != nil {
		panic(err.Error())
	}
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err.Error())
		}
	}()

	err = Process(conn)
	if err != nil {
		panic(err.Error())
	}
}

func Process(conn net.Conn) error {
	reader := bufio.NewReader(os.Stdin)

	for {
		buff, err := reader.ReadBytes('\n')
		if err != nil {
			return err
		}

		if len(buff) > 4096 {
			buff = buff[:4096]
		}

		if string(buff) == "quit\n" {
			return nil
		}

		_, err = conn.Write(buff)
		if err != nil {
			return err
		}
	}
}
