package main

import (
	"io"
	"net"
	"os"
)

func main() {
	internetAccess()
}

func internetAccess() {
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET / HTTP1.0\r\nHost: example.com\r\n\r\n")
	io.Copy(os.Stdout, conn)
}
