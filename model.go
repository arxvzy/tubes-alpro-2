package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func getData() Data {
	var raw, err = os.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
	}
	var data Data
	json.Unmarshal(raw, &data)
	return data
}

func writeBarang(data Data) error {
	var raw, err = json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile("data.json", raw, 0644)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
