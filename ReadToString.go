package main

import (
	"fmt"
	"io"
	"os"
)

func ReadToString(fileName string) (Text string) {
	File, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error Opening: " + fileName)
	}
	defer File.Close()
	bytes, err := io.ReadAll(File)
	if err != nil {
		fmt.Println("Error Reading: " + fileName)
	}

	return string(bytes)
}
