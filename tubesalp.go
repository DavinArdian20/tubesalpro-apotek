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

type Obat struct {
	ID          string
	Nama        string
	Gejala      string
	Stok        int
	Harga       float64
	Kadaluarsa  time.Time
}

var daftarObat []Obat

var reader = bufio.NewReader(os.Stdin)

const FORMAT_TANGGAL = "02-01-2006"

func inisialisasiData() {
	dummy := []struct {
		id, nama, gejala string
		stok             int
		harga            float64
		tgl              string
	}{
		{"OBT001", "Paracetamol", "Demam, Sakit Kepala", 50, 5000, "15-12-2025"},
		{"OBT002", "Bodrex", "Demam, Flu, Sakit Kepala", 10, 8500, "01-03-2027"},
		{"OBT003", "OBH Combi", "Batuk, Flu", 30, 15000, "20-06-2026"},
		{"OBT004", "Woods", "Batuk", 25, 18000, "10-09-2026"},
		{"OBT005", "Promag", "Maag, Mual", 20, 12000, "01-01-2027"},
		{"OBT006", "Antangin", "Masuk Angin, Mual, Perut Kembung", 40, 7000, "30-11-2025"},
		{"OBT007", "Tolak Angin", "Masuk Angin, Mual", 35, 9500, "05-07-2026"},
		{"OBT008", "Panadol", "Demam, Sakit Kepala", 15, 11000, "01-05-2027"},
		{"OBT009", "Decolgen", "Flu, Hidung Tersumbat, Demam", 10, 13500, "25-08-2026"},
		{"OBT010", "Mixagrip", "Flu, Sakit Kepala, Demam", 10, 14000, "01-04-2027"},
		{"OBT011", "Ibuprofen", "Sakit Kepala, Nyeri Sendi", 20, 6500, "12-02-2026"},
		{"OBT012", "Antasida", "Maag, Sakit Perut", 18, 8000, "18-07-2026"},
		{"OBT013", "Cetirizine", "Alergi, Gatal-gatal, Flu", 12, 9000, "22-10-2026"},
		{"OBT014", "Amoxicillin", "Infeksi Bakteri, Batuk Berdahak", 8, 25000, "15-01-2026"},
		{"OBT015", "Vitamin C", "Daya Tahan Tubuh, Flu", 60, 4500, "31-12-2027"},
	}

	for _, d := range dummy {
		tgl, _ := time.Parse(FORMAT_TANGGAL, d.tgl)
		daftarObat = append(daftarObat, Obat{
			ID:         d.id,
			Nama:       d.nama,
			Gejala:     d.gejala,
			Stok:       d.stok,
			Harga:      d.harga,
			Kadaluarsa: tgl,
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
		fmt.Println("  [!] Format tanggal tidak valid. Gunakan format dd-mm-yyyy. Coba lagi.")
	}
}

// ============================================================
// UTILITAS TAMPILAN
// ============================================================

func cetakGaris() {
	fmt.Println(strings.Repeat("=", 100))
}

func cetakJudul() {
	cetakGaris()
	fmt.Printf("%-8s %-20s %-35s %-6s %-12s %-12s\n",
		"ID", "Nama Obat", "Gejala", "Stok", "Harga (Rp)", "Kadaluarsa")
	cetakGaris()
}

func cetakBaris(o Obat) {
	gejala := o.Gejala
	if len(gejala) > 33 {
		gejala = gejala[:30] + "..."
	}
	fmt.Printf("%-8s %-20s %-35s %-6d %-12.0f %-12s\n",
		o.ID, o.Nama, gejala, o.Stok,
		o.Harga, o.Kadaluarsa.Format(FORMAT_TANGGAL))
}

func cariIndexByID(id string) int {
	for i, o := range daftarObat {
		if strings.EqualFold(o.ID, id) {
			return i
		}
	}
	return -1
}


func tambahObat() {
	fmt.Println("\n  [ TAMBAH DATA OBAT ]")
	cetakGaris()

	var id string
	for {
		id = bacaInput("  ID Obat        : ")
		if id == "" {
			fmt.Println("  [!] ID tidak boleh kosong.")
			continue
		}
		if cariIndexByID(id) != -1 {
			fmt.Println("  [!] ID sudah digunakan. Gunakan ID lain.")
			continue
		}
		break
	}

	var nama string
	for {
		nama = bacaInput("  Nama Obat      : ")
		if nama == "" {
			fmt.Println("  [!] Nama tidak boleh kosong.")
			continue
		}
		break
	}

	gejala := bacaInput("  Gejala         : ")

	// Baca dan validasi Stok
	var stok int
	for {
		stok = bacaInt("  Stok           : ")
		if stok < 0 {
			fmt.Println("  [!] Stok tidak boleh negatif.")
			continue
		}
		break
	}

	var harga float64
	for {
		harga = bacaFloat("  Harga (Rp)     : ")
		if harga < 0 {
			fmt.Println("  [!] Harga tidak boleh negatif.")
			continue
		}
		break
	}

	kadaluarsa := bacaTanggal("  Tanggal Kadaluarsa")

	obatBaru := Obat{
		ID:         id,
		Nama:       nama,
		Gejala:     gejala,
		Stok:       stok,
		Harga:      harga,
		Kadaluarsa: kadaluarsa,
	}
	daftarObat = append(daftarObat, obatBaru)

	fmt.Println("\n  [✓] Data obat berhasil ditambahkan!")
}

// ============================================================
// FITUR 2: UBAH OBAT
// ============================================================

func ubahObat() {
	fmt.Println("\n  [ UBAH DATA OBAT ]")
	cetakGaris()

	id := bacaInput("  Masukkan ID Obat yang akan diubah: ")
	idx := cariIndexByID(id)
	if idx == -1 {
		fmt.Println("  [!] Obat dengan ID tersebut tidak ditemukan.")
		return
	}

	o := daftarObat[idx]
	fmt.Printf("\n  Data saat ini:\n")
	cetakJudul()
	cetakBaris(o)
	fmt.Println()

	fmt.Println("  Masukkan data baru (tekan Enter untuk tetap menggunakan data lama):")

	// Ubah Nama
	namaBaru := bacaInput(fmt.Sprintf("  Nama Obat [%s]: ", o.Nama))
	if namaBaru != "" {
		daftarObat[idx].Nama = namaBaru
	}

	// Ubah Gejala
	gejalaBaru := bacaInput(fmt.Sprintf("  Gejala [%s]: ", o.Gejala))
	if gejalaBaru != "" {
		daftarObat[idx].Gejala = gejalaBaru
	}

	// Ubah Stok
	stokStr := bacaInput(fmt.Sprintf("  Stok [%d]: ", o.Stok))
	if stokStr != "" {
		stokBaru, err := strconv.Atoi(stokStr)
		if err != nil || stokBaru < 0 {
			fmt.Println("  [!] Input stok tidak valid, stok tidak diubah.")
		} else {
			daftarObat[idx].Stok = stokBaru
		}
	}

	// Ubah Harga
	hargaStr := bacaInput(fmt.Sprintf("  Harga [%.0f]: ", o.Harga))
	if hargaStr != "" {
		hargaBaru, err := strconv.ParseFloat(hargaStr, 64)
		if err != nil || hargaBaru < 0 {
			fmt.Println("  [!] Input harga tidak valid, harga tidak diubah.")
		} else {
			daftarObat[idx].Harga = hargaBaru
		}
	}

	// Ubah Tanggal Kadaluarsa
	tglStr := bacaInput(fmt.Sprintf("  Kadaluarsa [%s] (dd-mm-yyyy): ", o.Kadaluarsa.Format(FORMAT_TANGGAL)))
	if tglStr != "" {
		tglBaru, err := time.Parse(FORMAT_TANGGAL, tglStr)
		if err != nil {
			fmt.Println("  [!] Format tanggal tidak valid, tanggal tidak diubah.")
		} else {
			daftarObat[idx].Kadaluarsa = tglBaru
		}
	}

	fmt.Println("\n  [✓] Data obat berhasil diperbarui!")
}

// ============================================================
// FITUR 3: HAPUS OBAT
// ============================================================

func hapusObat() {
	fmt.Println("\n  [ HAPUS DATA OBAT ]")
	cetakGaris()

	id := bacaInput("  Masukkan ID Obat yang akan dihapus: ")
	idx := cariIndexByID(id)
	if idx == -1 {
		fmt.Println("  [!] Obat dengan ID tersebut tidak ditemukan.")
		return
	}

	fmt.Println("\n  Data yang akan dihapus:")
	cetakJudul()
	cetakBaris(daftarObat[idx])
	fmt.Println()

	konfirmasi := bacaInput("  Apakah Anda yakin ingin menghapus data ini? (y/n): ")
	if strings.ToLower(konfirmasi) != "y" {
		fmt.Println("  [i] Penghapusan dibatalkan.")
		return
	}

	// Hapus elemen dari slice dengan menggabungkan slice sebelum dan sesudah indeks
	daftarObat = append(daftarObat[:idx], daftarObat[idx+1:]...)
	fmt.Println("  [✓] Data obat berhasil dihapus!")
}

// ============================================================
// FITUR 4: TAMPILKAN SEMUA OBAT
// ============================================================

func tampilkanSemua() {
	fmt.Println("\n  [ DAFTAR SEMUA OBAT ]")

	if len(daftarObat) == 0 {
		fmt.Println("  [i] Belum ada data obat.")
		return
	}

	cetakJudul()
	for _, o := range daftarObat {
		cetakBaris(o)
	}
	cetakGaris()
	fmt.Printf("  Total: %d jenis obat\n", len(daftarObat))
}

// ============================================================
// FITUR 5: CARI BERDASARKAN GEJALA
// ============================================================


func cariByGejala() {
	fmt.Println("\n  [ CARI OBAT BERDASARKAN GEJALA ]")
	fmt.Println("  Contoh gejala: Demam, Batuk, Flu, Maag, Sakit Kepala")
	cetakGaris()

	keyword := bacaInput("  Masukkan gejala: ")
	if keyword == "" {
		fmt.Println("  [!] Kata kunci tidak boleh kosong.")
		return
	}

	keywordLower := strings.ToLower(keyword)
	var hasil []Obat
	for _, o := range daftarObat {
		if strings.Contains(strings.ToLower(o.Gejala), keywordLower) {
			hasil = append(hasil, o)
		}
	}

	fmt.Printf("\n  Hasil pencarian untuk gejala \"%s\":\n", keyword)
	if len(hasil) == 0 {
		fmt.Println("  [i] Tidak ada obat yang cocok dengan gejala tersebut.")
		return
	}

	cetakJudul()
	for _, o := range hasil {
		cetakBaris(o)
	}
	cetakGaris()
	fmt.Printf("  Ditemukan: %d obat\n", len(hasil))
}

// ============================================================
// FITUR 6: CARI BERDASARKAN NAMA
// ============================================================

func cariByNama() {
	fmt.Println("\n  [ CARI OBAT BERDASARKAN NAMA ]")
	cetakGaris()

	keyword := bacaInput("  Masukkan nama obat: ")
	if keyword == "" {
		fmt.Println("  [!] Nama obat tidak boleh kosong.")
		return
	}

	keywordLower := strings.ToLower(keyword)
	var hasil []Obat

	// Sequential Search dengan partial match dan case-insensitive
	for _, o := range daftarObat {
		if strings.Contains(strings.ToLower(o.Nama), keywordLower) {
			hasil = append(hasil, o)
		}
	}

	fmt.Printf("\n  Hasil pencarian untuk nama \"%s\":\n", keyword)
	if len(hasil) == 0 {
		fmt.Println("  [i] Tidak ada obat dengan nama tersebut.")
		return
	}

	cetakJudul()
	for _, o := range hasil {
		cetakBaris(o)
	}
	cetakGaris()
	fmt.Printf("  Ditemukan: %d obat\n", len(hasil))
}

// ============================================================
// FITUR 7: SORTING BERDASARKAN TANGGAL KADALUARSA
// ============================================================

func sortKadaluarsa(ascending bool) {
	// Buat salinan slice agar data asli tidak berubah
	hasil := make([]Obat, len(daftarObat))
	copy(hasil, daftarObat)

	n := len(hasil)

	// Selection Sort
	for i := 0; i < n-1; i++ {
		idxPilih := i
		for j := i + 1; j < n; j++ {
			if ascending {
				// Cari tanggal paling dekat (paling kecil)
				if hasil[j].Kadaluarsa.Before(hasil[idxPilih].Kadaluarsa) {
					idxPilih = j
				}
			} else {
				// Cari tanggal paling jauh (paling besar)
				if hasil[j].Kadaluarsa.After(hasil[idxPilih].Kadaluarsa) {
					idxPilih = j
				}
			}
		}
		// Tukar elemen
		hasil[i], hasil[idxPilih] = hasil[idxPilih], hasil[i]
	}

	// Tentukan label tampilan
	arah := "ASCENDING (Kadaluarsa Terdekat → Terjauh)"
	if !ascending {
		arah = "DESCENDING (Kadaluarsa Terjauh → Terdekat)"
	}

	fmt.Printf("\n  [ URUTAN KADALUARSA: %s ]\n", arah)
	cetakJudul()
	for _, o := range hasil {
		cetakBaris(o)
	}
	cetakGaris()
}

func sortMultiKriteria() {
	hasil := make([]Obat, len(daftarObat))
	copy(hasil, daftarObat)
	sort.Slice(hasil, func(i, j int) bool {
		if hasil[i].Stok != hasil[j].Stok {
			return hasil[i].Stok < hasil[j].Stok
		}
		return hasil[i].Kadaluarsa.Before(hasil[j].Kadaluarsa)
	})

	fmt.Println("\n  [ URUTAN BERDASARKAN STOK & KADALUARSA ]")
	fmt.Println("  Prioritas: Stok Terkecil → Stok Terbesar")
	fmt.Println("             Jika stok sama: Kadaluarsa Terdekat → Terjauh")
	cetakJudul()
	for _, o := range hasil {
		cetakBaris(o)
	}
	cetakGaris()
}

func statistikStok() {
	fmt.Println("\n  [ STATISTIK STOK OBAT ]")
	cetakGaris()

	if len(daftarObat) == 0 {
		fmt.Println("  [i] Belum ada data obat.")
		return
	}

	totalStok := 0
	idxTerbanyak := 0
	idxTersedikit := 0
	sekarang := time.Now()
	jumlahKadaluarsa30Hari := 0

	for i, o := range daftarObat {
		totalStok += o.Stok

		// Cari stok terbanyak
		if o.Stok > daftarObat[idxTerbanyak].Stok {
			idxTerbanyak = i
		}

		// Cari stok paling sedikit
		if o.Stok < daftarObat[idxTersedikit].Stok {
			idxTersedikit = i
		}

		// Hitung obat yang kadaluarsa dalam <= 30 hari
		selisih := o.Kadaluarsa.Sub(sekarang)
		if selisih >= 0 && selisih <= 30*24*time.Hour {
			jumlahKadaluarsa30Hari++
		}
	}

	fmt.Printf("  %-40s: %d unit\n", "Total seluruh stok", totalStok)
	fmt.Printf("  %-40s: %s (%d unit)\n", "Obat dengan stok terbanyak",
		daftarObat[idxTerbanyak].Nama, daftarObat[idxTerbanyak].Stok)
	fmt.Printf("  %-40s: %s (%d unit)\n", "Obat dengan stok paling sedikit",
		daftarObat[idxTersedikit].Nama, daftarObat[idxTersedikit].Stok)
	fmt.Printf("  %-40s: %d jenis\n", "Jumlah jenis obat tersedia", len(daftarObat))
	fmt.Printf("  %-40s: %d obat\n", "Obat kadaluarsa dalam <= 30 hari", jumlahKadaluarsa30Hari)
	cetakGaris()

	// Peringatan jika ada obat yang hampir kadaluarsa
	if jumlahKadaluarsa30Hari > 0 {
		fmt.Printf("\n  [!] PERINGATAN: Ada %d obat yang akan kadaluarsa dalam 30 hari ke depan!\n", jumlahKadaluarsa30Hari)
		fmt.Println("  Obat-obat tersebut:")
		for _, o := range daftarObat {
			selisih := o.Kadaluarsa.Sub(sekarang)
			if selisih >= 0 && selisih <= 30*24*time.Hour {
				sisaHari := int(selisih.Hours() / 24)
				fmt.Printf("    - %-20s (kadaluarsa: %s, sisa %d hari)\n",
					o.Nama, o.Kadaluarsa.Format(FORMAT_TANGGAL), sisaHari)
			}
		}
	}
}

// ============================================================
// MENU UTAMA
// ============================================================

func tampilkanMenu() {
	fmt.Println()
	cetakGaris()
	fmt.Println("          SISTEM MANAJEMEN STOK APOTEK          ")
	cetakGaris()
	fmt.Println("   1. Tambah Obat")
	fmt.Println("   2. Ubah Obat")
	fmt.Println("   3. Hapus Obat")
	fmt.Println("   4. Tampilkan Semua Obat")
	fmt.Println("   5. Cari Berdasarkan Gejala")
	fmt.Println("   6. Cari Berdasarkan Nama")
	fmt.Println("   7. Urutkan Kadaluarsa Ascending")
	fmt.Println("   8. Urutkan Kadaluarsa Descending")
	fmt.Println("   9. Urutkan Berdasarkan Stok & Kadaluarsa")
	fmt.Println("  10. Statistik Stok")
	fmt.Println("   0. Keluar")
	cetakGaris()
}

// ============================================================
// FUNGSI UTAMA
// ============================================================

func main() {
	// Buat header awal
	fmt.Println()
	cetakGaris()
	fmt.Println("  Selamat datang di Sistem Manajemen Stok Apotek")
	fmt.Println("  Menginisialisasi data dummy...")
	cetakGaris()

	// Isi data dummy saat program pertama kali dijalankan
	inisialisasiData()
	fmt.Printf("  [✓] %d data obat berhasil dimuat.\n", len(daftarObat))

	// Loop menu utama
	for {
		tampilkanMenu()
		pilihan := bacaInput("  Pilih menu [0-10]: ")

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
			cariByGejala()
		case "6":
			cariByNama()
		case "7":
			sortKadaluarsa(true)
		case "8":
			sortKadaluarsa(false)
		case "9":
			sortMultiKriteria()
		case "10":
			statistikStok()
		case "0":
			fmt.Println("\n  Terima kasih telah menggunakan Sistem Manajemen Stok Apotek.")
			fmt.Println("  Program selesai.")
			cetakGaris()
			return
		default:
			fmt.Println("\n  [!] Pilihan tidak valid. Masukkan angka 0-10.")
		}

		// Jeda sebelum kembali ke menu
		fmt.Print("\n  Tekan Enter untuk kembali ke menu...")
		reader.ReadString('\n')
	}
}