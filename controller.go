package main

import (
	"fmt"
)

func lihatBarang(data Data) {
	clearScreen()
	fmt.Println("Menu -> Lihat Barang")
	var inputMenuBarang string
	if data.JumlahBarang == 0 {
		fmt.Println("Barang masih kosong.")
	} else {
		fmt.Println("Data Barang :")
		printBarang(data)
	}
	fmt.Println()
	fmt.Println("1. Urutkan Berdasarkan Nama Barang")
	fmt.Println("2. Urutkan Berdasarkan Harga Barang")
	fmt.Println("3. Cari Berdasarkan Nama")
	fmt.Println("0. Keluar")
	fmt.Print("Pilihan anda : ")
	fmt.Scanln(&inputMenuBarang)
	switch inputMenuBarang {
	case "1":
		urutBerdasarkanNama(&data)
		fmt.Println("Barang berhasil diurutkan berdasarkan nama:")
		printBarang(data)
		pauseScreen()
	case "2":
		urutBerdasarkanHarga(&data)
		fmt.Println("Barang berhasil diurutkan berdasarkan harga:")
		printBarang(data)
		pauseScreen()
	case "3":
		var namaBarang string
		fmt.Print("Masukkan nama barang yang dicari: ")
		fmt.Scanln(&namaBarang)
		urutBerdasarkanNama(&data)
		index := cariBarangBerdasarkanNama(data, namaBarang)
		if index != -1 {
			fmt.Printf("Barang ditemukan:\n\n")
			fmt.Printf("Nama Barang: %s\n", data.Barang[index].Nama)
			fmt.Printf("Harga Jual: %d\n", data.Barang[index].HargaJual)
			fmt.Printf("Stok: %d\n", data.Barang[index].Stok)
			fmt.Printf("Harga Modal: %d\n", data.Barang[index].HargaModal)
		} else {
			fmt.Println("Barang tidak ditemukan")
		}
		pauseScreen()
	case "0":
		return
	default:
		fmt.Println("Menu Tidak Ada")
		pauseScreen()
	}
}

func tambahBarang(data *Data) {
	clearScreen()
	fmt.Println("Menu -> Tambah Barang")
	if data.JumlahBarang >= NBARANG {
		fmt.Println("Penyimpanan barang sudah penuh")
		return
	}
	var nama string
	var harga, modal, stok int
	fmt.Print("Nama barang : ")
	fmt.Scan(&nama)
	fmt.Print("Harga modal : ")
	fmt.Scan(&modal)
	fmt.Print("Harga jual : ")
	fmt.Scan(&harga)
	fmt.Print("Stok barang: ")
	fmt.Scan(&stok)
	var barang = Barang{Nama: nama, HargaJual: harga, Terjual: 0, HargaModal: modal, Stok: stok}
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
	clearScreen()
	fmt.Println("Menu -> Edit Barang")
	var pilihan int
	if data.JumlahBarang == 0 {
		fmt.Println("Barang masih kosong.")
		pauseScreen()
		return
	}

	fmt.Println("Daftar Barang:")
	for i := 0; i < data.JumlahBarang; i++ {
		fmt.Printf("%d. %s - Rp. %d - Modal: Rp. %d - Stok: %d\n", i+1, data.Barang[i].Nama, data.Barang[i].HargaJual, data.Barang[i].HargaModal, data.Barang[i].Stok)
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
	var hargaModalBaru, hargaJualBaru, stokBaru int

	fmt.Printf("Nama baru untuk barang (%s): ", data.Barang[index].Nama)
	fmt.Scan(&namaBaru)
	fmt.Printf("Harga modal baru untuk barang (%d): ", data.Barang[index].HargaModal)
	fmt.Scan(&hargaModalBaru)
	fmt.Printf("Harga jual baru untuk barang (%d): ", data.Barang[index].HargaJual)
	fmt.Scan(&hargaJualBaru)
	fmt.Printf("Stok baru untuk barang (%d): ", data.Barang[index].Stok)
	fmt.Scan(&stokBaru)

	data.Barang[index].Nama = namaBaru
	data.Barang[index].HargaModal = hargaModalBaru
	data.Barang[index].HargaJual = hargaJualBaru
	data.Barang[index].Stok = stokBaru

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

func hapusBarang(data *Data) {
	clearScreen()
	fmt.Println("Menu -> Hapus Barang")
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

func menuTransaksi(data *Data) {
	clearScreen()
	fmt.Println("Menu -> Transaksi")
	var pilihan int
	fmt.Println("Transaksi")
	fmt.Println("1. Beli Barang")
	fmt.Println("2. Lihat Transaksi")
	fmt.Println("3. Hapus Transaksi")
	fmt.Println("0. Keluar")
	fmt.Print("Pilihan anda : ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		beliBarang(data)
	case 2:
		lihatTransaksi(data)
	case 3:
		hapusTransaksi(data)
	case 0:
		return
	default:
		fmt.Println("Menu Tidak Ada")
		pauseScreen()
		pauseScreen()
	}
}

func beliBarang(data *Data) {
	var pilihan, jumlah, index int
	fmt.Println("Beli Barang")
	printBarang(*data)
	fmt.Println("0. Keluar")
	fmt.Print("Masukkan nomor barang yang ingin dibeli: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > data.JumlahBarang {
		fmt.Println("Barang Tidak Ada.")
		pauseScreen()
		return
	}
	fmt.Print("Masukkan jumlah barang yang ingin dibeli: ")
	fmt.Scan(&jumlah)

	index = pilihan - 1

	data.Barang[index].Terjual += jumlah
	var transaksi = Transaksi{TotalHarga: data.Barang[index].HargaJual * jumlah, NamaBarang: data.Barang[index].Nama, JumlahBarang: jumlah}
	data.Transaksi[data.JumlahTransaksi] = transaksi
	data.JumlahTransaksi++

	err := writeData(*data)
	if err != nil {
		fmt.Println("Terjadi kesalahan saat menyimpan data:", err)
	}

	fmt.Println("Transaksi berhasil.")
	fmt.Printf("Nama Barang: %s\n", data.Barang[index].Nama)
	fmt.Printf("Jumlah Barang: %d\n", jumlah)
	fmt.Printf("Total Harga: %d\n", data.Barang[index].HargaJual*jumlah)
	pauseScreen()
	pauseScreen()
}

func lihatTransaksi(data *Data) {
	fmt.Println("Daftar Transaksi:")
	for i := 0; i < data.JumlahTransaksi; i++ {
		fmt.Printf("%d. Nama Barang: %s | Jumlah Terjual: %d  | Total Harga: %d\n", i+1, data.Transaksi[i].NamaBarang, data.Transaksi[i].JumlahBarang, data.Transaksi[i].TotalHarga)
	}
	pauseScreen()
	pauseScreen()
}

func hapusTransaksi(data *Data) {
	if data.JumlahTransaksi == 0 {
		fmt.Println("Tidak ada transaksi yang bisa dihapus")
		return
	}
	fmt.Println("Daftar Transaksi:")
	for i := 0; i < data.JumlahTransaksi; i++ {
		fmt.Printf("%d. Nama Barang: %s | Total Harga: %d\n", i+1, data.Transaksi[i].NamaBarang, data.Transaksi[i].TotalHarga)
	}
	fmt.Print("Masukkan nomor transaksi yang ingin dihapus: ")
	var index int
	fmt.Scan(&index)
	if index < 1 || index > data.JumlahTransaksi {
		fmt.Println("Transaksi Tidak Ada")
		return
	}

	var indexBarang = cariBarangBerdasarkanNama(*data, data.Transaksi[index-1].NamaBarang)

	data.Barang[indexBarang].Terjual -= data.Transaksi[index-1].JumlahBarang
	for i := index - 1; i < data.JumlahTransaksi-1; i++ {
		data.Transaksi[i] = data.Transaksi[i+1]
	}
	data.Transaksi[data.JumlahTransaksi-1] = Transaksi{}
	data.JumlahTransaksi--
	fmt.Println("Transaksi berhasil dihapus")
	pauseScreen()
	pauseScreen()
}

func laporanPenjualan(data Data) {
    // Total Modal
    totalModal := 0
    for i := 0; i < data.JumlahBarang; i++ {
        totalModal += data.Barang[i].HargaModal * data.Barang[i].Terjual
    }

    // Pendapatan Kotor
    pendapatanKotor := 0
    for i := 0; i < data.JumlahTransaksi; i++ {
        pendapatanKotor += data.Transaksi[i].TotalHarga
    }

    // Pendapatan Bersih
    pendapatanBersih := pendapatanKotor - totalModal

    // Urutkan Barang Berdasarkan Jumlah Terjual (Manual Sorting - Bubble Sort)
    for i := 0; i < data.JumlahBarang-1; i++ {
        for j := 0; j < data.JumlahBarang-i-1; j++ {
            if data.Barang[j].Terjual < data.Barang[j+1].Terjual {
                // Tukar posisi barang[j] dengan barang[j+1]
                temp := data.Barang[j]
                data.Barang[j] = data.Barang[j+1]
                data.Barang[j+1] = temp
            }
        }
    }

    // Cetak Laporan
    fmt.Println("Laporan Penjualan:")
    fmt.Printf("Total Modal: Rp %d\n", totalModal)
    fmt.Printf("Pendapatan Kotor: Rp %d\n", pendapatanKotor)
    fmt.Printf("Pendapatan Bersih: Rp %d\n", pendapatanBersih)
    fmt.Println("5 Barang Paling Banyak Terjual:")
    for i := 0; i < 5 && i < data.JumlahBarang; i++ {
        barang := data.Barang[i]
        fmt.Printf("%d. %s - Terjual: %d\n", i+1, barang.Nama, barang.Terjual)
    }

    pauseScreen()
}
