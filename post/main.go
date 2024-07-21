package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int8   `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("HELLO Post")
	postUrl := "https://jsonplaceholder.typicode.com/todos"

	todo := Todo{
		UserID:    1,
		Title:     "HeLLO THERE",
		Completed: true,
	}

	todoData, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("[ERROR]:", err)
		return
	}

	jsonData := strings.NewReader(string(todoData))

	res, err := http.Post(postUrl, "createing user todo", jsonData)

	if err != nil {
		fmt.Println("[ERROR]:", err)
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("[ERROR]:", err)
		return
	}

	fmt.Println("[Data]:", string(data))
	defer res.Body.Close()
}
