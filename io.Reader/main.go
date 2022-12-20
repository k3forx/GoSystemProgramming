package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	fmt.Println("-----------------------------------")
	fileInputOutput()

	fmt.Println("-----------------------------------")
	networkInputOutput()

	fmt.Println("-----------------------------------")
	endianTransform()

	fmt.Println("-----------------------------------")
	file, err := os.Open("PNG_transparency_demonstration_1.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	chunks := readChunks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}

	fmt.Println("-----------------------------------")
	newFile, err := os.Create("PNG_transparency_demonstration_secret.png")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	chunks = readChunks(file)
	// シグニチャを書き込み
	io.WriteString(newFile, "\x89PNG\r\n\x1a\n")
	// 先頭に必要なIHDRチャンクを書き込み
	io.Copy(newFile, chunks[0])
	// テキストチャンクを追加
	io.Copy(newFile, textChunk("Lambda Note++"))
	// 残りのチャンクを追加
	for _, chunk := range chunks[1:] {
		io.Copy(newFile, chunk)
	}

	newChunks := readChunks(newFile)
	for _, chunk := range newChunks {
		dumpChunk(chunk)
	}

	fmt.Println("-----------------------------------")
	multiReader()
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

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)

	if bytes.Equal(buffer, []byte("teXt")) {
		rawText := make([]byte, length)
		chunk.Read(rawText)
		fmt.Println(string(rawText))
	}
}

func readChunks(file *os.File) []io.Reader {
	var chunks []io.Reader

	// 最初の8バイトを飛ばす
	file.Seek(8, 0)
	var offset int64 = 8

	for {
		var length int32
		// 32ビット (4byte) 分だけPNGのバイナリデータを読み込む
		// PNGのバイナリデータの形式から最初の4byteはデータの長さの情報になっている
		if err := binary.Read(file, binary.BigEndian, &length); err == io.EOF {
			break
		}
		chunks = append(chunks,
			io.NewSectionReader(file, offset, int64(length)+12))
		// 次のチャンクの先頭に移動
		// 現在位置は長さを読み終わった箇所なので
		// チャンク名 (4バイト) + データ長 + CRC (4バイト) 先に移動
		offset, _ = file.Seek(int64(length+8), 1)
	}
	return chunks
}

func textChunk(text string) io.Reader {
	byteText := []byte(text)
	crc := crc32.NewIEEE()
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, int32(len(byteText)))
	// CRC計量とバッファへの書き込みを同時に行うMultiWriter
	writer := io.MultiWriter(&buffer, crc)
	io.WriteString(writer, "teXt") //2バイトめの5ビットめを立てる (小文字にする) とプライベート

	writer.Write(byteText)
	binary.Write(&buffer, binary.BigEndian, crc.Sum32())
	return &buffer
}

func multiReader() {
	header := bytes.NewBufferString("----- HEADER -----\n")
	content := bytes.NewBufferString("Example of io.MultiReader\n")
	footer := bytes.NewBufferString("----- FOOTER -----\n")

	reader := io.MultiReader(header, content, footer)
	// すべてのreaderをつなげた出力が表示
	io.Copy(os.Stdout, reader)
}
