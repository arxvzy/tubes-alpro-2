package main

import (
	"fmt"
)

const NBARANG int = 10
const NTRANSAKSI int = 10

type Barang struct {
	Nama       string `json:"nama"`
	HargaJual  int    `json:"hargaJual"`
	Terjual    int    `json:"terjual"`
	HargaModal int    `json:"hargaModal"`
}

type Transaksi struct {
	TotalHarga   int    `json:"totalHarga"`
	NamaBarang   string `json:"namaBarang"`
	JumlahBarang int    `json:"jumlahBarang"`
}

type Data struct {
	JumlahBarang    int                   `json:"jumlahBarang"`
	JumlahTransaksi int                   `json:"jumlahTransaksi"`
	Barang          [NBARANG]Barang       `json:"barang"`
	Transaksi       [NTRANSAKSI]Transaksi `json:"transaksi"`
}

func main() {
	var data Data = getData()
	var inputMenu string
	for inputMenu != "0" {
		clearScreen()
		fmt.Println("Aplikasi Penjualan")
		fmt.Println("1. Lihat Barang")
		fmt.Println("2. Tambah Barang")
		fmt.Println("3. Edit Barang")
		fmt.Println("4. Hapus Barang")
		fmt.Println("5. Transaksi")
		fmt.Println("6. Laporan Penjualan")
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan anda : ")
		fmt.Scanln(&inputMenu)

		switch inputMenu {
		case "1":
			lihatBarang(data)
		case "2":
			tambahBarang(&data)
		case "3":
			editBarang(&data)
		case "4":
			hapusBarang(&data)
		case "5":
			menuTransaksi(&data)
		case "6":
		case "0":
			return
		default:
			fmt.Println("Menu Tidak Ada")
		}

	}
}

// BUG:
// 1. MULTI WORD INPUT SCRAMBLE THE INPUT IN tambahBarang()
