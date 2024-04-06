package main

import (
	"flag"
	"io"
	"os"
)

func main() {
	execute1()
	execute2()
}

func execute1() {
	oldFile, _ := os.Open("old.txt")
	newFile, _ := os.Create("new.txt")

	io.Copy(newFile, oldFile)
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
