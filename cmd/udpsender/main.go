package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	udp, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		panic(err)
	}
	conn, err := net.DialUDP("udp", nil, udp)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Listening on 42069")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		readLine, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		conn.Write([]byte(readLine))
	}
}
