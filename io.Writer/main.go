package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	internetAccess()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	// JSON化する元のデータ
	source := map[string]string{
		"Hello": "World",
	}

	compresser := gzip.NewWriter(w)
	multiWriter := io.MultiWriter(compresser, os.Stdout)
	encoder := json.NewEncoder(multiWriter)
	encoder.SetIndent("", "  ")
	encoder.Encode(source)
	compresser.Flush()
}

func internetAccess() {
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET / HTTP1.0\r\nHost: example.com\r\n\r\n")
	io.Copy(os.Stdout, conn)
}
