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

	lines := getLinesChannel(file)
	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}
}

func getLinesChannel(file io.ReadCloser) <-chan string {
	lines := make(chan string)

	go func() {
		defer file.Close()
		defer close(lines)

		var line string
		for {
			byteArray := make([]byte, 8)
			_, err := file.Read(byteArray)
			if err != nil {
				file.Close()
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
	}()

	return lines
}
