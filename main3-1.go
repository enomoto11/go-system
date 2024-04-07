// go:build ignore

package main

import (
	"flag"
	"io"
	"os"
)

func main() {
	// prepare()
	execute1()
	// execute2()
}

func prepare() {
	oldFile, err := os.Create("old.txt")
	// oldFile, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	// defer oldFile.Close()

	oldFile.Write([]byte("oldファイルの内容です\n"))
	oldFile.Close()

}

func execute1() {
	oldFile, _ := os.Open("old.txt")

	byte, err := io.ReadAll(oldFile)
	if err != nil {
		panic(err)
	}
	println("byte: ", byte)

	newFile, _ := os.Create("new.txt")

	io.WriteString(newFile, string(byte))

	// io.Copy(newFile, oldFile)
}

func execute2() {
	flag.Parse()
	fileName := flag.Arg(0)

	copiedFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	newFile, _ := os.Create("new.txt")
	io.Copy(newFile, copiedFile)
}
