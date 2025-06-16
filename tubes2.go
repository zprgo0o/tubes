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

type Ide struct {
	ID          int
	Judul       string
	Deskripsi   string
	Kategori    string
	Tanggal     time.Time
	VotePositif int
}

var daftarIde []Ide
var idBerikutnya int = 1

func input(prompt string) string {
	fmt.Print(prompt)
	pembaca := bufio.NewReader(os.Stdin)
	teks, _ := pembaca.ReadString('\n')
	return strings.TrimSpace(teks)
}

func tambahIde() {
	judul := input("Judul ide: ")
	deskripsi := input("Deskripsi singkat: ")
	kategori := input("Kategori (Produk/Marketing/Jasa): ")
	ideBaru := Ide{
		ID:          idBerikutnya,
		Judul:       judul,
		Deskripsi:   deskripsi,
		Kategori:    kategori,
		Tanggal:     time.Now(),
		VotePositif: 0,
	}
	idBerikutnya++
	daftarIde = append(daftarIde, ideBaru)
	fmt.Println("✅ Ide berhasil ditambahkan!\n")
}

func lihatIde() {
	if len(daftarIde) == 0 {
		fmt.Println("Belum ada ide yang tercatat.\n")
		return
	}
	for _, ide := range daftarIde {
		fmt.Printf("[ID: %d] %s\nKategori: %s | Tanggal: %s\nDeskripsi: %s\n👍 %d\n\n",
			ide.ID, ide.Judul, ide.Kategori, ide.Tanggal.Format("02-01-2006"), ide.Deskripsi, ide.VotePositif)
	}
}

func voteIde() {
	idStr := input("Masukkan ID ide yang ingin diberi vote: ")
	id, _ := strconv.Atoi(idStr)
	for i := range daftarIde {
		if daftarIde[i].ID == id {
			daftarIde[i].VotePositif++
			fmt.Println("✅ Vote dicatat!\n")
			return
		}
	}
	fmt.Println("❌ ID tidak ditemukan.\n")
}

func cariIdeSequential(kataKunci string) {
	adaIde := false
	for _, ide := range daftarIde {
		if strings.Contains(strings.ToLower(ide.Judul), strings.ToLower(kataKunci)) ||
			strings.Contains(strings.ToLower(ide.Deskripsi), strings.ToLower(kataKunci)) {
			fmt.Printf("[ID: %d] %s\nKategori: %s | Tanggal: %s\nDeskripsi: %s\n👍 %d\n\n",
				ide.ID, ide.Judul, ide.Kategori, ide.Tanggal.Format("02-01-2006"), ide.Deskripsi, ide.VotePositif)
			adaIde = true
		}
	}
	if !adaIde {
		fmt.Println("❌ Tidak ada ide yang cocok dengan kata kunci tersebut.\n")
	}
}

func cariIdeBinary(kataKunci string) {
	sort.Slice(daftarIde, func(i, j int) bool {
		return strings.ToLower(daftarIde[i].Judul) < strings.ToLower(daftarIde[j].Judul)
	})
	low, high := 0, len(daftarIde)-1
	for low <= high {
		mid := (low + high) / 2
		if strings.Contains(strings.ToLower(daftarIde[mid].Judul), strings.ToLower(kataKunci)) {
			fmt.Printf("[ID: %d] %s\nKategori: %s | Tanggal: %s\nDeskripsi: %s\n👍 %d\n\n",
				daftarIde[mid].ID, daftarIde[mid].Judul, daftarIde[mid].Kategori, daftarIde[mid].Tanggal.Format("02-01-2006"), daftarIde[mid].Deskripsi, daftarIde[mid].VotePositif)
			return
		} else if strings.ToLower(daftarIde[mid].Judul) < strings.ToLower(kataKunci) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("❌ Tidak ada ide yang cocok dengan kata kunci tersebut.\n")
}

func urutkanIdeBerdasarkanUpvote() {
	sort.Slice(daftarIde, func(i, j int) bool {
		return daftarIde[i].VotePositif > daftarIde[j].VotePositif
	})
	fmt.Println("✅ Ide telah diurutkan berdasarkan jumlah upvote!\n")
}

func lihatIdePopuler(periodHari int) {
	batasWaktu := time.Now().AddDate(0, 0, -periodHari)
	populer := []Ide{}
	for _, ide := range daftarIde {
		if ide.Tanggal.After(batasWaktu) {
			populer = append(populer, ide)
		}
	}
	if len(populer) == 0 {
		fmt.Println("❌ Tidak ada ide yang populer dalam periode tersebut.\n")
		return
	}
	urutkanIdeBerdasarkanUpvote()
	for _, ide := range populer {
		fmt.Printf("[ID: %d] %s\nKategori: %s | Tanggal: %s\nDeskripsi: %s\n👍 %d\n\n",
			ide.ID, ide.Judul, ide.Kategori, ide.Tanggal.Format("02-01-2006"), ide.Deskripsi, ide.VotePositif)
	}
}

func main() {
	for {
		fmt.Println("=== IdeaBox - Pengelolaan Ide Startup ===")
		fmt.Println("1. Lihat Daftar Ide")
		fmt.Println("2. Tambah Ide Baru")
		fmt.Println("3. Voting Ide")
		fmt.Println("4. Cari Ide (Sequential Search)")
		fmt.Println("5. Cari Ide (Binary Search)")
		fmt.Println("6. Urutkan Ide Berdasarkan Upvote")
		fmt.Println("7. Lihat Ide Populer")
		fmt.Println("8. Keluar")

		pilih := input("Pilih menu: ")

		switch pilih {
		case "1":
			lihatIde()
		case "2":
			tambahIde()
		case "3":
			voteIde()
		case "4":
			kataKunci := input("Masukkan kata kunci pencarian: ")
			cariIdeSequential(kataKunci)
		case "5":
			kataKunci := input("Masukkan kata kunci pencarian: ")
			cariIdeBinary(kataKunci)
		case "6":
			urutkanIdeBerdasarkanUpvote()
		case "7":
			periodHariStr := input("Masukkan periode hari untuk ide populer: ")
			periodHari, _ := strconv.Atoi(periodHariStr)
			lihatIdePopuler(periodHari)
		case "8":
			fmt.Println("👋 Terima kasih, sampai jumpa!")
			return
		default:
			fmt.Println("❌ Menu tidak tersedia.\n")
		}
	}
}
