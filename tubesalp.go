package main

import "fmt"

// Struct Obat menyimpan informasi lengkap satu jenis obat
type Obat struct {
	ID, nama, gejala string
	Stok             int
	Harga            float64
	Kadaluarsa       string
}

// Kapasitas maksimum penyimpanan data obat
const nmax int = 1000

// dataObat adalah tipe array statis berkapasitas nmax
type dataObat [nmax]Obat

// seedData mengisi array dengan data awal (5 obat contoh)
// Parameter:
//   - a        : pointer ke array dataObat yang akan diisi
//   - stokObat : pointer ke jumlah obat yang tersimpan, di-set ke 5
func seedData(a *dataObat, stokObat *int) {
	a[0] = Obat{"OB001", "Paracetamol", "Demam", 100, 5000.0, "2025-12-31"}
	a[1] = Obat{"OB002", "Amoxicillin", "Infeksi", 50, 15000.0, "2024-06-30"}
	a[2] = Obat{"OB003", "Loperamide", "Diare", 200, 8000.0, "2026-03-31"}
	a[3] = Obat{"OB004", "Cetirizine", "Alergi", 75, 12000.0, "2025-09-30"}
	a[4] = Obat{"OB005", "Ibuprofen", "Nyeri", 150, 10000.0, "2024-12-31"}
	*stokObat = 5
}

// main adalah titik masuk program.
// Alur kerja:
//  1. Inisialisasi array dataObat dan isi dengan seedData
//  2. Tampilkan menu secara berulang (loop)
//  3. Baca pilihan pengguna, lalu panggil fungsi yang sesuai
//  4. Loop berhenti jika pengguna memilih 0 (keluar)
func main() {
	var a dataObat
	var stokObat int
	var pilihan int
	var done bool

	// Isi data awal
	seedData(&a, &stokObat)

	// Loop utama: tampilkan menu → baca pilihan → jalankan aksi
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

// cetakMenu menampilkan daftar pilihan menu ke layar
func cetakMenu() {
	fmt.Println("\n=== Aplikasi Apotek ===")
	fmt.Println("1. Tambah Obat")
	fmt.Println("2. Edit Data Obat")
	fmt.Println("3. Hapus Data Obat")
	fmt.Println("4. Sorting Data Obat")
	fmt.Println("5. Cari Data Obat")
	fmt.Println("6. Cetak Data Obat")
	fmt.Println("7. Statistik Obat")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih menu: ")
}

// tambahObat menambahkan satu atau lebih obat baru ke dalam array.
// Alur kerja:
//  1. Minta jumlah obat yang ingin ditambahkan
//  2. Cek apakah kapasitas array masih cukup
//  3. Untuk setiap obat baru:
//     a. Minta nama obat → validasi tidak duplikat via cekNamaObat
//     b. Jika nama sudah ada, minta ulang nama
//     c. Jika nama unik, isi data lengkap (gejala, stok, harga, kadaluarsa)
//     d. Generate ID otomatis menggunakan idObat
//     e. Increment stokObat
func tambahObat(a *dataObat, stokObat *int) {
	var i, newObat int
	var namaObat string

	fmt.Print("Masukkan jumlah obat yang ingin ditambahkan: ")
	fmt.Scan(&newObat)

	// Cek kapasitas: total setelah tambah tidak boleh melebihi nmax
	if newObat+*stokObat > nmax {
		fmt.Println("Jumlah obat melebihi kapasitas maksimum.")
		return
	}

	var tempIndex int
	for tempIndex < newObat {
		fmt.Print("Masukkan nama obat: ")
		fmt.Scan(&namaObat)

		// Validasi: nama obat tidak boleh duplikat
		if !cekNamaObat(*a, namaObat, *stokObat) {
			fmt.Printf("Nama obat '%s' sudah ada. Silakan masukkan nama lain.\n", namaObat)
		} else {
			// Nama unik → isi data di indeks berikutnya
			i = *stokObat
			a[i].nama = namaObat
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

// cekNamaObat memeriksa apakah nama obat belum ada dalam array (sequential search).
// Mengembalikan true  → nama BELUM ada (boleh dipakai)
// Mengembalikan false → nama SUDAH ada (duplikat)
func cekNamaObat(a dataObat, namaObat string, stokObat int) bool {
	var i int
	for i = 0; i < stokObat; i++ {
		if a[i].nama == namaObat {
			return false // ditemukan duplikat
		}
	}
	return true // tidak ada duplikat
}

// idObat menghasilkan string ID berformat "OBxxx" dari angka indeks.
// Contoh: indeks 1 → "OB001", indeks 12 → "OB012", indeks 100 → "OB100"
// Alur kerja:
//  1. Mulai dengan prefix "OB"
//  2. Tambahkan nol di depan sesuai jumlah digit angka
//  3. Konversi angka digit per digit menjadi string, lalu gabungkan
func idObat(index int) string {
	var id string
	id = "OB"

	// Tentukan banyaknya nol padding berdasarkan besar angka
	if index < 10 {
		id += "00"
	} else if index < 100 {
		id += "0"
	}

	// Konversi angka ke string digit per digit (tanpa fungsi bawaan)
	var tempNum, digit int
	var tempString string
	tempNum = index
	for tempNum > 0 {
		digit = tempNum % 10
		tempString = string('0'+rune(digit)) + tempString // susun dari belakang ke depan
		tempNum /= 10
	}
	return id + tempString
}

// editData mengubah salah satu field dari obat yang dipilih.
// Alur kerja:
//  1. Tampilkan pilihan field yang bisa diedit (nama/gejala/stok/harga/kadaluarsa)
//  2. Minta nama obat yang ingin diedit
//  3. Cari obat dengan cariNamaObat → jika tidak ditemukan, keluar
//  4. Sesuai pilihan field, minta nilai baru dan simpan ke array
//  5. Khusus edit nama: validasi tidak duplikat sebelum disimpan
func editData(a *dataObat, stokObat int) {
	var piilihan, idx int
	var namaBaru, cariNama string
	var done bool

	fmt.Println("=== Edit Data Obat ===")
	fmt.Println("1. Nama")
	fmt.Println("2. Gejala")
	fmt.Println("3. Stok")
	fmt.Println("4. Harga")
	fmt.Println("5. Kadaluarsa")
	fmt.Print("Pilih kategori yang ingin diubah: ")
	fmt.Scan(&piilihan)

	fmt.Print("Cari nama obat yang ingin diedit datanya: ")
	fmt.Scan(&cariNama)

	// Cari indeks obat secara sequential
	idx = cariNamaObat(*a, stokObat, cariNama)
	if idx < 0 {
		fmt.Println("Obat tidak ditemukan.")
		return
	}

	switch piilihan {
	case 1:
		// Edit nama: loop sampai nama baru unik
		fmt.Print("Masukkan nama baru: ")
		for done = false; !done; {
			fmt.Scan(&namaBaru)
			if !cekNamaObat(*a, namaBaru, stokObat) {
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

// hapusData menghapus satu obat dari array berdasarkan nama (sequential search).
// Alur kerja:
//  1. Minta nama obat yang ingin dihapus
//  2. Cari indeks obat dengan cariNamaObat → jika tidak ada, keluar
//  3. Geser semua elemen setelah indeks satu posisi ke kiri (overwrite elemen target)
//  4. Kurangi stokObat sebesar 1
func hapusData(a *dataObat, stokObat *int) {
	var i, idx int
	var target string

	fmt.Print("Masukkan nama obat yang ingin dihapus: ")
	fmt.Scan(&target)

	// Cari posisi obat yang akan dihapus
	idx = cariNamaObat(*a, *stokObat, target)
	if idx < 0 {
		fmt.Println("Nama obat yang dimasukkan tidak ada dalam data.")
		return
	}

	// Geser elemen dari idx+1 s.d. akhir ke kiri satu posisi
	for i = idx; i < *stokObat-1; i++ {
		a[i] = a[i+1]
	}
	*stokObat -= 1
	fmt.Printf("Obat '%s' berhasil dihapus.\n", target)
}

// sortingData mengurutkan data obat berdasarkan tanggal kadaluarsa.
// Alur kerja:
//  1. Tampilkan pilihan urutan: ascending (Insertion Sort) atau descending (Selection Sort)
//  2. Baca pilihan pengguna
//  3. Panggil fungsi sort yang sesuai dan tampilkan konfirmasi
func sortingData(a *dataObat, stokObat int) {
	var pilihan int
	fmt.Println("Urutkan berdasarkan tanggal kadaluarsa:")
	fmt.Println("1. Terdekat (Ascending) - Insertion Sort")
	fmt.Println("2. Terjauh (Descending) - Selection Sort")
	fmt.Print("Pilih (1/2): ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		sortKadalauarsaAsc(a, stokObat)
		fmt.Println("Data berhasil diurutkan dari kadaluarsa terdekat.")
	} else if pilihan == 2 {
		sortKadalauarsaDesc(a, stokObat)
		fmt.Println("Data berhasil diurutkan dari kadaluarsa terjauh.")
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

// cariData menjadi dispatcher untuk tiga metode pencarian obat.
// Alur kerja:
//  1. Tampilkan tiga opsi pencarian
//  2. Baca pilihan pengguna
//  3. Panggil fungsi pencarian yang sesuai
func cariData(a *dataObat, stokObat int) {
	var pilihan int
	fmt.Println("=== Cari Data Obat ===")
	fmt.Println("1. Cari berdasarkan Nama (Sequential Search)")
	fmt.Println("2. Cari berdasarkan Gejala (Sequential Search)")
	fmt.Println("3. Cari berdasarkan Nama (Binary Search)")
	fmt.Print("Pilih metode pencarian: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		cariSequentialNama(a, stokObat)
	case 2:
		cariSequentialGejala(a, stokObat)
	case 3:
		cariBinaryNama(a, stokObat)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

// cariSequentialNama mencari obat berdasarkan nama dengan Sequential Search.
// Alur kerja:
//  1. Minta nama target dari pengguna
//  2. Iterasi seluruh array dari indeks 0 s.d. stokObat-1
//  3. Jika nama cocok, cetak detail obat dan tandai ketemu = true
//  4. Jika setelah iterasi ketemu masih false, tampilkan pesan tidak ditemukan
func cariSequentialNama(a *dataObat, stokObat int) {
	var target string
	var i int
	var ketemu bool

	fmt.Print("Masukkan nama obat yang dicari: ")
	fmt.Scan(&target)

	ketemu = false
	for i = 0; i < stokObat; i++ {
		if a[i].nama == target {
			fmt.Println("=== Obat Ditemukan ===")
			fmt.Printf("ID        : %s\n", a[i].ID)
			fmt.Printf("Nama      : %s\n", a[i].nama)
			fmt.Printf("Gejala    : %s\n", a[i].gejala)
			fmt.Printf("Stok      : %d\n", a[i].Stok)
			fmt.Printf("Harga     : %.2f\n", a[i].Harga)
			fmt.Printf("Kadaluarsa: %s\n", a[i].Kadaluarsa)
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Printf("Obat dengan nama '%s' tidak ditemukan.\n", target)
	}
}

// cariSequentialGejala mencari semua obat yang cocok dengan gejala tertentu.
// Alur kerja:
//  1. Minta gejala target dari pengguna
//  2. Iterasi seluruh array; setiap obat yang gejalanya cocok langsung dicetak
//  3. Fungsi ini bisa menampilkan lebih dari satu obat (tidak berhenti saat pertama ditemukan)
//  4. Jika tidak ada yang cocok, tampilkan pesan tidak ditemukan
func cariSequentialGejala(a *dataObat, stokObat int) {
	var target string
	var i int
	var ketemu bool

	fmt.Print("Masukkan gejala yang dicari: ")
	fmt.Scan(&target)

	ketemu = false
	fmt.Println("=== Hasil Pencarian ===")
	for i = 0; i < stokObat; i++ {
		if a[i].gejala == target {
			fmt.Printf("ID: %s | Nama: %s | Stok: %d | Harga: %.2f | Kadaluarsa: %s\n",
				a[i].ID, a[i].nama, a[i].Stok, a[i].Harga, a[i].Kadaluarsa)
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Printf("Tidak ada obat untuk gejala '%s'.\n", target)
	}
}

// cariBinaryNama mencari obat berdasarkan nama menggunakan Binary Search.
// Karena Binary Search mengharuskan data terurut, fungsi ini:
// Alur kerja:
//  1. Salin data ke array sementara (temp) agar array asli tidak berubah
//  2. Urutkan temp berdasarkan nama secara ascending menggunakan Insertion Sort
//  3. Minta nama target dari pengguna
//  4. Lakukan Binary Search pada temp:
//     - Hitung mid = (low + high) / 2
//     - Jika temp[mid].nama == target → cetak detail obat, set ketemu = true, stop
//     - Jika target > temp[mid].nama  → geser low ke mid+1 (cari ke kanan)
//     - Jika target < temp[mid].nama  → geser high ke mid-1 (cari ke kiri)
//  5. Jika ketemu masih false setelah loop, tampilkan pesan tidak ditemukan
func cariBinaryNama(a *dataObat, stokObat int) {
	var target string
	var low, high, mid int
	var ketemu bool
	var temp dataObat
	var i, j int
	var tempObat Obat

	// Salin data ke array sementara (array asli tidak diubah)
	for i = 0; i < stokObat; i++ {
		temp[i] = a[i]
	}

	// Insertion Sort berdasarkan nama (ascending) untuk kebutuhan Binary Search
	for j = 1; j < stokObat; j++ {
		tempObat = temp[j]
		i = j
		for i > 0 && tempObat.nama < temp[i-1].nama {
			temp[i] = temp[i-1]
			i--
		}
		temp[i] = tempObat
	}

	fmt.Print("Masukkan nama obat yang dicari: ")
	fmt.Scan(&target)

	// Binary Search pada array yang sudah terurut
	low = 0
	high = stokObat - 1
	ketemu = false
	for low <= high {
		mid = (low + high) / 2
		if temp[mid].nama == target {
			fmt.Println("=== Obat Ditemukan (Binary Search) ===")
			fmt.Printf("ID        : %s\n", temp[mid].ID)
			fmt.Printf("Nama      : %s\n", temp[mid].nama)
			fmt.Printf("Gejala    : %s\n", temp[mid].gejala)
			fmt.Printf("Stok      : %d\n", temp[mid].Stok)
			fmt.Printf("Harga     : %.2f\n", temp[mid].Harga)
			fmt.Printf("Kadaluarsa: %s\n", temp[mid].Kadaluarsa)
			ketemu = true
			low = high + 1 // paksa loop berhenti
		} else if temp[mid].nama < target {
			low = mid + 1 // target ada di separuh kanan
		} else {
			high = mid - 1 // target ada di separuh kiri
		}
	}
	if !ketemu {
		fmt.Printf("Obat dengan nama '%s' tidak ditemukan.\n", target)
	}
}

// cariNamaObat adalah fungsi pembantu Sequential Search untuk mendapatkan indeks obat.
// Mengembalikan indeks pertama yang cocok, atau -1 jika tidak ditemukan.
// Digunakan oleh editData dan hapusData.
func cariNamaObat(a dataObat, stokObat int, target string) int {
	var i int
	for i = 0; i < stokObat; i++ {
		if a[i].nama == target {
			return i // kembalikan indeks langsung saat ditemukan
		}
	}
	return -1 // tidak ditemukan
}

// cetakStatistik menampilkan laporan ringkas kondisi stok apotek.
// Alur kerja:
//  1. Iterasi array → kumpulkan dan cetak obat dengan stok <= batasStok (hampir habis)
//  2. Iterasi array → kumpulkan dan cetak obat dengan kadaluarsa <= batasKadaluarsa
//     (perbandingan string format YYYY-MM-DD bekerja secara leksikografis/kronologis)
//  3. Tampilkan ringkasan: total jenis obat, jumlah hampir habis, jumlah hampir kadaluarsa
func cetakStatistik(a dataObat, stokObat int) {
	var i int
	var jumlahHampirHabis int
	var batasStok int
	var batasKadaluarsa string

	// Ambang batas: stok <= 20 → hampir habis
	batasStok = 20
	// Ambang batas: kadaluarsa di tahun ini atau sebelumnya → segera kadaluarsa
	batasKadaluarsa = "2025-12-31"

	fmt.Println("\n=== Statistik Obat ===")

	// --- Bagian 1: Obat hampir habis ---
	jumlahHampirHabis = 0
	fmt.Printf("--- Obat Hampir Habis (Stok <= %d) ---\n", batasStok)
	for i = 0; i < stokObat; i++ {
		if a[i].Stok <= batasStok {
			fmt.Printf("Nama: %-15s | Stok: %d\n", a[i].nama, a[i].Stok)
			jumlahHampirHabis++
		}
	}
	if jumlahHampirHabis == 0 {
		fmt.Println("Tidak ada obat yang hampir habis.")
	} else {
		fmt.Printf("Total obat hampir habis: %d\n", jumlahHampirHabis)
	}

	// --- Bagian 2: Obat segera kadaluarsa ---
	fmt.Printf("\n--- Obat Segera Kadaluarsa (Sebelum atau pada %s) ---\n", batasKadaluarsa)
	var jumlahKadaluarsa int
	jumlahKadaluarsa = 0
	for i = 0; i < stokObat; i++ {
		// Perbandingan string format YYYY-MM-DD berfungsi sebagai perbandingan tanggal
		if a[i].Kadaluarsa <= batasKadaluarsa {
			fmt.Printf("Nama: %-15s | Kadaluarsa: %s\n", a[i].nama, a[i].Kadaluarsa)
			jumlahKadaluarsa++
		}
	}
	if jumlahKadaluarsa == 0 {
		fmt.Println("Tidak ada obat yang akan segera kadaluarsa.")
	} else {
		fmt.Printf("Total obat segera kadaluarsa: %d\n", jumlahKadaluarsa)
	}

	// --- Bagian 3: Ringkasan akhir ---
	fmt.Println("\n--- Ringkasan ---")
	fmt.Printf("Total jenis obat       : %d\n", stokObat)
	fmt.Printf("Obat hampir habis      : %d\n", jumlahHampirHabis)
	fmt.Printf("Obat segera kadaluarsa : %d\n", jumlahKadaluarsa)
}

// sortKadalauarsaAsc mengurutkan data obat berdasarkan kadaluarsa dari terdekat ke terjauh
// menggunakan algoritma Insertion Sort (ascending).
// Alur kerja Insertion Sort:
//  1. Mulai dari elemen indeks 1 (pass = 1)
//  2. Simpan elemen saat ini ke temp
//  3. Geser elemen-elemen sebelumnya yang lebih besar satu posisi ke kanan
//  4. Sisipkan temp di posisi yang tepat
//  5. Ulangi hingga semua elemen terproses
func sortKadalauarsaAsc(a *dataObat, stokObat int) {
	var i, pass int
	var temp Obat

	pass = 1
	for pass < stokObat {
		i = pass
		temp = a[i] // simpan elemen yang akan disisipkan

		// Geser elemen yang lebih besar ke kanan
		for i > 0 && temp.Kadaluarsa < a[i-1].Kadaluarsa {
			a[i] = a[i-1]
			i--
		}
		a[i] = temp // sisipkan di posisi yang tepat
		pass++
	}
}

// sortKadalauarsaDesc mengurutkan data obat berdasarkan kadaluarsa dari terjauh ke terdekat
// menggunakan algoritma Selection Sort (descending).
// Alur kerja Selection Sort:
//  1. Mulai dari posisi 0 (pass = 1, idx = pass-1)
//  2. Cari elemen dengan kadaluarsa terbesar (terjauh) pada sisa array (indeks pass s.d. akhir)
//  3. Tukar elemen terbesar tersebut dengan elemen di posisi pass-1
//  4. Ulangi hingga semua posisi terisi elemen yang benar
func sortKadalauarsaDesc(a *dataObat, stokObat int) {
	var i, idx, pass int
	var temp Obat

	pass = 1
	for pass < stokObat {
		idx = pass - 1 // asumsikan elemen terbesar ada di posisi pass-1

		// Cari elemen dengan kadaluarsa terbesar di sisa array
		i = pass + 1
		for i < stokObat {
			if a[i].Kadaluarsa > a[idx].Kadaluarsa {
				idx = i // perbarui posisi terbesar
			}
			i++
		}

		// Tukar elemen terbesar ke posisi pass-1
		temp = a[idx]
		a[idx] = a[pass-1]
		a[pass-1] = temp
		pass++
	}
}

// cetakData menampilkan seluruh data obat dalam format tabel satu baris per obat.
// Alur kerja: iterasi array dari indeks 0 s.d. stokObat-1, cetak setiap elemen.
func cetakData(a dataObat, stokObat int) {
	var i int
	fmt.Println("=== Daftar Obat ===")
	for i = 0; i < stokObat; i++ {
		fmt.Printf("ID: %s | Nama: %-15s | Gejala: %-10s | Stok: %3d | Harga: %8.2f | Kadaluarsa: %s\n",
			a[i].ID, a[i].nama, a[i].gejala, a[i].Stok, a[i].Harga, a[i].Kadaluarsa)
	}
}