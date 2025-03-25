package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := make(chan string)
	go getLinesChannel(file, lines)
	for {
		line := <-lines
		fmt.Printf("read: %s\n", line)
	}
}

func getLinesChannel(file io.ReadCloser, lines chan string) <-chan string {
	var line string
	for {
		byteArray := make([]byte, 8)
		_, err := file.Read(byteArray)
		if err != nil {
			file.Close()
			close(lines)
		}
		readLine := string(byteArray)
		if len(strings.Split(readLine, "\n")) == 1 {
			line = line + readLine
		} else {
			for index, value := range strings.Split(readLine, "\n") {
				if index == 0 {
					lines <- line + value
					line = ""
					continue
				}
				line += value
			}
		}
	}
}
