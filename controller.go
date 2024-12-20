package main

import (
	"fmt"
)

func lihatBarang(data Data) {
	var inputMenuBarang string
	if data.JumlahBarang == 0 {
		fmt.Println("Belum ada barang")
	} else {
		fmt.Println("Data Barang :")
		printBarang(data)
	}
	fmt.Println()
	fmt.Println("1. Urutkan Berdasarkan Nama Barang")
	fmt.Println("2. Urutkan Berdasarkan Harga Barang")
	fmt.Println("0. Keluar")
	fmt.Print("Pilihan anda : ")
	fmt.Scanln(&inputMenuBarang)
	switch inputMenuBarang {
	case "1":
		urutBerdasarkanNama(data)
		pauseScreen()
	case "2":
		urutBerdasarkanHarga(data)
		pauseScreen()
	default:
		fmt.Println("Menu Tidak Ada")
		pauseScreen()
	}
}

func tambahBarang(data *Data) {
	if data.JumlahBarang >= NBARANG {
		fmt.Println("Barang sudah penuh")
		return
	}
	var nama string
	var harga, modal int
	fmt.Print("Nama barang : ")
	fmt.Scan(&nama)
	fmt.Print("Harga modal : ")
	fmt.Scan(&modal)
	fmt.Print("Harga jual : ")
	fmt.Scan(&harga)
	var barang = Barang{Nama: nama, HargaJual: harga, Terjual: 0, HargaModal: modal}
	data.Barang[data.JumlahBarang] = barang
	data.JumlahBarang++
	var err = writeData(*data)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Barang berhasil ditambahkan")
	pauseScreen()
	pauseScreen()
}

func editBarang(data *Data) {
	var pilihan int
	if data.JumlahBarang == 0 {
		fmt.Println("Belum ada barang untuk diubah.")
		pauseScreen()
		return
	}

	fmt.Println("Daftar Barang:")
	for i := 0; i < data.JumlahBarang; i++ {
		fmt.Printf("%d. %s - Rp. %d - Modal: Rp. %d\n", i+1, data.Barang[i].Nama, data.Barang[i].HargaJual, data.Barang[i].HargaModal)
	}

	fmt.Print("Masukkan nomor barang yang ingin diubah: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > data.JumlahBarang {
		fmt.Println("Barang Tidak Ada.")
		pauseScreen()
		return
	}

	index := pilihan - 1

	var namaBaru string
	var hargaModalBaru, hargaJualBaru int

	fmt.Printf("Nama baru untuk barang (%s): ", data.Barang[index].Nama)
	fmt.Scan(&namaBaru)
	fmt.Printf("Harga modal baru untuk barang (%d): ", data.Barang[index].HargaModal)
	fmt.Scan(&hargaModalBaru)
	fmt.Printf("Harga jual baru untuk barang (%d): ", data.Barang[index].HargaJual)
	fmt.Scan(&hargaJualBaru)

	data.Barang[index].Nama = namaBaru
	data.Barang[index].HargaModal = hargaModalBaru
	data.Barang[index].HargaJual = hargaJualBaru

	var err = writeData(*data)
	if err != nil {
		fmt.Println("Terjadi kesalahan saat menyimpan data:", err)
		pauseScreen()
		pauseScreen()
	} else {
		fmt.Println("Barang berhasil diperbarui.")
		pauseScreen()
		pauseScreen()
	}
}

func urutBerdasarkanNama(data Data) {
	for i := 0; i < data.JumlahBarang-1; i++ {
		minIdx := i
		for j := i + 1; j < data.JumlahBarang; j++ {
			if data.Barang[j].Nama < data.Barang[minIdx].Nama {
				minIdx = j
			}
		}
		data.Barang[i], data.Barang[minIdx] = data.Barang[minIdx], data.Barang[i]
	}
	fmt.Println("Barang berhasil diurutkan berdasarkan nama:")
	printBarang(data)
}

func urutBerdasarkanHarga(data Data) {
	for i := 1; i < data.JumlahBarang; i++ {
		key := data.Barang[i]
		j := i - 1
		// Geser elemen-elemen yang lebih besar dari key
		for j >= 0 && data.Barang[j].HargaJual > key.HargaJual {
			data.Barang[j+1] = data.Barang[j]
			j--
		}
		data.Barang[j+1] = key
	}
	fmt.Println("Barang berhasil diurutkan berdasarkan harga:")
	printBarang(data)
}

func hapusBarang(data *Data) {
	if data.JumlahBarang == 0 {
		fmt.Println("Tidak ada barang yang bisa dihapus")
		return
	}

	printBarang(*data)

	// Meminta input nomor barang yang ingin dihapus
	var index int
	fmt.Print("Masukkan nomor barang yang ingin dihapus: ")
	fmt.Scan(&index)

	// Validasi input
	if index < 1 || index > data.JumlahBarang {
		fmt.Println("Barang Tidak Ada")
		return
	}

	// Menghapus barang dengan menggeser elemen array
	for i := index - 1; i < data.JumlahBarang-1; i++ {
		data.Barang[i] = data.Barang[i+1]
	}
	data.Barang[data.JumlahBarang-1] = Barang{} // Kosongkan elemen terakhir
	data.JumlahBarang--

	// Simpan perubahan ke file
	err := writeData(*data)
	if err != nil {
		fmt.Println("Gagal menyimpan data:", err)
		return
	}

	fmt.Println("Barang berhasil dihapus")
	pauseScreen()
	pauseScreen()

}

func printBarang(data Data) {
	for i := 0; i < data.JumlahBarang; i++ {
		fmt.Printf("%d. %s - Rp. %d\n", i+1, data.Barang[i].Nama, data.Barang[i].HargaJual)
	}
}
