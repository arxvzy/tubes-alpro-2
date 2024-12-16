package main

import (
	"fmt"
)

func lihatBarang(data Data) {
	for i := 0; i < data.Count; i++ {
		fmt.Printf("%d. %s - Rp. %d | %v\n", i+1, data.Barang[i].Nama, data.Barang[i].Harga, data.Barang[i].Kategori)
	}
	pauseScreen()
}

func tambahBarang(data *Data) {
	if data.Count >= NMAX {
		fmt.Println("Barang sudah penuh")
		return
	}
	var nama, kategori string
	var harga int
	fmt.Print("Nama barang : ")
	fmt.Scanln(&nama)
	fmt.Print("Kategori barang : ")
	fmt.Scanln(&kategori)
	fmt.Print("Harga barang : ")
	fmt.Scanln(&harga)
	var barang = Barang{Nama: nama, Harga: harga, Terjual: 0, Kategori: kategori}
	data.Barang[data.Count] = barang
	data.Count++
	var err = writeBarang(*data)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Barang berhasil ditambahkan")
	pauseScreen()
}
