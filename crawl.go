package main

import (
	"fmt"
	"os"
	"net/http"
	// "io/ioutil"
	"github.com/jackdanger/collectlinks"

)

func main() {
	if (len(os.Args) != 2) {
		fmt.Println("Please specific start page")
		os.Exit(1)
	}
	retrieve(os.Args[1])

}

func retrieve(uri string) {
	resp, err := http.Get(uri)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	links := collectlinks.All(resp.Body)

	for _, link := range(links) {
		fmt.Println(link)
	}
// 	body, err := ioutil.ReadAll(resp.Body)
// 	fmt.Println(string(body))
}
