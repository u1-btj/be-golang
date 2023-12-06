package main

import (
    "fmt"
	"strconv"
)

// Inventaris awal kosong
type item map[string]string
var inventaris []item

// Fungsi untuk menambahkan produk ke inventaris
func tambah_produk(nama string, harga string, jumlah string){
	new_item := make(item)
	new_item["nama"] = nama
	new_item["harga"] = harga
	new_item["jumlah"] = jumlah
    inventaris = append(inventaris, new_item)
}

// Fungsi untuk memperbarui jumlah produk
func perbarui_jumlah(nama_produk string, jumlah_baru string){
	for _, item := range inventaris { 
		if nama_produk == item["nama"] {
			item["jumlah"] = jumlah_baru 
		}
	}
}

// Fungsi untuk menampilkan inventaris saat ini
func tampilkan_inventaris(){
    fmt.Println("Inventaris Saat Ini:")
    fmt.Println("No || Nama || Harga || Jumlah")
    for i, item := range inventaris { 
		fmt.Println(i + 1, item["nama"], item["harga"], item["jumlah"])
	}
}

// Fungsi untuk menghitung total nilai inventaris
func hitung_total_nilai(){
    var total_nilai int64 = 0
	tampilkan_inventaris()
	for _, item := range inventaris {
		harga, _ := strconv.ParseInt(item["harga"], 0, 64)
		jumlah, _ := strconv.ParseInt(item["jumlah"], 0, 64)
		total_nilai += harga * jumlah
	}
    fmt.Println("Total nilai inventaris: ", total_nilai)
}

func main() {
	tambah_produk("Sandal", "50000", "5")
	tambah_produk("Sepatu", "150000", "10")
	tambah_produk("Tas", "100000", "5")
	tampilkan_inventaris()

	perbarui_jumlah("Tas", "15")
	tampilkan_inventaris()

	hitung_total_nilai()
}