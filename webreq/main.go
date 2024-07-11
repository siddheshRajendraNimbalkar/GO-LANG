package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Hello!!!!")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("[ERROR]:", err)
		return
	}
	println("[RESPONSE]:", res)

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("[ERROR]:", err)
		return
	}

	fmt.Println("[DATA]: ", string(data))

	defer res.Body.Close()
}
