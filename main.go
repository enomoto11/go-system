package main

import (
	"io"
	"os"
)

func main() {
	// prepare()
	execute1()
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
