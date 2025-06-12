package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Konten struct {
	Judul    string
	Hashtag  string
	Platform string
	Viewers  int
	Tanggal  string
}

var daftarKonten = []Konten{
	{"Tips_Algoritma", "#ngoding", "Instagram", 1500, "01-05-2025"},
	{"Belajar_GoLang", "#golang", "YouTube", 2300, "03-05-2025"},
	{"UI/UX_Design", "#design", "TikTok", 1800, "05-05-2025"},
	{"Digital_Marketing", "#marketing", "Instagram", 2000, "07-05-2025"},
	{"AI_untuk_Pemula", "#ai", "YouTube", 3200, "10-05-2025"},
}

func clear() {
	if os.PathSeparator == '\\' {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	var awal int
	running := true
	for running {
		clear()
		fmt.Println("=== Aplikasi AI Pembuat Konten Sosial Media ===")
		fmt.Println("1. Lihat Semua Konten")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&awal)

		if awal == 1 {
			clear()
			tampilkanSemua()
			menuLengkap()
		} else if awal == 0 {
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			running = false
		} else {
			fmt.Println("Pilihan tidak valid.")
			fmt.Println("Tekan enter untuk melanjutkan...")
			fmt.Scanln()
			fmt.Scanln()
		}
	}
}

func menuLengkap() {
	for {
		fmt.Println("\n--- Menu Lengkap ---")
		fmt.Println("1. Tambah Ide Konten")
		fmt.Println("2. Edit Ide Konten")
		fmt.Println("3. Hapus Ide Konten")
		fmt.Println("4. Cari Ide (Sequential/Binary Search)")
		fmt.Println("5. Urutkan Konten (Selection/Insertion Sort)")
		fmt.Println("6. Evaluasi Viewers Tertinggi")
		fmt.Println("7. Rekomendasi Caption & Hashtag")
		fmt.Println("0. Kembali ke Menu Awal")
		fmt.Print("Pilih: ")
		var pilihan int
		fmt.Scan(&pilihan)

		clear()
		tampilkanSemua()

		switch pilihan {
		case 1:
			tambahKonten()
			fmt.Println("\nTekan enter untuk kembali ke menu...")
			fmt.Scanln()
			fmt.Scanln()
		case 2:
			editKonten()
			fmt.Println("\nTekan enter untuk kembali ke menu...")
			fmt.Scanln()
			fmt.Scanln()
		case 3:
			hapusKonten()
			fmt.Println("\nTekan enter untuk kembali ke menu...")
			fmt.Scanln()
			fmt.Scanln()
		case 4:
			menuPencarian()
		case 5:
			menuPengurutan()
			fmt.Println("\nTekan enter untuk kembali ke menu...")
			fmt.Scanln()
			fmt.Scanln()
		case 6:
			evaluasiViewers()
			fmt.Println("\nTekan enter untuk kembali ke menu...")
			fmt.Scanln()
			fmt.Scanln()
		case 7:
			rekomendasiCaptionHashtag()
			fmt.Println("\nTekan enter untuk kembali ke menu...")
			fmt.Scanln()
			fmt.Scanln()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
		clear()
		tampilkanSemua()
	}
}

func tampilkanSemua() {
	if len(daftarKonten) == 0 {
		fmt.Println("Belum ada data untuk ditampilkan.")
		return
	}
	fmt.Println("---------------------------------------------------------------------------------------------")
	fmt.Printf("| %-4s | %-20s | %-15s | %-15s | %-10s | %-12s |\n", "No", "Judul", "Hashtag", "Platform", "Viewers", "Tanggal")
	fmt.Println("---------------------------------------------------------------------------------------------")
	for i, k := range daftarKonten {
		fmt.Printf("| %-4d | %-20s | %-15s | %-15s | %10d | %-12s |\n", i+1, k.Judul, k.Hashtag, k.Platform, k.Viewers, k.Tanggal)
	}
	fmt.Println("---------------------------------------------------------------------------------------------")
}

func tambahKonten() {
	var k Konten
	fmt.Print("Judul: ")
	fmt.Scan(&k.Judul)
	fmt.Print("Hashtag: ")
	fmt.Scan(&k.Hashtag)
	fmt.Print("Platform: ")
	fmt.Scan(&k.Platform)
	fmt.Print("Viewers: ")
	fmt.Scan(&k.Viewers)
	fmt.Print("Tanggal (dd-mm-yyyy): ")
	fmt.Scan(&k.Tanggal)
	daftarKonten = append(daftarKonten, k)
	fmt.Println("Konten berhasil ditambahkan.")
}

func editKonten() {
	if len(daftarKonten) == 0 {
		fmt.Println("Belum ada konten untuk diedit.")
		return
	}
	var nomor int
	fmt.Print("Masukkan nomor konten yang ingin diedit: ")
	fmt.Scan(&nomor)

	if nomor < 1 || nomor > len(daftarKonten) {
		fmt.Println("Nomor tidak valid.")
		return
	}

	i := nomor - 1
	fmt.Printf("Edit konten nomor %d:\n", nomor)
	fmt.Printf("Judul lama: %s\n", daftarKonten[i].Judul)
	fmt.Print("Judul baru: ")
	fmt.Scan(&daftarKonten[i].Judul)

	fmt.Printf("Hashtag lama: %s\n", daftarKonten[i].Hashtag)
	fmt.Print("Hashtag baru: ")
	fmt.Scan(&daftarKonten[i].Hashtag)

	fmt.Printf("Platform lama: %s\n", daftarKonten[i].Platform)
	fmt.Print("Platform baru: ")
	fmt.Scan(&daftarKonten[i].Platform)

	fmt.Printf("Viewers lama: %d\n", daftarKonten[i].Viewers)
	fmt.Print("Viewers baru: ")
	fmt.Scan(&daftarKonten[i].Viewers)

	fmt.Printf("Tanggal lama: %s\n", daftarKonten[i].Tanggal)
	fmt.Print("Tanggal baru (dd-mm-yyyy): ")
	fmt.Scan(&daftarKonten[i].Tanggal)

	fmt.Println("Konten berhasil diedit.")
}

func hapusKonten() {
	if len(daftarKonten) == 0 {
		fmt.Println("Belum ada konten untuk dihapus.")
		return
	}
	var nomor int
	fmt.Print("Masukkan nomor konten yang ingin dihapus: ")
	fmt.Scan(&nomor)

	if nomor < 1 || nomor > len(daftarKonten) {
		fmt.Println("Nomor tidak valid.")
		return
	}

	i := nomor - 1
	daftarKonten = append(daftarKonten[:i], daftarKonten[i+1:]...)
	fmt.Println("Konten berhasil dihapus.")
}

func sequentialSearch(keyword string) {
	keyword = strings.ToLower(keyword)
	found := false
	for _, k := range daftarKonten {
		judulLower := strings.ToLower(k.Judul)
		hashtagLower := strings.ToLower(k.Hashtag)
		if strings.Contains(judulLower, keyword) || strings.Contains(hashtagLower, keyword) {
			fmt.Println(k)
			found = true
		}
	}
	if !found {
		fmt.Println("Konten tidak ditemukan.")
	}
}

func binarySearch(keyword string) {
	insertionSortByJudul() // Pastikan data sudah urut berdasarkan judul
	keyword = strings.ToLower(keyword)

	low, high := 0, len(daftarKonten)-1
	found := false
	for low <= high {
		mid := (low + high) / 2
		judulLower := strings.ToLower(daftarKonten[mid].Judul)

		if judulLower == keyword {
			fmt.Println(daftarKonten[mid])
			found = true
			break
		} else if judulLower < keyword {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !found {
		fmt.Println("Konten tidak ditemukan (binary search hanya cocok untuk exact match pada judul).")
	}
}

func menuPencarian() {
	clear()
	tampilkanSemua()
	fmt.Println("\n--- Pencarian Ide Konten ---")
	var keyword string
	var metode int
	fmt.Print("Masukkan kata kunci (dapat sebagian kata dari judul atau hashtag): ")
	fmt.Scan(&keyword)
	fmt.Println("1. Sequential Search (partial match)")
	fmt.Println("2. Binary Search (exact match pada judul saja)")
	fmt.Print("Pilih metode: ")
	fmt.Scan(&metode)

	clear()
	tampilkanSemua()

	fmt.Println("\n--- Hasil Pencarian ---")
	if metode == 1 {
		sequentialSearch(keyword)
	} else if metode == 2 {
		binarySearch(keyword)
	} else {
		fmt.Println("Metode pencarian tidak valid.")
	}

	fmt.Println("\nTekan enter untuk kembali ke menu...")
	fmt.Scanln()
	fmt.Scanln()

	clear()
	tampilkanSemua()
}

func selectionSortByViewers() {
	n := len(daftarKonten)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if daftarKonten[j].Viewers < daftarKonten[minIdx].Viewers {
				minIdx = j
			}
		}
		daftarKonten[i], daftarKonten[minIdx] = daftarKonten[minIdx], daftarKonten[i]
	}
}

func insertionSortByJudul() {
	for i := 1; i < len(daftarKonten); i++ {
		key := daftarKonten[i]
		j := i - 1
		for j >= 0 && daftarKonten[j].Judul > key.Judul {
			daftarKonten[j+1] = daftarKonten[j]
			j--
		}
		daftarKonten[j+1] = key
	}
}

func menuPengurutan() {
	var metode int
	fmt.Println("1. Urutkan berdasarkan Viewers (Selection Sort)")
	fmt.Println("2. Urutkan berdasarkan Judul (Insertion Sort)")
	fmt.Print("Pilih metode: ")
	fmt.Scan(&metode)
	if metode == 1 {
		selectionSortByViewers()
	} else if metode == 2 {
		insertionSortByJudul()
	} else {
		fmt.Println("Metode tidak valid.")
	}
}

func evaluasiViewers() {
	if len(daftarKonten) == 0 {
		fmt.Println("Belum ada konten.")
		return
	}
	max := daftarKonten[0]
	for _, k := range daftarKonten {
		if k.Viewers > max.Viewers {
			max = k
		}
	}
	fmt.Println("Konten dengan Viewers tertinggi:")
	fmt.Println(max)
}

func rekomendasiCaptionHashtag() {
	if len(daftarKonten) == 0 {
		fmt.Println("Belum ada data konten.")
		return
	}
	top := daftarKonten[0]
	for _, konten := range daftarKonten {
		if konten.Viewers > top.Viewers {
			top = konten
		}
	}
	fmt.Println("===== Rekomendasi Berdasarkan Viewers Tertinggi =====")
	fmt.Printf("Judul: %s\n", top.Judul)
	fmt.Printf("Hashtag: %s\n", top.Hashtag)
	fmt.Printf("Platform: %s\n", top.Platform)
	fmt.Printf("Viewers: %d\n", top.Viewers)
	fmt.Printf("Tanggal: %s\n", top.Tanggal)
	fmt.Println("\nRekomendasi Caption:")
	fmt.Printf("\"%s adalah konten dengan viewers %d, sangat direkomendasikan untuk ide konten kamu!\"\n", top.Judul, top.Viewers)
	fmt.Println("\nRekomendasi Hashtag Tambahan:")
	fmt.Println(top.Hashtag)
}
