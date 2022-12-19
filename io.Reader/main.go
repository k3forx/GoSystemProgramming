package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	fileInputOutput()
	networkInputOutput()
	endianTransform()
}

func fileInputOutput() {
	file, err := os.Open("file.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}

func networkInputOutput() {
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		panic(err)
	}

	io.WriteString(conn, "GET / HTTP1.0\r\nHost: example.com\r\n\r\n")
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println(res.Header)
	io.Copy(os.Stdout, res.Body)
}

func endianTransform() {
	// 32ビットのビックエンディアンのデータ (10000, 0x2710)
	data := []byte{0x0, 0x0, 0x27, 0x10}
	var i int32
	// エンディアンの変換
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("data: %d\n", i)
}
