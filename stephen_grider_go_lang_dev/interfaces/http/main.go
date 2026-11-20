package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// 1. why 99999 becos resp.Body.Read function will say that it is full stop ...
	// okay so if i have it 5 it should only have 5 chars....

	// bs := make([]byte, 5)
	// resp.Body.Read(bs) // benefit of Read interface so we dont need to create multiple functions and code reuse ..
	// like ReadBody .. ReadFile ... etc ... code reuse ...

	// fmt.Println(bs)
	// fmt.Println(string(bs))

	io.Copy(os.Stdout, resp.Body)

}
