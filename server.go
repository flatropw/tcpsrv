package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Launching server...")
	listener, err := net.Listen("tcp", config.address+":"+config.port)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Server is listening on port %s", config.port)

	connection, err := listener.Accept()
	if err != nil {
		panic(err)
	}


	for {
		message, err := bufio.NewReader(connection).ReadString(config.delimiter)
		if err != nil {
			panic(err)
		}

		_, err = connection.Write([]byte(message + "\n"))
		if err != nil {
			panic(err)
		}
	}
}
