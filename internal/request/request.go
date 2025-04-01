package request

import (
	"fmt"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	var request Request
	requestString, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("error while reading request: ", err)
		return nil, err
	}

	requestLineUnparsed := strings.Split(string(requestString), "\r\n")[0]
	if requestLineUnparsed == "" {
		fmt.Println("No Request line found")
		return nil, err
	}
	var requestLine RequestLine
	requestLineSlice := strings.Split(requestLineUnparsed, " ")
	if len(requestLineSlice) != 3 {
		return nil, fmt.Errorf("bad request, request = %s", requestLineUnparsed)
	}
	requestLine.Method = requestLineSlice[0]
	requestLine.RequestTarget = requestLineSlice[1]
	requestLine.HttpVersion = strings.Split(requestLineSlice[2], "/")[1]

	request.RequestLine = requestLine

	return &request, nil
}
