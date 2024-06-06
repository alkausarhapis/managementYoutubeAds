package main

import (
	"fmt"
	"os"
)

const MaxAkun = 100
const MaxVideo = 100

type Video struct {
	Judul         string
	Tayang        int
	Tanggal       string
	SudahDihitung bool
}

type Akun struct {
	Nama        string
	Channel     string
	Email       string
	Saldo       int
	Status      string
	VideoList   [MaxVideo]Video
	JumlahVideo int
}

var akunList [MaxAkun]Akun
var jumlahAkun int

func main() {
	for {
		welcomePage()
		var pilihan string
		fmt.Scan(&pilihan)
		if pilihan == "x" {
			menuUtama()
		} else {
			os.Exit(0)
		}
	}
}

func welcomePage() {
	fmt.Println("==================================================================================================================================")
	fmt.Println("                                        ,.   ,   ,.   .                      ,--,--' ")
	fmt.Println("                                        `|  /|  / ,-. |  ,-. ,-. ,-,-. ,-.   `- | ,-.")
	fmt.Println("                                         | / | /  |-' |  |   | | | | | |-'    , | | |")
	fmt.Println("                                         `'  `'   `-' `' `-' `-' ' ' ' `-'    `-' `-'")
	fmt.Println("    ,.       .    .                                                                   .      .               .                    ")
	fmt.Println("   / |   ,-. |  . | , ,-. ,-. .   ,-,-. ,-. ,-. ,-. . ,-. ,-,-. ,-. ,-.   . . ,-. . . |- . . |-. ,-.   ,-. ,-| ,-. ,-. ,-. ,-. ,-.")
	fmt.Println("  /~~|-. | | |  | |<  ,-| `-. |   | | | ,-| | | ,-| | |-' | | | |-' | |   | | | | | | |  | | | | |-'   ,-| | | `-. |-' | | `-. |-'")
	fmt.Println(",'   `-' |-' `' ' ' ` `-^ `-' '   ' ' ' `-^ ' ' `-^ | `-' ' ' ' `-' ' '   `-| `-' `-^ `' `-^ ^-' `-'   `-^ `-^ `-' `-' ' ' `-' `-'")
	fmt.Println("         |                                          |                      /|                                                     ")
	fmt.Println("         '                                         `'                     `-'                                                     ")
	fmt.Println("===========================================(Tekan 'x' untuk melanjutkan ke menu utama)============================================")
}

func menuUtama() {
	var pilihan int
	for {
		fmt.Println("")
		fmt.Println("========(Manajemen YouTube AdSense)========")
		fmt.Println("1. Manajemen Akun")
		fmt.Println("2. Manajemen Video")
		fmt.Println("3. Hitung Monetisasi")
		fmt.Println("4. Tampilkan Akun dan Video")
		fmt.Println("0. Keluar")
		fmt.Println("===========================================")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			menuAkun()
		case 2:
			menuVideo()
		case 3:
			menuUpdateSaldo()
		case 4:
			tampilkanAkun()
		case 0:
			fmt.Println("Terimakasih sudah menggunakan aplikasi ini.")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

// Submenu akun
func menuAkun() {
	fmt.Println("")
	var pilihan int
	fmt.Println("==============(Manajemen Akun)==============")
	fmt.Println("1. Tambah Akun")
	fmt.Println("2. Cari Akun")
	fmt.Println("3. Edit Akun")
	fmt.Println("4. Hapus Akun")
	fmt.Println("0. Kembali")
	fmt.Println("============================================")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		tambahAkun()
	case 2:
		menuCariAkun()
	case 3:
		editAkun()
	case 4:
		hapusAkun()
	case 0:
		return
	default:
		fmt.Println("Pilihan tidak valid, coba lagi.")
	}
}

// Submenu video
func menuVideo() {
	fmt.Println("")
	var pilihan int
	fmt.Println("==============(Manajemen Video)==============")
	fmt.Println("1. Tambah Video")
	fmt.Println("2. Cari Video")
	fmt.Println("3. Hapus Video")
	fmt.Println("4. Edit Video")
	fmt.Println("0. Kembali")
	fmt.Println("=============================================")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		tambahVideo()
	case 2:
		menuCariVideo()
	case 3:
		hapusVideo()
	case 4:
		editVideo()
	case 0:
		return
	default:
		fmt.Println("Pilihan tidak valid, coba lagi.")
	}
}

// Fungsi untuk menampilkan submenu pencarian akun
func menuCariAkun() {
	fmt.Println("")
	var pilihan int
	fmt.Println("================(Cari Akun)================")
	fmt.Println("1. Cari berdasarkan channel")
	fmt.Println("2. Cari berdasarkan status")
	fmt.Println("0. Kembali")
	fmt.Println("===========================================")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		cariAkunByChannel()
	case 2:
		cariAkunByStatus()
	case 0:
		return
	default:
		fmt.Println("Pilihan tidak valid, coba lagi.")
	}
}

// Fungsi untuk menampilkan submenu pencarian video
func menuCariVideo() {
	fmt.Println("")
	var pilihan int
	fmt.Println("================(Cari Video)================")
	fmt.Println("1. Cari berdasarkan judul")
	fmt.Println("2. Cari berdasarkan viewer")
	fmt.Println("0. Kembali")
	fmt.Println("============================================")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		cariVideoByJudul()
	case 2:
		cariVideoByView()
	case 0:
		return
	default:
		fmt.Println("Pilihan tidak valid, coba lagi.")
	}
}

// Fungsi untuk menambah akun baru
func tambahAkun() {
	var akun Akun
	var channelFound, valid bool
	var i int
	if jumlahAkun >= MaxAkun {
		fmt.Println("Jumlah akun maksimal telah tercapai")
		return
	}

	fmt.Print("Masukkan nama: ")
	fmt.Scan(&akun.Nama)

	channelFound = true
	for channelFound {
		fmt.Print("Masukkan channel: ")
		fmt.Scan(&akun.Channel)
		channelFound = false
		for i = 0; i < jumlahAkun && !channelFound; i++ {
			if akunList[i].Channel == akun.Channel {
				channelFound = true
				fmt.Println("Channel sudah ada. Masukkan channel yang berbeda.")
			}
		}
	}

	fmt.Print("Masukkan email: ")
	fmt.Scan(&akun.Email)
	akun.Saldo = 0

	valid = true
	for valid {
		fmt.Print("Masukkan status (silver/platinum/gold): ")
		fmt.Scan(&akun.Status)
		valid = false

		if akun.Status != "silver" && akun.Status != "platinum" && akun.Status != "gold" {
			valid = true
			fmt.Println("Harap masukkan status dengan benar (silver/platinum/gold).")
		}
	}

	akunList[jumlahAkun] = akun
	jumlahAkun++
	fmt.Println("Akun berhasil ditambahkan")
}

// Fungsi untuk mencari akun berdasarkan channel
func cariAkunByChannel() {
	var channel string
	var index, j int
	var akun Akun
	fmt.Print("Masukkan channel: ")
	fmt.Scan(&channel)

	index = cariAkunBerdasarkanChannel(channel)
	if index != -1 {
		akun = akunList[index]
		fmt.Println("")
		fmt.Printf("Nama: %s | Channel: %s | Email: %s | Saldo: %d | Status: %s| \n",
			akun.Nama, akun.Channel, akun.Email, akun.Saldo, akun.Status)

		fmt.Println("")
		fmt.Printf("%-5s %-50s %-20s %-10s\n", "No", "Judul", "Tayang", "Tanggal")
		for j = 0; j < akun.JumlahVideo; j++ {
			fmt.Printf("%-5d %-50s %-20d %-10s\n", j+1, akun.VideoList[j].Judul, akun.VideoList[j].Tayang, akun.VideoList[j].Tanggal)
		}
		fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("")
	} else {
		fmt.Println("Channel tidak ditemukan")
	}
}

// Fungsi untuk mencari akun berdasarkan status
func cariAkunByStatus() {
	var status string
	var valid bool
	var i, j int

	valid = true
	for valid {
		fmt.Print("Masukkan status (silver/platinum/gold): ")
		fmt.Scan(&status)
		valid = false

		if status != "silver" && status != "platinum" && status != "gold" {
			valid = true
			fmt.Println("Harap masukkan status dengan benar (silver/platinum/gold).")
		}
	}

	for i = 0; i < jumlahAkun; i++ {
		if akunList[i].Status == status {
			fmt.Println("")
			fmt.Printf("Nama: %s | Channel: %s | Email: %s | Saldo: %d | Status: %s| \n",
				akunList[i].Nama, akunList[i].Channel, akunList[i].Email, akunList[i].Saldo, akunList[i].Status)

			fmt.Println("")
			fmt.Printf("%-5s %-50s %-20s %-10s\n", "No", "Judul", "Tayang", "Tanggal")
			for j = 0; j < akunList[i].JumlahVideo; j++ {
				fmt.Printf("%-5d %-50s %-20d %-10s\n", j+1, akunList[i].VideoList[j].Judul, akunList[i].VideoList[j].Tayang, akunList[i].VideoList[j].Tanggal)
			}
			fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------")
			fmt.Println("")
		}
	}
}

// Fungsi untuk mengedit akun
func editAkun() {
	var channel string
	var index int
	var statusValid bool

	fmt.Print("Masukkan channel dari akun yang ingin diedit: ")
	fmt.Scan(&channel)

	index = cariAkunBerdasarkanChannel(channel)
	if index == -1 {
		fmt.Println("Channel tidak ditemukan")
		return
	}

	fmt.Print("Masukkan nama baru: ")
	fmt.Scan(&akunList[index].Nama)

	fmt.Print("Masukkan email baru: ")
	fmt.Scan(&akunList[index].Email)

	statusValid = true
	for statusValid {
		fmt.Print("Masukkan status baru (silver/platinum/gold): ")
		fmt.Scan(&akunList[index].Status)
		statusValid = false

		if akunList[index].Status != "silver" && akunList[index].Status != "platinum" && akunList[index].Status != "gold" {
			statusValid = true
			fmt.Println("Harap masukkan status dengan benar (silver/platinum/gold).")
		}
	}

	fmt.Println("Akun berhasil diedit")
}

// Fungsi untuk menghapus akun
func hapusAkun() {
	var channel string
	var index, i int

	fmt.Print("Masukkan channel dari akun yang ingin dihapus: ")
	fmt.Scan(&channel)

	index = cariAkunBerdasarkanChannel(channel)
	if index == -1 {
		fmt.Println("Channel tidak ditemukan")
		return
	}

	for i = index; i < jumlahAkun-1; i++ {
		akunList[i] = akunList[i+1]
	}

	jumlahAkun--
	fmt.Println("Akun berhasil dihapus")
}

// Fungsi untuk menambah video ke dalam akun
func tambahVideo() {
	var channel string
	var index int

	fmt.Print("Masukkan channel dari akun yang ingin ditambahkan video: ")
	fmt.Scan(&channel)

	index = cariAkunBerdasarkanChannel(channel)
	if index == -1 {
		fmt.Println("Channel tidak ditemukan")
		return
	}

	if akunList[index].JumlahVideo >= MaxVideo {
		fmt.Println("Jumlah video maksimal telah tercapai")
		return
	}

	var video Video
	fmt.Print("Masukkan judul video: ")
	fmt.Scan(&video.Judul)

	fmt.Print("Masukkan jumlah tayang: ")
	fmt.Scan(&video.Tayang)

	fmt.Print("Masukkan tanggal upload (format: YYYY-MM-DD): ")
	fmt.Scan(&video.Tanggal)

	akunList[index].VideoList[akunList[index].JumlahVideo] = video
	akunList[index].JumlahVideo++
	fmt.Println("Video berhasil ditambahkan")
}

// Fungsi untuk mencari video berdasarkan judul dan menampilkan hasilnya
func cariVideoByJudul() {
	var judul string
	var found bool
	var i, index int
	fmt.Print("Masukkan judul video: ")
	fmt.Scan(&judul)

	fmt.Println("")
	fmt.Printf("%-20s %-50s %-10s %-10s\n", "Channel", "Judul", "Tayang", "Tanggal")
	for i = 0; i < jumlahAkun; i++ {
		index = cariVideoBerdasarkanJudul(i, judul)
		if index != -1 {
			fmt.Printf("%-20s %-50s %-10d %-10s\n", akunList[i].Channel, akunList[i].VideoList[index].Judul, akunList[i].VideoList[index].Tayang, akunList[i].VideoList[index].Tanggal)
			found = true
		}
	}
	if !found {
		fmt.Println("")
		fmt.Println("Video dengan judul tersebut tidak ditemukan.")
		fmt.Println("")
	}
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("")
}

// Fungsi untuk mencari video berdasarkan jumlah tayang dengan binary search dan menampilkan hasilnya
func cariVideoByView() {
	var tayang int
	var found bool
	var i, index int

	fmt.Print("Masukkan jumlah tayang: ")
	fmt.Scan(&tayang)

	fmt.Println("")
	fmt.Printf("%-20s %-50s %-10s %-10s\n", "Channel", "Judul", "Tayang", "Tanggal")
	for i = 0; i < jumlahAkun; i++ {
		index = cariVideoBerdasarkanView(i, tayang)
		if index != -1 {
			fmt.Printf("%-20s %-50s %-10d %-10s\n", akunList[i].Channel, akunList[i].VideoList[index].Judul, akunList[i].VideoList[index].Tayang, akunList[i].VideoList[index].Tanggal)
			found = true
		}
	}
	if !found {
		fmt.Println("")
		fmt.Println("Video dengan jumlah tayang tersebut tidak ditemukan.")
		fmt.Println("")
	}
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("")
}

// Fungsi untuk menghapus video dari akun
func hapusVideo() {
	var channel string
	var judul string
	var index, i int

	fmt.Print("Masukkan channel dari akun yang ingin dihapus videonya: ")
	fmt.Scan(&channel)

	index = cariAkunBerdasarkanChannel(channel)
	if index == -1 {
		fmt.Println("Channel tidak ditemukan")
		return
	}

	fmt.Print("Masukkan judul video yang ingin dihapus: ")
	fmt.Scan(&judul)

	var videoIndex = cariVideoBerdasarkanJudul(index, judul)
	if videoIndex == -1 {
		fmt.Println("Video tidak ditemukan")
		return
	}

	for i = videoIndex; i < akunList[index].JumlahVideo-1; i++ {
		akunList[index].VideoList[i] = akunList[index].VideoList[i+1]
	}

	akunList[index].JumlahVideo--
	fmt.Println("Video berhasil dihapus")
}

// Fungsi untuk mengedit video dari akun
func editVideo() {
	var channel string
	var judul string
	var index, videoIndex int

	fmt.Print("Masukkan channel dari akun yang ingin diedit videonya: ")
	fmt.Scan(&channel)

	index = cariAkunBerdasarkanChannel(channel)
	if index == -1 {
		fmt.Println("Channel tidak ditemukan")
		return
	}

	fmt.Print("Masukkan judul video yang ingin diedit: ")
	fmt.Scan(&judul)

	videoIndex = cariVideoBerdasarkanJudul(index, judul)
	if akunList[index].Saldo > 0 {
		akunList[index].Saldo -= hitungMonetisasi(index, videoIndex)
	}
	if videoIndex == -1 {
		fmt.Println("Video tidak ditemukan")
		return
	}

	fmt.Print("Masukkan judul baru: ")
	fmt.Scan(&akunList[index].VideoList[videoIndex].Judul)

	fmt.Print("Masukkan jumlah tayang baru: ")
	fmt.Scan(&akunList[index].VideoList[videoIndex].Tayang)

	akunList[index].VideoList[videoIndex].SudahDihitung = false
	fmt.Println("Video berhasil diedit")
}

// Fungsi untuk menghitung saldo dari video tertentu di akun
func menuUpdateSaldo() {
	var channel string
	var judul string
	var index, videoIndex int

	fmt.Print("Masukkan channel dari akun yang ingin dihitung saldonya: ")
	fmt.Scan(&channel)

	index = cariAkunBerdasarkanChannel(channel)
	if index == -1 {
		fmt.Println("Channel tidak ditemukan")
		return
	}

	fmt.Print("Masukkan judul video yang ingin dihitung monetisasinya: ")
	fmt.Scan(&judul)

	videoIndex = cariVideoBerdasarkanJudul(index, judul)
	if videoIndex == -1 {
		fmt.Println("Video tidak ditemukan")
		return
	}

	if akunList[index].VideoList[videoIndex].SudahDihitung {
		fmt.Println("Monetisasi untuk video ini sudah pernah dihitung.")
		return
	}

	akunList[index].Saldo += hitungMonetisasi(index, videoIndex)
	akunList[index].VideoList[videoIndex].SudahDihitung = true
	fmt.Println("Saldo berhasil diperbarui")
}

// Fungsi untuk menampilkan daftar akun
func tampilkanAkun() {
	for {
		var pilihan, i, j int

		fmt.Println("")
		fmt.Println("==============(Tampilkan data)==============")
		fmt.Println("1. Tampilkan secara ascending")
		fmt.Println("2. Tampilkan secara descending")
		fmt.Println("0. Kembali")
		fmt.Println("============================================")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 || pilihan == 2 || pilihan == 0 {
			for i = 0; i < jumlahAkun; i++ {
				if pilihan == 1 {
					insertionSortAscending(&akunList[i].VideoList, akunList[i].JumlahVideo)
				} else if pilihan == 2 {
					selectionSortDescending(&akunList[i].VideoList, akunList[i].JumlahVideo)
				} else if pilihan == 0 {
					return
				}
				fmt.Println("")
				fmt.Printf("Nama: %s | Channel: %s | Email: %s | Saldo: %d | Status: %s| \n",
					akunList[i].Nama, akunList[i].Channel, akunList[i].Email, akunList[i].Saldo, akunList[i].Status)

				fmt.Println("")
				fmt.Printf("%-50s %-20s %-10s\n", "Judul", "Tayang", "Tanggal")
				for j = 0; j < akunList[i].JumlahVideo; j++ {
					fmt.Printf("%-50s %-20d %-10s\n", akunList[i].VideoList[j].Judul, akunList[i].VideoList[j].Tayang, akunList[i].VideoList[j].Tanggal)
				}
				fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------")
				fmt.Println("")
			}
			return
		} else {
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}

// Fungsi untuk mencari akun berdasarkan channel
func cariAkunBerdasarkanChannel(channel string) int {
	var i int
	for i = 0; i < jumlahAkun; i++ {
		if akunList[i].Channel == channel {
			return i
		}
	}
	return -1
}

// Fungsi untuk mencari video berdasarkan judul
func cariVideoBerdasarkanJudul(indexAkun int, judul string) int {
	var i int
	for i = 0; i < akunList[indexAkun].JumlahVideo; i++ {
		if akunList[indexAkun].VideoList[i].Judul == judul {
			return i
		}
	}
	return -1
}

// Fungsi untuk mencari video berdasarkan view (gunakan binary search)
func cariVideoBerdasarkanView(indexAkun int, tayang int) int {
	var left, right, mid int
	selectionSortByView(&akunList[indexAkun].VideoList, akunList[indexAkun].JumlahVideo)

	left = 0
	right = akunList[indexAkun].JumlahVideo - 1
	for left <= right {
		mid = (left + right) / 2
		if akunList[indexAkun].VideoList[mid].Tayang == tayang {
			return mid
		} else if akunList[indexAkun].VideoList[mid].Tayang < tayang {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// Fungsi untuk mengurutkan video secara ascending berdasarkan view
func selectionSortByView(videoList *[MaxVideo]Video, jumlahVideo int) {
	var i, j, minIndex int
	var temp Video

	for i = 0; i < jumlahVideo-1; i++ {
		minIndex = i
		for j = i + 1; j < jumlahVideo; j++ {
			if videoList[j].Tayang < videoList[minIndex].Tayang {
				minIndex = j
			}
		}

		temp = videoList[i]
		videoList[i] = videoList[minIndex]
		videoList[minIndex] = temp
	}
}

// Fungsi untuk mengurutkan video secara ascending berdasarkan tanggal
func insertionSortAscending(videoList *[MaxVideo]Video, jumlahVideo int) {
	var i, j int
	var temp Video

	for i = 1; i < jumlahVideo; i++ {
		temp = videoList[i]
		j = i - 1
		for j >= 0 && videoList[j].Tanggal > temp.Tanggal {
			videoList[j+1] = videoList[j]
			j = j - 1
		}
		videoList[j+1] = temp
	}
}

// Fungsi untuk mengurutkan video secara descending berdasarkan tanggal
func selectionSortDescending(videoList *[MaxVideo]Video, jumlahVideo int) {
	var i, j, maxIndex int
	var temp Video

	for i = 0; i < jumlahVideo-1; i++ {
		maxIndex = i
		for j = i + 1; j < jumlahVideo; j++ {
			if videoList[j].Tanggal > videoList[maxIndex].Tanggal {
				maxIndex = j
			}
		}

		temp = videoList[i]
		videoList[i] = videoList[maxIndex]
		videoList[maxIndex] = temp
	}
}

// Fungsi untuk menghitung monetisasi
func hitungMonetisasi(indexAkun, indexVideo int) int {
	var rate int

	if akunList[indexAkun].Status == "silver" {
		rate = 1000
	} else if akunList[indexAkun].Status == "platinum" {
		rate = 2000
	} else if akunList[indexAkun].Status == "gold" {
		rate = 5000
	} else {
		rate = 0
	}

	return akunList[indexAkun].VideoList[indexVideo].Tayang * rate
}
