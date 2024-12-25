package main

import "fmt"

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func pauseScreen() {
	fmt.Scanln()
}

func printBarang(data Data) {
	for i := 0; i < data.JumlahBarang; i++ {
		fmt.Printf("%d. %s - Rp. %d - Stok: %d\n", i+1, data.Barang[i].Nama, data.Barang[i].HargaJual, data.Barang[i].Stok)
	}
}

func cariBarangBerdasarkanNama(data Data, nama string) int {
	low := 0
	high := data.JumlahBarang - 1

	for low <= high {
		mid := (low + high) / 2
		if data.Barang[mid].Nama == nama {
			return mid // Barang ditemukan
		} else if data.Barang[mid].Nama < nama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1 // Barang tidak ditemukan
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
