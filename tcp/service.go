package main

import (
	"fmt"
	"net"
)

func conntent(conn net.Conn) error {
	buffer := make([]byte, 1024)

	defer conn.Close()

	for {
		read, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("err: %v", err)
			return err
		}
		str := string(buffer[:read])
		fmt.Println("客户端说：", str)

		if str == "off" {
			conn.Write([]byte("bye"))
			break
		}
	}
	return nil
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8888")

	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	conn, err := listener.Accept()

	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	conntent(conn)
}
