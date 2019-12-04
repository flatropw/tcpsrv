package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Launching server...")
	listener, err := net.Listen("tcp", address+":"+port)
	if err != nil {
		panic(err)
	}

	connection, err := listener.Accept()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Server is listening on port %s", port)

	for {
		message, err := bufio.NewReader(connection).ReadString(delimiter)
		if err != nil {
			panic(err)
		}

		_, err = connection.Write([]byte(message + string(delimiter)))
		if err != nil {
			panic(err)
		}
	}
}
