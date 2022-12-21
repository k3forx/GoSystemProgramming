package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")

	for {
		fmt.Println("before accept")
		conn, err := listener.Accept()
		fmt.Println("accept is called")
		if err != nil {
			panic(err)
		}
		fmt.Println("conn is created")
		go func() {
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			// リクエストを読み込む
			request, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}

			dump, err := httputil.DumpRequest(request, true)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(dump))

			response := http.Response{
				StatusCode: http.StatusOK,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body: io.NopCloser(
					strings.NewReader("Hello World\n"),
				),
			}
			response.Write(conn)
			conn.Close()
		}()
		fmt.Println("aaa")
	}
}
