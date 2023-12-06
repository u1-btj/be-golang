# Inventaris awal kosong
inventaris = []

# Fungsi untuk menambahkan produk ke inventaris
def tambah_produk(nama, harga, jumlah):
    inventaris.append({'nama': nama, 'harga': harga, 'jumlah': jumlah})

# Fungsi untuk memperbarui jumlah produk
def perbarui_jumlah(nama_produk, jumlah_baru):
    for i in inventaris:
        if i['nama'] == nama_produk:
            i['jumlah'] = jumlah_baru

# Fungsi untuk menampilkan inventaris saat ini
def tampilkan_inventaris():
    print("Inventaris Saat Ini:")
    for i in inventaris:
        print(i)

# Fungsi untuk menghitung total nilai inventaris
def hitung_total_nilai():
    total_nilai = 0
    # proses untuk menghitung total nilai
    for i in inventaris:
        total_nilai += i['harga'] * i['jumlah']
    print("Total nilai inventaris: ", total_nilai)

tambah_produk('sepatu', 50000, 2)
tambah_produk('sandal', 15000, 5)
tambah_produk('tas', 25000, 15)
perbarui_jumlah('sepatu', 10)
tampilkan_inventaris()
hitung_total_nilai()