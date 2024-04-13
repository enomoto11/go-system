// go:build ignore
package main

import (
	"archive/zip"
	"strings"
)

func main() {
	execute()
}

// io.Writerって直接インスタンス化できないのか？
func execute() {
	text := "圧縮される内容"
	reader := strings.NewReader(text)

	// buffer := make([]byte, 1024)
	// reader.Read(buffer)

	// TODO: strings.Reader -> io.Writerへの変換

	zipWriter := zip.NewWriter()
	defer zipWriter.Close()
}
