package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("test.csv")

	if err != nil {
		panic(err)
	}

	arr := [][]string{
		{
			"csv",
			"csv",
		},
		{
			"csv",
			"csv",
			"csv",
		}, {
			"csv",
			"csv",
			"csv",
			"csv",
		},
	}

	writer := csv.NewWriter(file)
	writer.Comma = '¥'

	writer.WriteAll(arr)
	writer.Flush()

}
