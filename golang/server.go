package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	conn_port := flag.String("conn_port", "55555", "bind port")
	bufsize := flag.Int("bufsize", 8192, "send/recv bufsize")

	flag.Parse()

	l, err := net.Listen("tcp", ":"+*conn_port)

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
		go handleRequest(conn, *bufsize)
	}
}

func handleRequest(conn net.Conn, bufsize int) {
	var (
		buf = make([]byte, bufsize)
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
