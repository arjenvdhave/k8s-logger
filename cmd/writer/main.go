package main

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/ugorji/go/codec"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

var (
	v  interface{} // value to decode/encode into
	r  io.Reader
	w  io.Writer
	b  []byte
	mh codec.MsgpackHandle
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	var dec = codec.NewDecoder(conn, new(codec.MsgpackHandle))
	var m interface{}
	var e = dec.Decode(&m)
	if e != nil {
		// Send a response back to person contacting us.
		conn.Write([]byte("Error"))
	} else {
		fmt.Print("Received: ")
		fmt.Print(m)
		// Send a response back to person contacting us.
		conn.Write([]byte("OK"))
	}
	conn.Close()
}
