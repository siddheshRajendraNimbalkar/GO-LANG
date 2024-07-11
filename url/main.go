package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("URL")
	myUrl := "https://jsonplaceholder.typicode.com/todos/1?key=value"
	fmt.Println("[URL]: ", myUrl)

	parserUrl, err := url.Parse(myUrl)

	if err != nil {
		fmt.Print("ERROR", err)
		return
	}

	fmt.Println(parserUrl)
	fmt.Println("[Scheme]:", parserUrl.Scheme)
	fmt.Println("[Host]:", parserUrl.Host)
	fmt.Println("[Path]:", parserUrl.Path)
	fmt.Println("[Query]:", parserUrl.RawQuery)
}
