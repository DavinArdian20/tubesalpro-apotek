package main

import "fmt"

type Obat struct {
	ID, nama, gejala string
	Stok             int
	Harga            float64
	Kadaluarsa       string
}

const nmax int = 1000

type dataObat [nmax]Obat

func seedData(a *dataObat, stokObat *int) {
	// Data dummy untuk mengisi array obat
	a[0] = Obat{"OB001", "Paracetamol", "Demam", 100, 5000.0, "2025-12-31"}
	a[1] = Obat{"OB002", "Amoxicillin", "Infeksi", 50, 15000.0, "2024-06-30"}
	a[2] = Obat{"OB003", "Loperamide", "Diare", 200, 8000.0, "2026-03-31"}
	a[3] = Obat{"OB004", "Cetirizine", "Alergi", 75, 12000.0, "2025-09-30"}
	a[4] = Obat{"OB005", "Ibuprofen", "Nyeri", 150, 10000.0, "2024-12-31"}
	*stokObat = 5
}
func main() {
	var a dataObat
	var stokObat int
	var pilihan int
	var done bool
	seedData(&a, &stokObat)
	for done = false; !done; {
		cetakMenu()
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			tambahObat(&a, &stokObat)
		case 2:
			editData(&a, stokObat)
		case 3:
			hapusData(&a, &stokObat)
		case 4:
			sortingData(&a, stokObat)
		case 5:
			cariData(&a, stokObat)
		case 6:
			cetakData(a, stokObat)
		case 7:
			cetakStatistik(a, stokObat)
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi apotek.")
			done = true
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}

}

func cetakMenu() {
	fmt.Println("=== Aplikasi Apotek ===")
	fmt.Println("1. Tambah Obat")
	fmt.Println("2. Edit Data Obat")
	fmt.Println("3. Hapus Data Obat")
	fmt.Println("4. Sorting Data Obat")
	fmt.Println("5. Cari Data Obat")
	fmt.Println("6. Cetak Data Obat")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih menu: ")
}

func tambahObat(a *dataObat, stokObat *int) {
	var i, newObat int
	var namaObat string // Asumsi nama obat unik
	fmt.Print("Masukkan jumlah obat yang ingin ditambahkan: ")
	fmt.Scan(&newObat)
	// Cek apakah jumlah obat melebihi kapasitas array
	if newObat+*stokObat > nmax {
		fmt.Println("Jumlah obat melebihi kapasitas maksimum.")
		return
	}
	var tempIndex int
	for tempIndex < newObat {
		fmt.Print("Masukkan nama obat: ")
		fmt.Scan(&namaObat)
		// Memanggil fungsi cekNamaObat untuk mengecek jika namaObat yang diinput sudah ada di dalam data.
		if !cekNamaObat(*a, namaObat, *stokObat) {
			fmt.Printf("Nama obat '%s' sudah ada. Silakan masukkan nama lain.\n", namaObat)
		} else {
			i = *stokObat
			a[i].nama = namaObat
			// Memanggil fungsi idObat untuk membuat id obat secara otomatis.
			a[i].ID = idObat(i + 1)
			fmt.Print("Masukkan gejala yang dapat diobati: ")
			fmt.Scan(&a[i].gejala)
			fmt.Print("Masukkan stok obat: ")
			fmt.Scan(&a[i].Stok)
			fmt.Print("Masukkan harga obat: ")
			fmt.Scan(&a[i].Harga)
			fmt.Print("Masukkan tanggal kadaluarsa (YYYY-MM-DD): ")
			fmt.Scan(&a[i].Kadaluarsa)
			*stokObat++
			tempIndex++
		}
	}

}

func cekNamaObat(a dataObat, namaObat string, stokObat int) bool {
	// menerima nama obat dan mengecek apakah nama obat sudah ada dalam data obat
	// Dengan sequential search
	var i int
	for i = 0; i < stokObat; i++ {
		if a[i].nama == namaObat {
			return false
		}
	}
	return true
}
func idObat(index int) string {
	// Generate ID obat berdasarkan index
	var id string
	// Inisialisasi semua ID obat dengan "OB"
	id = "OB"
	if index < 10 {
		id += "00"
	} else if index < 100 {
		id += "0"
	}
	var tempNum, digit int
	var tempString string
	tempNum = index
	// While loop untuk mengambil digit terakhir dari index lalu diubah ke character dan diubah lagi ke string
	for tempNum > 0 {
		digit = tempNum % 10
		tempString = string('0'+rune(digit)) + tempString
		tempNum /= 10
	}
	return id + tempString
}
func editData(a *dataObat, stokObat int) {
	// Fungsi untuk mengedit data dari obat
	var piilihan, idx int
	var namaBaru, cariNama string
	var done bool
	fmt.Print("Pilih kategori yang ingin diubah:")
	fmt.Scan(&piilihan)
	fmt.Print("Cari nama obat yang ingin diedit datanya:")
	fmt.Scan(&cariNama)
	idx = cariNamaObat(a, stokObat, cariNama)
	switch piilihan {
	case 1:
		for done = false; !done; {
			fmt.Scan(&namaBaru)
			if !cekNamaObat(a, namaBaru, stokObat) {
				fmt.Printf("Nama obat '%s' sudah ada. Silakan masukkan nama lain.\n", namaBaru)
			} else {
				a[idx].nama = namaBaru
				done = true
			}
		}
	case 2:
		fmt.Print("Masukkan gejala baru yang dapat diobati: ")
		fmt.Scan(&a[idx].gejala)
	case 3:
		fmt.Print("Masukkan stok baru obat: ")
		fmt.Scan(&a[idx].Stok)
	case 4:
		fmt.Print("Masukkan harga baru obat: ")
		fmt.Scan(&a[idx].Harga)
	case 5:
		fmt.Print("Masukkan tanggal kadaluarsa baru (YYYY-MM-DD): ")
		fmt.Scan(&a[idx].Kadaluarsa)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}
func hapusData(a *dataObat, stokObat *int) {
	var i, idx int
	var target string
	fmt.Print("Masukkan nama obat yang ingin dihapus:")
	fmt.Scan(&target)
	idx = cariNamaObat(*a, *stokObat, target)
	if idx < 0 {
		fmt.Println("Nama obat yang dimasukkan tidak ada dalam data.")
		return
	}
	for i = idx; i < *stokObat-1; i++ {
		a[i] = a[i+1]
	}
	*stokObat -= 1
}
func sortingData(a *dataObat, stokObat int) {
	// Fungsi untuk mengurutkan data dengan memilih mengurut naik atau turun.
	var pilihan int
	fmt.Print("Pilih ... (1/2):")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		sortKadalauarsaAsc(a, stokObat)
	} else if pilihan == 2 {
		sortKadalauarsaDesc(a, stokObat)
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}
func cariData(a *dataObat, stokObat int) {

}
func cariNamaObat(a dataObat, stokObat int, target string) int {
	// Fungsi untuk mencari nama obat dengan menerima string target lalu mengembalikan indexnya atau -1 jika tidka ditemukan.
	var i int
	for i = 0; i < stokObat; i++ {
		if a[i].nama == target {
			return i
		}
	}
	return -1
}

func sortKadalauarsaAsc(a *dataObat, stokObat int) {
	// Mengurutkan data dengan metode Insertion Sort
	var i, pass int
	var temp Obat
	pass = 1
	for pass < stokObat {
		i = pass
		temp = a[i]
		for i > 0 && temp.Kadaluarsa < a[i-1].Kadaluarsa {
			a[i] = a[i-1]
			i--
		}
		a[i] = temp
		pass++
	}
}

func sortKadalauarsaDesc(a *dataObat, stokObat int) {
	// Mengurutkan data dengan metode Selection Sort
	var i, idx, pass int
	var temp Obat
	pass = 1
	for pass < stokObat {
		idx = pass - 1
		i = pass + 1
		for i < stokObat {
			if a[i].Kadaluarsa > a[idx].Kadaluarsa {
				idx = i
			}
			i++
		}
		temp = a[idx]
		a[idx] = a[pass-1]
		a[pass-1] = temp
		pass++
	}
}

func cetakData(a dataObat, stokObat int) {
	var i int
	for i = 0; i < stokObat; i++ {
		fmt.Printf("ID: %s, Nama: %s, Gejala: %s, Stok: %d, Harga: %.2f, | Kadaluarsa: %s\n", a[i].ID, a[i].nama, a[i].gejala, a[i].Stok, a[i].Harga, a[i].Kadaluarsa)
	}
}
