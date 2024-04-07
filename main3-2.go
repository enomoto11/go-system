// go:build ignore
package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	execute1()
	// assert()
}

func execute1() {
	buffer := make([]byte, 1024)
	rand.Reader.Read(buffer)
	// Readメソッドを使うには一時的にバッファに溜め込む必要がある
	// Readの実態はio.ReadFullだから

	file3_2, err := os.Create("file3_2")
	if err != nil {
		panic(err)
	}

	file3_2.Write(buffer)
	file3_2.Close()
}

func assert() {
	file, err := os.Open("file3_2")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(os.Stdout, file)
}
