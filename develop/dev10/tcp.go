package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		fmt.Println("connected to localhost:8080")

		bufReader := bufio.NewReader(conn)
		fmt.Println("Start reading")

		go func(conn net.Conn) {
			defer conn.Close()
			for {
				content, err := bufReader.ReadString('\n')
				if err != nil {
					fmt.Println(err)
					fmt.Println("cannot read from a buf")
				}
				fmt.Fprintln(conn, "message: ", string(content[:len(content)-1]), "received")
				fmt.Println(string(content))
			}

		}(conn)

	}

}
