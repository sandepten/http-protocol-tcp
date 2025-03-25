package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:42069")
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on 42069")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Connection Accepted...")

		lines := getLinesChannel(conn)
		for line := range lines {
			fmt.Printf("%s\n", line)
		}
	}
}

func getLinesChannel(conn net.Conn) <-chan string {
	lines := make(chan string)

	go func() {
		reader := bufio.NewReader(conn)

		for {
			readLine, err := reader.ReadString('\n')
			if err != nil {
				conn.Close()
			}
			readLine = strings.TrimSpace(readLine)
			if len(readLine) == 0 {
				conn.Close()
				close(lines)
				fmt.Println("Connection Closed!")
				break
			}
			lines <- readLine
		}
	}()

	return lines
}
