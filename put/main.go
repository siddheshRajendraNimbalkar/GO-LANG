package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("HELLO PUT REQUEST")

	PutUrl := "https://jsonplaceholder.typicode.com/todos/1"

	todo := Todo{
		UserID:    21311,
		Title:     "LET's GO",
		Completed: false,
	}

	todoData, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	req, err := http.NewRequest(http.MethodPut, PutUrl, strings.NewReader(string(todoData)))

	if err != nil {
		fmt.Println("[ERROR]:", err)
		return
	}

	req.Header.Set("Content-type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	fmt.Println("[DATA]:", string(data))
	defer res.Body.Close()
	defer req.Body.Close()
}
