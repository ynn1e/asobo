/*
https://pkg.go.dev/net
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

const (
	network = "tcp"
	address = "localhost:8080"
)

func main() {
	listener, err := net.Listen(network, address)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("[listen] network=%s, address=%s\n", network, address)

	// If you need, SetDeadline, SetReadDeadline, SetWriteDeadline.
	conn, err := listener.Accept()
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	log.Printf("[accept] remote_addr=%v, local_addr=%v\n", conn.RemoteAddr(), conn.LocalAddr())

	req, err := http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		log.Panicln(err)
	}

	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Panicln(err)
	}

	log.Println("--- read response ---")
	fmt.Println(string(dump))

	resp := http.Response{
		StatusCode: http.StatusOK,
		ProtoMajor: 1,
		ProtoMinor: 0,
		Body:       io.NopCloser(strings.NewReader("Hello.")),
	}

	if err := resp.Write(conn); err != nil {
		log.Panic(err)
	}
}
