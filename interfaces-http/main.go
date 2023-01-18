package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// resp.Body.Close()
	// fmt.Println(string(bs))

	// body, err := io.ReadAll(resp.Body)
	// resp.Body.Close()
	// if resp.StatusCode > 299 {
	// 	fmt.Printf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	// }
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("%s", body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	bw, err := fmt.Println(string(bs))
	fmt.Println("wrote", bw, "bytes")
	return bw, err
}
