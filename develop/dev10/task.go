package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"time"
)

//Реализовать простейший telnet-клиент.
//
//Примеры вызовов:
//go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

type Conn struct {
	host    string
	port    string
	timeout time.Duration
}

func parse() (t Conn) {
	flag.DurationVar(&t.timeout, "timeout", time.Second*10, "timeout for tcp connection")
	flag.Parse()
	if len(flag.Args()) < 2 {
		fmt.Println("error in parsing arguments... expected host and port")
	}
	t.host = flag.Arg(0)
	t.port = flag.Arg(1)
	return
}

func main() {
	p := parse()

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(p.host, p.port), p.timeout)
	if err != nil {
		time.Sleep(p.timeout)
		fmt.Fprintln(os.Stderr, err)
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		// reading from a socket
		reader := bufio.NewReader(conn)
		b, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				conn.Close()
				fmt.Println("Connection was closed by host")
			}
			return
		}
		fmt.Println(b)
	}()
	writer := bufio.NewWriter(conn)
	_, err = writer.ReadFrom(os.Stdin)
	if err != nil {
		fmt.Println(os.Stderr, err)
	}
	fmt.Println("EOF")
	conn.Close()
	wg.Wait()
	fmt.Println("connection closed")
}
