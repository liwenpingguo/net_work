package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	buffer := make([]byte,1024)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)

	for{
		line, _, _ := reader.ReadLine()

		conn.Write(line)

		read, _ := conn.Read(buffer)
		str := string(buffer[:read])

		if str == "bye"{
			break
		}
		fmt.Println("服务端说：",str)
	}

}
