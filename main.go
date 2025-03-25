package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for {
		byteArray := make([]byte, 8)
		_, err := file.Read(byteArray)
		if err != nil {
			return
		}
		fmt.Printf("%s\n", string(byteArray))
	}
}
