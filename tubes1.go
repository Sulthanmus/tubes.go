package main

import "fmt"

type Konten struct {
	Judul      string
	Hashtag    string
	Platform   string
	Viewers int
	Tanggal    string
}

var daftarKonten []Konten

func main() {
	var pilihan int

	for {
		fmt.Println("\n=== Aplikasi AI Pembuat Konten Sosial Media ===")
		fmt.Println("1. Tambah Ide Konten")
		fmt.Println("2. Edit Ide Konten")
		fmt.Println("3. Hapus Ide Konten")
		fmt.Println("4. Cari Ide (Sequential/Binary Search)")
		fmt.Println("5. Urutkan Konten (Selection/Insertion Sort)")
		fmt.Println("6. Evaluasi Viewers Tertinggi")
		fmt.Println("7. Tampilkan Semua Konten")
		fmt.Println("8. Rekomendasi Caption & Hashtag")
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahKonten()
		case 2:
			editKonten()
		case 3:
			hapusKonten()
		case 4:
			menuPencarian()
		case 5:
			menuPengurutan()
		case 6:
			evaluasiViewers()
		case 7:
			tampilkanSemua()
		case 8:
			rekomendasiCaptionHashtag()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
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
	fmt.Print("Tanggal (yyyy-mm-dd): ")
	fmt.Scan(&k.Tanggal)
	daftarKonten = append(daftarKonten, k)
	fmt.Println("Konten berhasil ditambahkan.")
}

func editKonten() {
	var judul string
	fmt.Print("Judul konten yang ingin diedit: ")
	fmt.Scan(&judul)
	for i := range daftarKonten {
		if daftarKonten[i].Judul == judul {
			fmt.Print("Judul baru: ")
			fmt.Scan(&daftarKonten[i].Judul)
			fmt.Print("Hashtag baru: ")
			fmt.Scan(&daftarKonten[i].Hashtag)
			fmt.Print("Platform baru: ")
			fmt.Scan(&daftarKonten[i].Platform)
			fmt.Print("Viewers baru: ")
			fmt.Scan(&daftarKonten[i].Viewers)
			fmt.Print("Tanggal baru: ")
			fmt.Scan(&daftarKonten[i].Tanggal)
			fmt.Println("Konten berhasil diedit.")
			return
		}
	}
	fmt.Println("Konten tidak ditemukan.")
}

func hapusKonten() {
	var judul string
	fmt.Print("Judul konten yang ingin dihapus: ")
	fmt.Scan(&judul)
	for i := range daftarKonten {
		if daftarKonten[i].Judul == judul {
			daftarKonten = append(daftarKonten[:i], daftarKonten[i+1:]...)
			fmt.Println("Konten berhasil dihapus.")
			return
		}
	}
	fmt.Println("Konten tidak ditemukan.")
}

func sequentialSearch(keyword string) {
	found := false
	for _, k := range daftarKonten {
		if k.Judul == keyword {
			fmt.Println(k)
			found = true
		}
	}
	if !found {
		fmt.Println("Konten tidak ditemukan.")
	}
}

func binarySearch(keyword string) {
	insertionSortByJudul()
	low, high := 0, len(daftarKonten)-1
	for low <= high {
		mid := (low + high) / 2
		if daftarKonten[mid].Judul == keyword {
			fmt.Println(daftarKonten[mid])
			return
		} else if daftarKonten[mid].Judul < keyword {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("Konten tidak ditemukan.")
}

func menuPencarian() {
	var keyword string
	var metode int
	fmt.Print("Masukkan kata kunci: ")
	fmt.Scan(&keyword)
	fmt.Println("1. Sequential Search\n2. Binary Search")
	fmt.Print("Pilih metode: ")
	fmt.Scan(&metode)
	if metode == 1 {
		sequentialSearch(keyword)
	} else {
		binarySearch(keyword)
	}
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
	} else {
		insertionSortByJudul()
	}
	tampilkanSemua()
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

	fmt.Println("\n Rekomendasi Caption:")
	fmt.Printf("“%s adalah konten dengan viewers %d, sehingga dapat menjadi rekomendasi konten kamu!”\n", top.Judul, top.Viewers)

	fmt.Println("\n Rekomendasi Hashtag Tambahan:")
	fmt.Println(top.Hashtag)
}

func tampilkanSemua() {
	for _, k := range daftarKonten {
		fmt.Printf("Judul: %s | Hashtag: %s | Platform: %s | Viewers: %d | Tanggal: %s\n",
			k.Judul, k.Hashtag, k.Platform, k.Viewers, k.Tanggal)
	}
}
