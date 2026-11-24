//go:build ignore

package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// for _, link := range links {
	// 	checkLink(link)
	// }

	// for _, link := range links {
	// 	go checkLink(link)
	// }

	channel_go := make(chan string)
	for _, link := range links {
		go checkLinkWithChannel(link, channel_go)
	}

	// fmt.Println(<-channel_go) // blocking code.
	// blocks the main routine. main thing to learn, main routine waits here.

	// fmt.Println(<-channel_go) // waiting twice. will show two in the list.
	// this makes it very critical
	// becos what if the program hangs wiating to recieve
	// so be very thought full when using goroutine channels

	for i := 0; i < len(links); i++ {
		fmt.Println(<-channel_go)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}

func checkLinkWithChannel(link string, channel_go chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
	channel_go <- "done"
}
