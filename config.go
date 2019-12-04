package main

type configList struct {
	address string
	port string
	delimiter byte
}

var config = configList{
	address: "127.0.0.1",
	port: "8081",
	delimiter: '\n',
}
