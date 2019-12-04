package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", config.address+":"+config.port)
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Request:")

		text, err := reader.ReadString(config.delimiter)
		if err != nil {
			panic(err)
		}

		_, err = fmt.Fprintf(conn, text+"\n")
		if err != nil {
			panic(err)
		}

		message, err := bufio.NewReader(conn).ReadString(config.delimiter)
		if err != nil {
			panic(err)
		}

		fmt.Print("Response: " + message)
	}
}
