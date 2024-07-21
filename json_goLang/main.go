package main

import (
	"encoding/json"
	"fmt"
)

type School struct {
	Name    string
	Roll_No int8
	Student bool
}

func main() {
	fmt.Println("Hello From json")

	school := School{Name: "sidd", Roll_No: 3, Student: true}
	fmt.Println("[DATA]:", school)

	// json code

	// converting
	jsonData, err := json.Marshal(school)

	if err != nil {
		fmt.Println("[ERROR]:", err)
		return
	}

	Data := string(jsonData)
	fmt.Println("[JSON DATA]:", Data)

	//encoding

	var studentData School
	err = json.Unmarshal(jsonData, &studentData)
	if err != nil {
		fmt.Println("[ERROR]:", err)
		return
	}
	fmt.Println("[ENCODED DATA]:", studentData)
}
