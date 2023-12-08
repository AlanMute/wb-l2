package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":13")
	if err != nil {
		fmt.Println(err)
	}
	conn, _ := ln.Accept()

	for {
		mes, err := bufio.NewReader(conn).ReadString('\n')
		if err == io.EOF {
			return
		}

		fmt.Print("Server: ", mes)

		mes = "Socket: " + mes

		conn.Write([]byte(mes + "\n"))
	}
}
