package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	// json化する元のデータ
	source := map[string]string{
		"Hello": "World",
	}

	// ここにコードを書く
	jsonSource, err := json.Marshal(source)
	if err != nil {
		fmt.Println("JSON marshal error: ", err)
	}

	// gzip-------------------
	gzipWriter := gzip.NewWriter(w)
	gzipWriter.Header.Name = "test.txt"
	// gzip-------------------

	// file--------------------
	file, _ := os.Create("test2.txt")
	defer file.Close()
	// file--------------------

	writer := io.MultiWriter(gzipWriter, os.Stdout, file)
	io.WriteString(writer, string(jsonSource))
	gzipWriter.Flush()

}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
