package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	Name string
	Age  int16
}

func main() {
	fmt.Println("Welcome to JSON!")

	EncodeJson()
}

func EncodeJson() {
	persons := []person{
		{"Hella", 21},
		{"Thor", 1500},
	}

	jsonBytes, _ := json.Marshal(persons)

	writeJsonFile("data", jsonBytes)
}

func writeJsonFile(name string, data []byte) {
	os.WriteFile(name+".json", data, os.ModePerm)
}
