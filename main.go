package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var line string
	for {
		byteArray := make([]byte, 8)
		_, err := file.Read(byteArray)
		if err != nil {
			return
		}
		readLine := string(byteArray)
		if len(strings.Split(readLine, "\n")) == 1 {
			line = line + readLine
		} else {
			for index, value := range strings.Split(readLine, "\n") {
				if index == 0 {
					fmt.Printf("read: %s\n", line+value)
					line = ""
					continue
				}
				line += value
			}
		}
	}
}
