package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Launching server...")
	listener, err := net.Listen("tcp", config.address+":"+config.port)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Server is listening on port %s\n", config.port)

	connection, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
	}
	defer connection.Close()

	for {
		message, err := bufio.NewReader(connection).ReadString(config.delimiter)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break

		}

		var response string
		message = strings.TrimSuffix(message, string(config.delimiter))

		if len(message) > 0 {
			response, err = MultiplyIfInt(message)
			if err != nil {
				response = strings.ToUpper(message)
			}
		} else {
			response = "Message can not be empty."
		}

		_, err = connection.Write([]byte(response + string(config.delimiter)))
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func MultiplyIfInt(input string) (string, error) {
	parsedFloat, err := strconv.ParseFloat(input, 10)
	if err != nil {
		return "", err
	}

	multiplier := 1.0
	isInt := parsedFloat == float64(int(parsedFloat))
	if isInt {
		multiplier = 2.0
	}

	return fmt.Sprintf("%v", parsedFloat*multiplier), nil
}
