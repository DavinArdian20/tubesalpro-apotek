package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// ============================================================
// STRUKTUR DATA
// ============================================================

// Obat merepresentasikan satu data obat di apotek.
type Obat struct {
	ID         string
	Nama       string
	KategoriID string 
	Gejala     string 
	Stok       int
	Harga      float64
	Kadaluarsa time.Time
}

type KategoriGejala struct {
	ID   string
	Nama string
}
type TransaksiStok struct {
	ID        string
	ObatID    string
	Jumlah    int
	Tanggal   time.Time
	Keterangan string
}

// ============================================================
// PENYIMPANAN DATA (Slice)
// ============================================================

var daftarObat      []Obat
var daftarKategori  []KategoriGejala
var daftarTransaksi []TransaksiStok

// reader untuk membaca input terminal
var reader = bufio.NewReader(os.Stdin)

// Format tanggal standar
const FORMAT_TANGGAL = "02-01-2006"

// ============================================================
// DATA DUMMY
// ============================================================

func inisialisasiData() {
	// --- Kategori Gejala ---
	daftarKategori = []KategoriGejala{
		{"KAT01", "Demam & Flu"},
		{"KAT02", "Batuk & Pilek"},
		{"KAT03", "Sakit Kepala & Nyeri"},
		{"KAT04", "Gangguan Pencernaan"},
		{"KAT05", "Alergi & Gatal"},
		{"KAT06", "Vitamin & Suplemen"},
		{"KAT07", "Infeksi & Antibiotik"},
	}

	// --- Data Obat ---
	obatDummy := []struct {
		id, nama, katID, gejala string
		stok                    int
		harga                   float64
		tgl                     string
	}{
		{"OBT001", "Paracetamol", "KAT01", "Demam, Sakit Kepala, Nyeri Ringan", 50, 5000, "15-12-2025"},
		{"OBT002", "Bodrex", "KAT01", "Demam, Flu, Sakit Kepala", 10, 8500, "01-03-2027"},
		{"OBT003", "OBH Combi", "KAT02", "Batuk Berdahak, Flu, Hidung Tersumbat", 30, 15000, "20-06-2026"},
		{"OBT004", "Woods", "KAT02", "Batuk Kering, Batuk Berdahak", 25, 18000, "10-09-2026"},
		{"OBT005", "Promag", "KAT04", "Maag, Mual, Kembung, Nyeri Lambung", 20, 12000, "01-01-2027"},
		{"OBT006", "Antangin", "KAT04", "Masuk Angin, Mual, Perut Kembung", 40, 7000, "30-11-2025"},
		{"OBT007", "Tolak Angin", "KAT04", "Masuk Angin, Mual, Badan Pegal", 35, 9500, "05-07-2026"},
		{"OBT008", "Panadol", "KAT01", "Demam, Sakit Kepala, Nyeri Otot", 15, 11000, "01-05-2027"},
		{"OBT009", "Decolgen", "KAT01", "Flu, Hidung Tersumbat, Demam", 10, 13500, "25-08-2026"},
		{"OBT010", "Mixagrip", "KAT01", "Flu, Sakit Kepala, Demam, Bersin", 10, 14000, "01-04-2027"},
		{"OBT011", "Ibuprofen", "KAT03", "Sakit Kepala, Nyeri Sendi, Demam", 20, 6500, "12-02-2026"},
		{"OBT012", "Antasida", "KAT04", "Maag, Sakit Perut, Mulas", 18, 8000, "18-07-2026"},
		{"OBT013", "Cetirizine", "KAT05", "Alergi, Gatal-gatal, Bersin, Flu Alergi", 12, 9000, "22-10-2026"},
		{"OBT014", "Amoxicillin", "KAT07", "Infeksi Bakteri, Batuk Berdahak", 8, 25000, "15-01-2026"},
		{"OBT015", "Vitamin C", "KAT06", "Daya Tahan Tubuh, Kelelahan, Flu Ringan", 60, 4500, "31-12-2027"},
	}

	for _, d := range obatDummy {
		tgl, _ := time.Parse(FORMAT_TANGGAL, d.tgl)
		daftarObat = append(daftarObat, Obat{
			ID: d.id, Nama: d.nama, KategoriID: d.katID,
			Gejala: d.gejala, Stok: d.stok, Harga: d.harga, Kadaluarsa: tgl,
		})
	}

	// --- Transaksi Stok Masuk ---
	transaksiDummy := []struct {
		id, obatID, ket string
		jumlah          int
		tgl             string
	}{
		{"TRX001", "OBT001", "Pembelian rutin dari distributor", 50, "01-11-2025"},
		{"TRX002", "OBT015", "Stok awal apotek", 60, "01-11-2025"},
		{"TRX003", "OBT003", "Restock OBH Combi", 30, "05-11-2025"},
		{"TRX004", "OBT006", "Pembelian grosir Antangin", 40, "10-11-2025"},
		{"TRX005", "OBT014", "Restock Amoxicillin resep dokter", 8, "15-11-2025"},
	}
	for _, t := range transaksiDummy {
		tgl, _ := time.Parse(FORMAT_TANGGAL, t.tgl)
		daftarTransaksi = append(daftarTransaksi, TransaksiStok{
			ID: t.id, ObatID: t.obatID, Jumlah: t.jumlah,
			Tanggal: tgl, Keterangan: t.ket,
		})
	}
}

// ============================================================
// UTILITAS INPUT
// ============================================================

func bacaInput(prompt string) string {
	fmt.Print(prompt)
	teks, _ := reader.ReadString('\n')
	return strings.TrimSpace(teks)
}

func bacaInt(prompt string) int {
	for {
		teks := bacaInput(prompt)
		angka, err := strconv.Atoi(teks)
		if err == nil {
			return angka
		}
		fmt.Println("  [!] Input harus berupa angka bulat. Coba lagi.")
	}
}

func bacaFloat(prompt string) float64 {
	for {
		teks := bacaInput(prompt)
		angka, err := strconv.ParseFloat(teks, 64)
		if err == nil {
			return angka
		}
		fmt.Println("  [!] Input harus berupa angka. Coba lagi.")
	}
}

func bacaTanggal(prompt string) time.Time {
	for {
		teks := bacaInput(prompt + " (dd-mm-yyyy): ")
		tgl, err := time.Parse(FORMAT_TANGGAL, teks)
		if err == nil {
			return tgl
		}
		fmt.Println("  [!] Format tidak valid. Gunakan dd-mm-yyyy.")
	}
}

// ============================================================
// UTILITAS TAMPILAN
// ============================================================

func cetakGaris() {
	fmt.Println(strings.Repeat("=", 105))
}

func cetakGarisPendek() {
	fmt.Println(strings.Repeat("-", 105))
}

// cetakHeaderObat mencetak header tabel obat.
func cetakHeaderObat() {
	cetakGaris()
	fmt.Printf("  %-8s %-18s %-8s %-30s %-5s %-11s %-12s\n",
		"ID", "Nama Obat", "KatID", "Gejala", "Stok", "Harga(Rp)", "Kadaluarsa")
	cetakGaris()
}

// cetakBarisObat mencetak satu baris data obat.
func cetakBarisObat(o Obat) {
	gejala := o.Gejala
	if len(gejala) > 28 {
		gejala = gejala[:25] + "..."
	}
	fmt.Printf("  %-8s %-18s %-8s %-30s %-5d %-11.0f %-12s\n",
		o.ID, o.Nama, o.KategoriID, gejala, o.Stok, o.Harga,
		o.Kadaluarsa.Format(FORMAT_TANGGAL))
}

// cariIndexObatByID mencari indeks obat berdasarkan ID (Linear Search).
func cariIndexObatByID(id string) int {
	for i, o := range daftarObat {
		if strings.EqualFold(o.ID, id) {
			return i
		}
	}
	return -1
}

// cariIndexKategoriByID mencari indeks kategori berdasarkan ID.
func cariIndexKategoriByID(id string) int {
	for i, k := range daftarKategori {
		if strings.EqualFold(k.ID, id) {
			return i
		}
	}
	return -1
}

// namaKategori mengembalikan nama kategori berdasarkan ID.
func namaKategori(id string) string {
	idx := cariIndexKategoriByID(id)
	if idx == -1 {
		return "-"
	}
	return daftarKategori[idx].Nama
}

// namaObat mengembalikan nama obat berdasarkan ID.
func namaObat(id string) string {
	idx := cariIndexObatByID(id)
	if idx == -1 {
		return "-"
	}
	return daftarObat[idx].Nama
}

// ============================================================
// ============================================================
//                   MENU OBAT
// ============================================================
// ============================================================
func tambahObat() {
	fmt.Println("\n  [ TAMBAH DATA OBAT ]")
	cetakGaris()

	// Validasi ID
	var id string
	for {
		id = bacaInput("  ID Obat          : ")
		if id == "" {
			fmt.Println("  [!] ID tidak boleh kosong.")
			continue
		}
		if cariIndexObatByID(id) != -1 {
			fmt.Println("  [!] ID sudah digunakan.")
			continue
		}
		break
	}

	// Validasi Nama
	var nama string
	for {
		nama = bacaInput("  Nama Obat        : ")
		if nama != "" {
			break
		}
		fmt.Println("  [!] Nama tidak boleh kosong.")
	}

	// Pilih Kategori
	tampilkanKategori()
	var katID string
	for {
		katID = bacaInput("  ID Kategori       : ")
		if cariIndexKategoriByID(katID) != -1 {
			break
		}
		fmt.Println("  [!] ID Kategori tidak ditemukan. Coba lagi.")
	}

	gejala := bacaInput("  Gejala Detail     : ")

	// Validasi Stok
	var stok int
	for {
		stok = bacaInt("  Stok              : ")
		if stok >= 0 {
			break
		}
		fmt.Println("  [!] Stok tidak boleh negatif.")
	}

	// Validasi Harga
	var harga float64
	for {
		harga = bacaFloat("  Harga (Rp)        : ")
		if harga >= 0 {
			break
		}
		fmt.Println("  [!] Harga tidak boleh negatif.")
	}

	kadaluarsa := bacaTanggal("  Tanggal Kadaluarsa")

	daftarObat = append(daftarObat, Obat{
		ID: id, Nama: nama, KategoriID: katID,
		Gejala: gejala, Stok: stok, Harga: harga, Kadaluarsa: kadaluarsa,
	})
	fmt.Println("  [✓] Data obat berhasil ditambahkan!")
}

func ubahObat() {
	fmt.Println("\n  [ UBAH DATA OBAT ]")
	cetakGaris()

	id := bacaInput("  ID Obat yang diubah: ")
	idx := cariIndexObatByID(id)
	if idx == -1 {
		fmt.Println("  [!] Obat tidak ditemukan.")
		return
	}
	o := daftarObat[idx]
	fmt.Println("\n  Data saat ini:")
	cetakHeaderObat()
	cetakBarisObat(o)
	fmt.Println("\n  Tekan Enter untuk melewati field (tidak mengubah):")

	namaBaru := bacaInput(fmt.Sprintf("  Nama [%s]: ", o.Nama))
	if namaBaru != "" {
		daftarObat[idx].Nama = namaBaru
	}

	tampilkanKategori()
	katBaru := bacaInput(fmt.Sprintf("  ID Kategori [%s]: ", o.KategoriID))
	if katBaru != "" && cariIndexKategoriByID(katBaru) != -1 {
		daftarObat[idx].KategoriID = katBaru
	} else if katBaru != "" {
		fmt.Println("  [!] ID Kategori tidak valid, tidak diubah.")
	}

	gejalaBaru := bacaInput(fmt.Sprintf("  Gejala [%s]: ", o.Gejala))
	if gejalaBaru != "" {
		daftarObat[idx].Gejala = gejalaBaru
	}

	stokStr := bacaInput(fmt.Sprintf("  Stok [%d]: ", o.Stok))
	if stokStr != "" {
		if v, err := strconv.Atoi(stokStr); err == nil && v >= 0 {
			daftarObat[idx].Stok = v
		} else {
			fmt.Println("  [!] Stok tidak valid, tidak diubah.")
		}
	}

	hargaStr := bacaInput(fmt.Sprintf("  Harga [%.0f]: ", o.Harga))
	if hargaStr != "" {
		if v, err := strconv.ParseFloat(hargaStr, 64); err == nil && v >= 0 {
			daftarObat[idx].Harga = v
		} else {
			fmt.Println("  [!] Harga tidak valid, tidak diubah.")
		}
	}

	tglStr := bacaInput(fmt.Sprintf("  Kadaluarsa [%s] (dd-mm-yyyy): ", o.Kadaluarsa.Format(FORMAT_TANGGAL)))
	if tglStr != "" {
		if tgl, err := time.Parse(FORMAT_TANGGAL, tglStr); err == nil {
			daftarObat[idx].Kadaluarsa = tgl
		} else {
			fmt.Println("  [!] Format tanggal tidak valid, tidak diubah.")
		}
	}

	fmt.Println("  [✓] Data obat berhasil diperbarui!")
}

func hapusObat() {
	fmt.Println("\n  [ HAPUS DATA OBAT ]")
	cetakGaris()

	id := bacaInput("  ID Obat yang dihapus: ")
	idx := cariIndexObatByID(id)
	if idx == -1 {
		fmt.Println("  [!] Obat tidak ditemukan.")
		return
	}

	fmt.Println("\n  Data yang akan dihapus:")
	cetakHeaderObat()
	cetakBarisObat(daftarObat[idx])
	fmt.Println()

	konfirmasi := bacaInput("  Yakin hapus? (y/n): ")
	if strings.ToLower(konfirmasi) != "y" {
		fmt.Println("  [i] Penghapusan dibatalkan.")
		return
	}
	daftarObat = append(daftarObat[:idx], daftarObat[idx+1:]...)
	fmt.Println("  [✓] Data obat berhasil dihapus!")
}

func tampilkanSemua() {
	fmt.Println("\n  [ DAFTAR SEMUA OBAT ]")
	if len(daftarObat) == 0 {
		fmt.Println("  [i] Belum ada data obat.")
		return
	}
	cetakHeaderObat()
	for _, o := range daftarObat {
		cetakBarisObat(o)
	}
	cetakGaris()
	fmt.Printf("  Total: %d jenis obat\n", len(daftarObat))
}

// ============================================================
//                   MENU KATEGORI GEJALA
// ============================================================

func tampilkanKategori() {
	fmt.Println("\n  [ DAFTAR KATEGORI GEJALA ]")
	cetakGarisPendek()
	fmt.Printf("  %-8s %-30s\n", "ID", "Nama Kategori")
	cetakGarisPendek()
	for _, k := range daftarKategori {
		fmt.Printf("  %-8s %-30s\n", k.ID, k.Nama)
	}
	cetakGarisPendek()
}

func tambahKategori() {
	fmt.Println("\n  [ TAMBAH KATEGORI GEJALA ]")
	cetakGaris()

	var id string
	for {
		id = bacaInput("  ID Kategori  : ")
		if id == "" {
			fmt.Println("  [!] ID tidak boleh kosong.")
			continue
		}
		if cariIndexKategoriByID(id) != -1 {
			fmt.Println("  [!] ID sudah digunakan.")
			continue
		}
		break
	}

	var nama string
	for {
		nama = bacaInput("  Nama Kategori: ")
		if nama != "" {
			break
		}
		fmt.Println("  [!] Nama tidak boleh kosong.")
	}

	daftarKategori = append(daftarKategori, KategoriGejala{ID: id, Nama: nama})
	fmt.Println("  [✓] Kategori berhasil ditambahkan!")
}

// ubahKategori mengubah data kategori berdasarkan ID.
func ubahKategori() {
	fmt.Println("\n  [ UBAH KATEGORI GEJALA ]")
	tampilkanKategori()

	id := bacaInput("  ID Kategori yang diubah: ")
	idx := cariIndexKategoriByID(id)
	if idx == -1 {
		fmt.Println("  [!] Kategori tidak ditemukan.")
		return
	}

	namaBaru := bacaInput(fmt.Sprintf("  Nama Baru [%s]: ", daftarKategori[idx].Nama))
	if namaBaru != "" {
		daftarKategori[idx].Nama = namaBaru
		fmt.Println("  [✓] Kategori berhasil diperbarui!")
	} else {
		fmt.Println("  [i] Tidak ada perubahan.")
	}
}
func hapusKategori() {
	fmt.Println("\n  [ HAPUS KATEGORI GEJALA ]")
	tampilkanKategori()

	id := bacaInput("  ID Kategori yang dihapus: ")
	idx := cariIndexKategoriByID(id)
	if idx == -1 {
		fmt.Println("  [!] Kategori tidak ditemukan.")
		return
	}

	konfirmasi := bacaInput(fmt.Sprintf("  Hapus kategori '%s'? (y/n): ", daftarKategori[idx].Nama))
	if strings.ToLower(konfirmasi) != "y" {
		fmt.Println("  [i] Penghapusan dibatalkan.")
		return
	}
	daftarKategori = append(daftarKategori[:idx], daftarKategori[idx+1:]...)
	fmt.Println("  [✓] Kategori berhasil dihapus!")
}

// ============================================================
//                   MENU TRANSAKSI STOK MASUK
// ============================================================

func tambahTransaksi() {
	fmt.Println("\n  [ CATAT STOK MASUK ]")
	cetakGaris()

	// Generate ID otomatis
	idBaru := fmt.Sprintf("TRX%03d", len(daftarTransaksi)+1)
	fmt.Printf("  ID Transaksi  : %s (otomatis)\n", idBaru)

	tampilkanSemua()
	obatID := bacaInput("  ID Obat       : ")
	idxObat := cariIndexObatByID(obatID)
	if idxObat == -1 {
		fmt.Println("  [!] Obat tidak ditemukan.")
		return
	}
	fmt.Printf("  Obat dipilih  : %s\n", daftarObat[idxObat].Nama)
	fmt.Printf("  Stok saat ini : %d unit\n", daftarObat[idxObat].Stok)

	var jumlah int
	for {
		jumlah = bacaInt("  Jumlah Masuk  : ")
		if jumlah > 0 {
			break
		}
		fmt.Println("  [!] Jumlah harus lebih dari 0.")
	}

	keterangan := bacaInput("  Keterangan    : ")
	tanggal := time.Now()

	// Simpan transaksi
	daftarTransaksi = append(daftarTransaksi, TransaksiStok{
		ID:         idBaru,
		ObatID:     obatID,
		Jumlah:     jumlah,
		Tanggal:    tanggal,
		Keterangan: keterangan,
	})
	daftarObat[idxObat].Stok += jumlah

	fmt.Printf("\n  [✓] Transaksi berhasil! Stok %s: %d → %d unit\n",
		daftarObat[idxObat].Nama,
		daftarObat[idxObat].Stok-jumlah,
		daftarObat[idxObat].Stok)
}
func tampilkanTransaksi() {
	fmt.Println("\n  [ RIWAYAT TRANSAKSI STOK MASUK ]")
	if len(daftarTransaksi) == 0 {
		fmt.Println("  [i] Belum ada transaksi.")
		return
	}

	cetakGarisPendek()
	fmt.Printf("  %-8s %-8s %-20s %-6s %-12s %-25s\n",
		"ID Trx", "ID Obat", "Nama Obat", "Jumlah", "Tanggal", "Keterangan")
	cetakGarisPendek()
	for _, t := range daftarTransaksi {
		ket := t.Keterangan
		if len(ket) > 23 {
			ket = ket[:20] + "..."
		}
		fmt.Printf("  %-8s %-8s %-20s %-6d %-12s %-25s\n",
			t.ID, t.ObatID, namaObat(t.ObatID), t.Jumlah,
			t.Tanggal.Format(FORMAT_TANGGAL), ket)
	}
	cetakGarisPendek()
	fmt.Printf("  Total transaksi: %d\n", len(daftarTransaksi))
}
// ============================================================
//          SEARCHING: SEQUENTIAL & BINARY SEARCH
// ============================================================
func cariByGejalaSekuensial() {
	fmt.Println("\n  [ CARI OBAT BERDASARKAN GEJALA — Sequential Search ]")
	cetakGaris()

	keyword := bacaInput("  Masukkan gejala (misal: Demam, Batuk, Flu, Maag): ")
	if keyword == "" {
		fmt.Println("  [!] Kata kunci tidak boleh kosong.")
		return
	}
	keywordLower := strings.ToLower(keyword)

	fmt.Printf("\n  Mencari menggunakan Sequential Search...\n")
	langkah := 0
	var hasil []Obat
	for _, o := range daftarObat {
		langkah++
		if strings.Contains(strings.ToLower(o.Gejala), keywordLower) ||
			strings.Contains(strings.ToLower(namaKategori(o.KategoriID)), keywordLower) {
			hasil = append(hasil, o)
		}
	}

	fmt.Printf("  Total langkah pemeriksaan: %d dari %d data\n", langkah, len(daftarObat))
	fmt.Printf("\n  Hasil pencarian gejala \"%s\":\n", keyword)

	if len(hasil) == 0 {
		fmt.Println("  [i] Tidak ada obat yang cocok.")
		return
	}
	cetakHeaderObat()
	for _, o := range hasil {
		cetakBarisObat(o)
	}
	cetakGaris()
	fmt.Printf("  Ditemukan: %d obat\n", len(hasil))
}

func cariByNamaBinarySearch() {
	fmt.Println("\n  [ CARI OBAT BERDASARKAN NAMA — Binary Search ]")
	cetakGaris()

	keyword := bacaInput("  Masukkan nama obat: ")
	if keyword == "" {
		fmt.Println("  [!] Nama obat tidak boleh kosong.")
		return
	}
	keywordLower := strings.ToLower(keyword)

	terurut := make([]Obat, len(daftarObat))
	copy(terurut, daftarObat)
	sort.Slice(terurut, func(i, j int) bool {
		return strings.ToLower(terurut[i].Nama) < strings.ToLower(terurut[j].Nama)
	})

	fmt.Printf("\n  Data diurutkan berdasarkan nama untuk Binary Search...\n")

	low, high := 0, len(terurut)-1
	langkah := 0
	idxDitemukan := -1

	for low <= high {
		langkah++
		mid := (low + high) / 2
		namaMid := strings.ToLower(terurut[mid].Nama)

		fmt.Printf("  Langkah %d: Periksa indeks %d → \"%s\"\n", langkah, mid, terurut[mid].Nama)

		if strings.HasPrefix(namaMid, keywordLower) || namaMid == keywordLower {
			idxDitemukan = mid
			break
		} else if namaMid < keywordLower {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	fmt.Printf("  Total langkah Binary Search: %d (dari %d data)\n", langkah, len(terurut))

	if idxDitemukan == -1 {
		fmt.Println("\n  [i] Obat dengan nama tersebut tidak ditemukan.")
		return
	}
	var hasil []Obat
	for i := idxDitemukan; i >= 0; i-- {
		if strings.Contains(strings.ToLower(terurut[i].Nama), keywordLower) {
			hasil = append([]Obat{terurut[i]}, hasil...)
		} else {
			break
		}
	}
	for i := idxDitemukan + 1; i < len(terurut); i++ {
		if strings.Contains(strings.ToLower(terurut[i].Nama), keywordLower) {
			hasil = append(hasil, terurut[i])
		} else {
			break
		}
	}

	fmt.Printf("\n  Hasil pencarian nama \"%s\":\n", keyword)
	cetakHeaderObat()
	for _, o := range hasil {
		cetakBarisObat(o)
	}
	cetakGaris()
	fmt.Printf("  Ditemukan: %d obat\n", len(hasil))
}
func sortSelectionSort(ascending bool) {
	hasil := make([]Obat, len(daftarObat))
	copy(hasil, daftarObat)
	n := len(hasil)

	for i := 0; i < n-1; i++ {
		idxPilih := i
		for j := i + 1; j < n; j++ {
			if ascending {
				if hasil[j].Kadaluarsa.Before(hasil[idxPilih].Kadaluarsa) {
					idxPilih = j
				}
			} else {
				if hasil[j].Kadaluarsa.After(hasil[idxPilih].Kadaluarsa) {
					idxPilih = j
				}
			}
		}
		hasil[i], hasil[idxPilih] = hasil[idxPilih], hasil[i]
	}

	arah := "ASCENDING (Kadaluarsa Terdekat → Terjauh)"
	if !ascending {
		arah = "DESCENDING (Kadaluarsa Terjauh → Terdekat)"
	}
	fmt.Printf("\n  [ SELECTION SORT — Kadaluarsa %s ]\n", arah)
	cetakHeaderObat()
	for _, o := range hasil {
		cetakBarisObat(o)
	}
	cetakGaris()
}
func sortInsertionSort(ascending bool) {
	hasil := make([]Obat, len(daftarObat))
	copy(hasil, daftarObat)
	n := len(hasil)

	for i := 1; i < n; i++ {
		kunci := hasil[i]
		j := i - 1

		for j >= 0 {
			lebihBesar := false
			if ascending {
				lebihBesar = hasil[j].Kadaluarsa.After(kunci.Kadaluarsa)
			} else {
				lebihBesar = hasil[j].Kadaluarsa.Before(kunci.Kadaluarsa)
			}

			if lebihBesar {
				hasil[j+1] = hasil[j]
				j--
			} else {
				break
			}
		}
		hasil[j+1] = kunci
	}

	arah := "ASCENDING (Kadaluarsa Terdekat → Terjauh)"
	if !ascending {
		arah = "DESCENDING (Kadaluarsa Terjauh → Terdekat)"
	}
	fmt.Printf("\n  [ INSERTION SORT — Kadaluarsa %s ]\n", arah)
	cetakHeaderObat()
	for _, o := range hasil {
		cetakBarisObat(o)
	}
	cetakGaris()
}

func sortMultiKriteria() {
	hasil := make([]Obat, len(daftarObat))
	copy(hasil, daftarObat)
	n := len(hasil)
	for i := 1; i < n; i++ {
		kunci := hasil[i]
		j := i - 1
		for j >= 0 {
			harus_geser := false
			if hasil[j].Stok != kunci.Stok {
				harus_geser = hasil[j].Stok > kunci.Stok
			} else {
				harus_geser = hasil[j].Kadaluarsa.After(kunci.Kadaluarsa)
			}

			if harus_geser {
				hasil[j+1] = hasil[j]
				j--
			} else {
				break
			}
		}
		hasil[j+1] = kunci
	}

	fmt.Println("\n  [ SORT MULTI-KRITERIA: Stok (Kecil→Besar) + Kadaluarsa (Terdekat→Terjauh) ]")
	fmt.Println("  Prioritas 1: Stok terkecil ke terbesar")
	fmt.Println("  Prioritas 2: Jika stok sama → kadaluarsa terdekat dulu")
	cetakHeaderObat()
	for _, o := range hasil {
		cetakBarisObat(o)
	}
	cetakGaris()
}

// ============================================================
// ============================================================
//                   STATISTIK STOK
// ============================================================
// ============================================================

func statistikStok() {
	fmt.Println("\n  [ STATISTIK STOK OBAT ]")
	cetakGaris()

	if len(daftarObat) == 0 {
		fmt.Println("  [i] Belum ada data obat.")
		return
	}

	totalStok := 0
	idxTerbanyak, idxTersedikit := 0, 0
	sekarang := time.Now()
	var obatHampirHabis []Obat    
	var obatKadaluarsa30 []Obat   
	for i, o := range daftarObat {
		totalStok += o.Stok

		if o.Stok > daftarObat[idxTerbanyak].Stok {
			idxTerbanyak = i
		}
		if o.Stok < daftarObat[idxTersedikit].Stok {
			idxTersedikit = i
		}
		if o.Stok < 15 {
			obatHampirHabis = append(obatHampirHabis, o)
		}

		selisih := o.Kadaluarsa.Sub(sekarang)
		if selisih >= 0 && selisih <= 30*24*time.Hour {
			obatKadaluarsa30 = append(obatKadaluarsa30, o)
		}
	}

	fmt.Printf("  %-42s: %d unit\n", "Total seluruh stok", totalStok)
	fmt.Printf("  %-42s: %s (%d unit)\n", "Obat stok terbanyak",
		daftarObat[idxTerbanyak].Nama, daftarObat[idxTerbanyak].Stok)
	fmt.Printf("  %-42s: %s (%d unit)\n", "Obat stok paling sedikit",
		daftarObat[idxTersedikit].Nama, daftarObat[idxTersedikit].Stok)
	fmt.Printf("  %-42s: %d jenis\n", "Jumlah jenis obat", len(daftarObat))
	fmt.Printf("  %-42s: %d jenis\n", "Jumlah kategori gejala", len(daftarKategori))
	fmt.Printf("  %-42s: %d transaksi\n", "Total transaksi stok masuk", len(daftarTransaksi))
	fmt.Printf("  %-42s: %d obat\n", "Obat hampir habis (stok < 15)", len(obatHampirHabis))
	fmt.Printf("  %-42s: %d obat\n", "Obat kadaluarsa dalam ≤ 30 hari", len(obatKadaluarsa30))
	cetakGaris()

	// Daftar obat hampir habis
	if len(obatHampirHabis) > 0 {
		fmt.Println("\n  [!] OBAT YANG STOKNYA HAMPIR HABIS (< 15 unit):")
		cetakGarisPendek()
		for _, o := range obatHampirHabis {
			fmt.Printf("    ⚠  %-20s | Stok: %d unit\n", o.Nama, o.Stok)
		}
	}

	// Daftar obat yang akan segera kadaluarsa
	if len(obatKadaluarsa30) > 0 {
		fmt.Println("\n  [!] OBAT YANG AKAN SEGERA KADALUARSA (≤ 30 hari):")
		cetakGarisPendek()
		for _, o := range obatKadaluarsa30 {
			sisaHari := int(o.Kadaluarsa.Sub(sekarang).Hours() / 24)
			fmt.Printf("    ⚠  %-20s | Kadaluarsa: %s | Sisa: %d hari\n",
				o.Nama, o.Kadaluarsa.Format(FORMAT_TANGGAL), sisaHari)
		}
	} else {
		fmt.Println("\n  [✓] Tidak ada obat yang akan kadaluarsa dalam 30 hari ke depan.")
	}
}

// ============================================================
// MENU UTAMA
// ============================================================

func tampilkanMenu() {
	fmt.Println()
	cetakGaris()
	fmt.Println("              APOTEK-SMART: SISTEM MANAJEMEN STOK & INVENTARIS APOTEK              ")
	cetakGaris()
	fmt.Println("  ── MENU OBAT ──────────────────────────────────────────")
	fmt.Println("   1. Tambah Obat            2. Ubah Obat")
	fmt.Println("   3. Hapus Obat             4. Tampilkan Semua Obat")
	fmt.Println()
	fmt.Println("  ── MENU KATEGORI GEJALA ───────────────────────────────")
	fmt.Println("   5. Tambah Kategori        6. Ubah Kategori")
	fmt.Println("   7. Hapus Kategori         8. Tampilkan Semua Kategori")
	fmt.Println()
	fmt.Println("  ── MENU TRANSAKSI STOK MASUK ──────────────────────────")
	fmt.Println("   9. Catat Stok Masuk      10. Riwayat Transaksi")
	fmt.Println()
	fmt.Println("  ── PENCARIAN ──────────────────────────────────────────")
	fmt.Println("  11. Cari Gejala (Sequential Search)")
	fmt.Println("  12. Cari Nama   (Binary Search)")
	fmt.Println()
	fmt.Println("  ── PENGURUTAN ─────────────────────────────────────────")
	fmt.Println("  13. Selection Sort  — Kadaluarsa Ascending")
	fmt.Println("  14. Selection Sort  — Kadaluarsa Descending")
	fmt.Println("  15. Insertion Sort  — Kadaluarsa Ascending")
	fmt.Println("  16. Insertion Sort  — Kadaluarsa Descending")
	fmt.Println("  17. Sort Multi-Kriteria (Stok + Kadaluarsa)")
	fmt.Println()
	fmt.Println("  ── LAINNYA ────────────────────────────────────────────")
	fmt.Println("  18. Statistik Stok")
	fmt.Println("   0. Keluar")
	cetakGaris()
}

func main() {
	fmt.Println()
	cetakGaris()
	fmt.Println("  Selamat Datang di APOTEK-SMART")
	fmt.Println("  Sistem Manajemen Stok dan Inventaris Apotek")
	fmt.Println("  Menginisialisasi data dummy...")
	cetakGaris()

	inisialisasiData()
	fmt.Printf("  [✓] %d data obat, %d kategori, %d transaksi berhasil dimuat.\n",
		len(daftarObat), len(daftarKategori), len(daftarTransaksi))

	for {
		tampilkanMenu()
		pilihan := bacaInput("  Pilih menu [0-18]: ")

		switch pilihan {
		case "1":
			tambahObat()
		case "2":
			ubahObat()
		case "3":
			hapusObat()
		case "4":
			tampilkanSemua()
		case "5":
			tambahKategori()
		case "6":
			ubahKategori()
		case "7":
			hapusKategori()
		case "8":
			tampilkanKategori()
		case "9":
			tambahTransaksi()
		case "10":
			tampilkanTransaksi()
		case "11":
			cariByGejalaSekuensial()
		case "12":
			cariByNamaBinarySearch()
		case "13":
			sortSelectionSort(true)
		case "14":
			sortSelectionSort(false)
		case "15":
			sortInsertionSort(true)
		case "16":
			sortInsertionSort(false)
		case "17":
			sortMultiKriteria()
		case "18":
			statistikStok()
		case "0":
			fmt.Println("\n  Terima kasih telah menggunakan Apotek-Smart. Program selesai.")
			cetakGaris()
			return
		default:
			fmt.Println("  [!] Pilihan tidak valid. Masukkan angka 0–18.")
		}

		fmt.Print("\n  Tekan Enter untuk kembali ke menu...")
		reader.ReadString('\n')
	}
}