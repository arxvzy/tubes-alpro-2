package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func getBarang() [10]Barang {
	var raw, err = os.ReadFile("barang.bin")
	if err != nil {
		fmt.Println(err)
	}
	var barang [10]Barang
	json.Unmarshal(raw, &barang)
	return barang
}

func writeBarang(barang [10]Barang) {
	var raw, err = json.Marshal(barang)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(raw)
	err = os.WriteFile("barang.bin", raw, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
