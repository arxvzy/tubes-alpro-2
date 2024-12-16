package main

import (
	"fmt"
)

const NMAX int = 10

type Barang struct {
	Nama     string `json:"nama"`
	Harga    int    `json:"harga"`
	Stok     int    `json:"stok"`
	Terjual  int    `json:"terjual"`
	Kategori string `json:"kategori"`
}

func main() {

	var barang [NMAX]Barang
	var inputMenu string

	for inputMenu != "0" {
		clearScreen()
		fmt.Println("Aplikasi Penjualan")
		fmt.Println("1. Lihat Barang") // ada menu cari, urut berdasarkan ..., kembali ke menu utama
		fmt.Println("2. Tambah Barang")
		fmt.Println("3. Edit Barang")
		fmt.Println("4. Hapus Barang")
		fmt.Println("5. Kelola Transaksi Penjualan") // tambah, edit, hapus transaksi, kembali ke menu utama
		fmt.Println("6. Laporan Penjualan")          // top 5 penjualan, data modal, pendapatan, dll
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan anda : ")
		fmt.Scanln(&inputMenu)

		switch inputMenu {
		case "1":
			lihatBarang()
		case "2":
			writeBarang(barang)
		}
	}
}
