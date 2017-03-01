package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = ""
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	l, err := net.Listen(CONN_TYPE, ":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	var (
		buf = make([]byte, 8192)
		r   = bufio.NewReader(conn)
		w   = bufio.NewWriter(conn)
	)

	defer conn.Close()

	for {
		n, err := r.Read(buf)
		if err != nil {
			return
		}

		w.Write(buf[:n])
		w.Flush()
	}
}
