package main

import (
	"fmt"
)

const NMAX int = 10

type Barang struct {
	Nama     string `json:"nama"`
	Harga    int    `json:"harga"`
	Terjual  int    `json:"terjual"`
	Kategori string `json:"kategori"`
}

type Transaksi struct {
	Harga int `json:"harga"`
}

type Data struct {
	Count  int            `json:"count"`
	Barang [NMAX]Barang   `json:"barang"`
	Harga  [100]Transaksi `json:"harga"`
}

func main() {
	var data Data = getData()
	var inputMenu string
	for inputMenu != "0" {
		clearScreen()
		fmt.Println("Aplikasi Penjualan")
		fmt.Println("1. Lihat Barang") // ada menu cari, urut berdasarkan ..., kembali ke menu utama
		fmt.Println("2. Tambah Barang")
		fmt.Println("3. Edit Barang")
		fmt.Println("4. Hapus Barang")
		fmt.Println("5. Transaksi")         // tambah, edit, hapus transaksi, kembali ke menu utama
		fmt.Println("6. Laporan Penjualan") // top 5 penjualan, data modal, pendapatan, dll
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan anda : ")
		fmt.Scanln(&inputMenu)

		switch inputMenu {
		case "1":
			lihatBarang(data)
		case "2":
			tambahBarang(&data)
		}
	}
}

// harga modal pada input barang,
