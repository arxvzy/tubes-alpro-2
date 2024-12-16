package main

import (
	"fmt"
)

func lihatBarang() {
	var result [NMAX]Barang = getBarang()
	for i := 0; i < NMAX; i++ {
		fmt.Printf("%d. %s - Rp. %d | %v\n", i+1, result[i].Nama, result[i].Harga, result[i].Kategori)
	}
}

func tambahBarang() {
	var nama, kategori string
	var harga, stok int
	fmt.Print("Nama barang : ")
	fmt.Scanln(&nama)
	fmt.Print("Kategori barang : ")
	fmt.Scanln(&kategori)
	fmt.Print("Harga barang : ")
	fmt.Scanln(&harga)
	fmt.Print("Stok barang : ")
	fmt.Scanln(&stok)
	var barang = Barang{Nama: nama, Harga: harga, Stok: stok, Terjual: 0, Kategori: kategori}
}
