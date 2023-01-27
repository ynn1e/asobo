package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
)

const (
	network = "tcp"
	scheme  = "http://"
	address = "localhost:8080"
)

func main() {
	conn, err := net.Dial(network, address)
	if err != nil {
		log.Panic(err)
	}

	req, err := http.NewRequest(http.MethodGet, scheme+address, nil)
	if err != nil {
		log.Panic(err)
	}

	if err := req.Write(conn); err != nil {
		log.Panic(err)
	}

	resp, err := http.ReadResponse(bufio.NewReader(conn), req)
	if err != nil {
		log.Panic(err)
	}

	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(dump))
}