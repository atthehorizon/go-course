package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]
	fmt.Println("filename:", filename)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}
