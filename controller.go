package main

import (
	"fmt"
)

func lihatBarang(data Data) {
	var inputMenuBarang string
	if data.JumlahBarang == 0 {
		fmt.Println("Belum ada barang")
	} else {
		for i := 0; i < data.JumlahBarang; i++ {
			fmt.Printf("%d. %s - Rp. %d | %v\n", i+1, data.Barang[i].Nama, data.Barang[i].HargaJual, data.Barang[i].Kategori)
		}
	}
	fmt.Println("1. Urutkan Berdasarkan Nama Barang")
	fmt.Println("2. Urutkan Berdasarkan Harga Barang")
	fmt.Println("0. Keluar")
	fmt.Print("Pilihan anda : ")
	fmt.Scan(&inputMenuBarang)
	switch inputMenuBarang {
	case "1":
	case "2":
	case "0":
		return
	default:
		fmt.Println("Menu Tidak Ada")
	}
	pauseScreen()
}

func tambahBarang(data *Data) {
	if data.JumlahBarang >= NBARANG {
		fmt.Println("Barang sudah penuh")
		return
	}
	var nama, kategori string
	var harga, modal int
	fmt.Print("Nama barang : ")
	fmt.Scan(&nama)
	fmt.Print("Kategori barang : ")
	fmt.Scan(&kategori)
	fmt.Print("Harga modal : ")
	fmt.Scan(&modal)
	fmt.Print("Harga jual : ")
	fmt.Scan(&harga)
	var barang = Barang{Nama: nama, HargaJual: harga, Terjual: 0, Kategori: kategori, HargaModal: modal}
	data.Barang[data.JumlahBarang] = barang
	data.JumlahBarang++
	var err = writeData(*data)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Barang berhasil ditambahkan")
	pauseScreen()
}
