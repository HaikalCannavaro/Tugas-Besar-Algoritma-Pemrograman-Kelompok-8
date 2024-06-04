package main

import "fmt"

const NMAX = 100

type dataBuku struct {
	judul     string
	pengarang string
	ISBN      int
	eksemplar int
	tersedia  bool
}

type dataPeminjam struct {
	bukuDipinjam dataBuku
	nama         string
	nomorTelp    string
	lamaPinjaman int
}

func tambahBuku(data *[NMAX]dataBuku, sumBuku *int) {
	var n, i int
	fmt.Print("Masukkan banyak buku yang ingin ditambah : ")
	fmt.Scan(&n)
	fmt.Println()

	if *sumBuku+n > NMAX {
		fmt.Println("Perpustakaan penuh")
	} else {
		for i = 0; i <= n-1; i++ {
			fmt.Printf("Masukkan judul buku ke-%d : ", i+1)
			fmt.Scan(&data[i].judul)
			fmt.Printf("Masukkan pengarang buku ke-%d : ", i+1)
			fmt.Scan(&data[i].pengarang)
			fmt.Printf("Masukkan nomor ISBN buku ke-%d: ", i+1)
			fmt.Scan(&data[i].ISBN)
			fmt.Printf("Masukkan banyak eksemplar buku ke-%d : ", i+1)
			fmt.Scan(&data[i].eksemplar)
			data[i].tersedia = true
			*sumBuku++
			fmt.Println()
		}
	}
}

func cariJudulBuku(data *[NMAX]dataBuku, sumBuku int, inputanJudul string) {
	var found bool = false
	var i int

	for i = 0; i <= sumBuku-1; i++ {
		if data[i].judul == inputanJudul && i <= sumBuku-1 {
			fmt.Println("**Buku ditemukan**")
			fmt.Printf("Judul buku : %s \n", data[i].judul)
			fmt.Printf("Nama pengarangnya : %s \n", data[i].pengarang)
			fmt.Printf("Nomor ISBN : %d \n", data[i].ISBN)
			fmt.Printf("Banyak eksemplar : %d \n", data[i].eksemplar)
			fmt.Println()
			found = true
		}
	}
	if found == false {
		fmt.Println("**Buku tidak ditemukan**")
	}
}

func cariPengarangBuku(data *[NMAX]dataBuku, sumBuku int, inputanPengarang string) {
	var found bool = false
	var i int

	for i = 0; i <= sumBuku-1; i++ {
		if data[i].pengarang == inputanPengarang {
			fmt.Println("**Buku ditemukan**")
			fmt.Printf("Judul buku : %s \n", data[i].judul)
			fmt.Printf("Nama pengarangnya : %s \n", data[i].pengarang)
			fmt.Printf("Nomor ISBN : %d \n", data[i].ISBN)
			fmt.Printf("Banyak eksemplar : %d \n", data[i].eksemplar)
			fmt.Println()
			found = true
		}
	}
	if found == false || sumBuku <= 0 {
		fmt.Println("**Buku tidak ditemukan**")
	}
}

func hapusData(data *[NMAX]dataBuku, sumBuku *int, hapusan string) {
	var found bool = false
	var newData [NMAX]dataBuku
	var i, j int

	for i = 0; i <= *sumBuku-1; i++ {
		if data[i].judul != hapusan {
			newData[j] = data[i]
			j++
			found = true
		}
	}
	*data = newData
	*sumBuku = j
	fmt.Println("**Buku berhasil dihapus**")

	if found == false {
		fmt.Println("**Buku tidak ditemukan**")
	}
}

func editData(data *[NMAX]dataBuku, sumBuku *int, inputEdit string) {
	var found bool = false
	var i int
	if *sumBuku == 0 {
		fmt.Println("**Tidak ada buku di perpustakaan**")
	}
	for i = 0; i <= *sumBuku-1; i++ {
		if data[i].judul == inputEdit {
			fmt.Print("Masukkan judul baru : ")
			fmt.Scan(&data[i].judul)
			fmt.Print("Masukkan nama pengarang baru : ")
			fmt.Scan(&data[i].pengarang)
			fmt.Print("Masukkan nomor ISBN baru : ")
			fmt.Scan(&data[i].ISBN)
			fmt.Print("Masukkan jumlah eksemplar baru : ")
			fmt.Scan(&data[i].eksemplar)
			fmt.Println()
			fmt.Println("**Buku berhasil diedit**")
			fmt.Println()
			found = true
		}
	}
	if found == false {
		fmt.Println("**Buku tidak ditemukan**")
	}
}

func sortingJudul(data [NMAX]dataBuku, sumBuku int) [NMAX]dataBuku {
	var pass, idx, i int
	var temp dataBuku
	var dataBaru [NMAX]dataBuku
	dataBaru = data
	pass = 1
	for pass <= sumBuku-1 {
		idx = pass - 1
		i = pass
		for i < sumBuku {
			if dataBaru[idx].judul > dataBaru[i].judul {
				idx = i
			}
			i++
		}
		temp = dataBaru[pass-1]
		dataBaru[pass-1] = dataBaru[idx]
		dataBaru[idx] = temp
		pass++
	}
	return dataBaru
}

func sortingPengarang(data [NMAX]dataBuku, sumBuku int) [NMAX]dataBuku {
	var pass, idx, i int
	var temp dataBuku
	var dataBaru [NMAX]dataBuku
	dataBaru = data
	pass = 1
	for pass <= sumBuku-1 {
		idx = pass - 1
		i = pass
		for i < sumBuku {
			if dataBaru[idx].pengarang > dataBaru[i].pengarang {
				idx = i
			}
			i++
		}
		temp = dataBaru[pass-1]
		dataBaru[pass-1] = dataBaru[idx]
		dataBaru[idx] = temp
		pass++
	}
	return dataBaru
}

func tampilDataJudul(data *[NMAX]dataBuku, sumBuku *int) {
	var dataSort [NMAX]dataBuku
	var i int
	dataSort = sortingJudul(*data, *sumBuku)

	if *sumBuku == 0 {
		fmt.Println("**Tidak ada buku di perpustakaan**")
	}
	fmt.Println()
	fmt.Printf("## Terdapat %d buku dalam perpustakaan ##", *sumBuku)
	for i = 0; i <= *sumBuku-1; i++ {
		fmt.Println()
		fmt.Printf("**Buku ke-%d**\n", i+1)
		fmt.Println("Judul : ", dataSort[i].judul)
		fmt.Println("Nama Pengarang : ", dataSort[i].pengarang)
		fmt.Println("Nomor ISBN : ", dataSort[i].ISBN)
		fmt.Println("Jumlah Eksemplar : ", dataSort[i].eksemplar)
		if dataSort[i].tersedia {
			fmt.Println("Buku tersedia")
		} else {
			fmt.Println("Buku sedang dipinjam")
		}
	}
}

func tampilDataPengarang(data *[NMAX]dataBuku, sumBuku *int) {
	var dataSort [NMAX]dataBuku
	var i int
	dataSort = sortingPengarang(*data, *sumBuku)

	if *sumBuku == 0 {
		fmt.Println("**Tidak ada buku di perpustakaan**")
		return
	}
	fmt.Println()
	fmt.Printf("## Terdapat %d buku dalam perpustakaan ##", *sumBuku)
	for i = 0; i <= *sumBuku-1; i++ {
		fmt.Println()
		fmt.Printf("**Buku ke-%d**\n", i+1)
		fmt.Println("Judul : ", dataSort[i].judul)
		fmt.Println("Nama Pengarang : ", dataSort[i].pengarang)
		fmt.Println("Nomor ISBN : ", dataSort[i].ISBN)
		fmt.Println("Jumlah Eksemplar : ", dataSort[i].eksemplar)
		if dataSort[i].tersedia {
			fmt.Println("Buku tersedia")
		} else {
			fmt.Println("Buku sedang dipinjam")
		}
	}
}

func tampilDataBiasa(data *[NMAX]dataBuku, sumBuku *int) {
	if *sumBuku == 0 {
		fmt.Println("**Tidak ada buku di perpustakaan**")
	}
	fmt.Println()
	fmt.Printf("## Terdapat %d buku dalam perpustakaan ##", *sumBuku)
	for i := 0; i <= *sumBuku-1; i++ {
		fmt.Println()
		fmt.Printf("**Buku ke-%d**\n", i+1)
		fmt.Println("Judul : ", data[i].judul)
		fmt.Println("Nama Pengarang : ", data[i].pengarang)
		fmt.Println("Nomor ISBN : ", data[i].ISBN)
		fmt.Println("Jumlah Eksemplar : ", data[i].eksemplar)
		if data[i].tersedia == true {
			fmt.Println("Buku tersedia")
		} else {
			fmt.Println("Buku sedang dipinjam")
		}
	}
}

func tampilDataPinjam(data *[NMAX]dataBuku, dataPinjam *[NMAX]dataPeminjam, sumPinjam *int) {
	var i int
	if *sumPinjam <= 0 {
		fmt.Println("**Tidak ada pinjaman yang sedang berlangsung**")
	} else {
		fmt.Printf("**Data Peminjam Buku Sebanyak %d orang**\n", *sumPinjam)

		for i = 0; i <= *sumPinjam-1; i++ {
			fmt.Println()
			fmt.Printf("**Peminjam ke-%d**\n", i+1)
			fmt.Println(dataPinjam[i].nama)
			fmt.Println(dataPinjam[i].nomorTelp)
			fmt.Println(dataPinjam[i].bukuDipinjam.judul)
		}
	}
}

func pinjamBuku(data *[NMAX]dataBuku, dataPinjam *[NMAX]dataPeminjam, sumBuku int, sumPinjam *int, inputPinjam string) {
	var found int = -1
	var i int
	for i = 0; i <= sumBuku-1; i++ {
		if inputPinjam == data[i].judul {
			found = i
		}
	}
	if found != -1 && data[found].tersedia == true {
		fmt.Print("Masukkan nama peminjam : ")
		fmt.Scan(&dataPinjam[*sumPinjam].nama)
		fmt.Print("Masukkan nomor telepon peminjam : ")
		fmt.Scan(&dataPinjam[*sumPinjam].nomorTelp)
		fmt.Print("Masukkan lama waktu meminjam (maksimal 7 hari): ")
		fmt.Scan(&dataPinjam[*sumPinjam].lamaPinjaman)
		for dataPinjam[*sumPinjam].lamaPinjaman > 7 {
			fmt.Print("Masukkan lama waktu meminjam (maksimal 7 hari): ")
			fmt.Scan(&dataPinjam[*sumPinjam].lamaPinjaman)
		}
		fmt.Println("**Buku berhasil dipinjam**")
		data[found].tersedia = false
		dataPinjam[*sumPinjam].bukuDipinjam.judul = data[found].judul
		*sumPinjam++
	} else if found != -1 && data[found].tersedia == false {
		fmt.Println("**Buku sedang dipinjam**")
	} else if found == -1 {
		fmt.Println("**Buku tidak ditemukan**")
	}
}

func kembaliBuku(data *[NMAX]dataBuku, dataPinjam *[NMAX]dataPeminjam, sumBuku int, sumPinjam *int, inputKembali string) {
	var i, j, totalWaktu, denda int
	var dataPinjamBaru [NMAX]dataPeminjam
	var found bool = false

	for i = 0; i <= *sumPinjam-1; i++ {
		if inputKembali == dataPinjam[i].bukuDipinjam.judul {
			found = true
		}
		if inputKembali != dataPinjam[i].bukuDipinjam.judul {
			dataPinjamBaru[j] = dataPinjam[i]
			j++
		}
	}
	*dataPinjam = dataPinjamBaru
	*sumPinjam = j

	for i = 0; i <= sumBuku-1; i++ {
		if inputKembali == data[i].judul {
			data[i].tersedia = true
		}
	}
	if found == true {
		fmt.Print("Masukkan total waktu pinjaman : ")
		fmt.Scan(&totalWaktu)
	}
	if found == false {
		fmt.Println("**Buku tidak ditemukan**")
	} else if found == true && totalWaktu > 7 {
		denda = (totalWaktu - 7) * 5000
		fmt.Printf("!! Buku berhasil dikembalikan namun terkena denda senilai Rp. %d,00 !!\n", denda)
	} else if found == true && totalWaktu <= 7 {
		fmt.Println("**Buku berhasil dikembalikan**\n")
	}
}

func main() {

	var pilih, pilihCari, sumBuku, sumPinjam int
	var buku [NMAX]dataBuku
	var cariJudul, cariPengarang string
	var inputHapus, inputEdit string
	var inputPinjam string
	var peminjam [NMAX]dataPeminjam
	var inputKembali string
	var pilihTampilan int

	for pilih != 8 {
		fmt.Println("******************************************")
		fmt.Println("*                                        *")
		fmt.Println("*       Manajemen Buku Perpustakaan      *")
		fmt.Println("*       ============================     *")
		fmt.Println("*       1. Pendataan Buku                *")
		fmt.Println("*       2. Cari Buku                     *")
		fmt.Println("*       3. Hapus Buku                    *")
		fmt.Println("*       4. Edit Data Buku                *")
		fmt.Println("*       5. Tampilkan Data Buku           *")
		fmt.Println("*       6. Peminjaman Buku               *")
		fmt.Println("*       7. Pengembalian Buku             *")
		fmt.Println("*       8. Keluar                        *")
		fmt.Println("*       ============================     *")
		fmt.Println("*                                        *")
		fmt.Println("******************************************")
		fmt.Print("Silahkan Tekan Angka : ")

		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahBuku(&buku, &sumBuku)

		}

		if pilih == 2 {
			if sumBuku == 0 {
				fmt.Println("**Tidak ada buku di perpustakaan**")
			} else {
				fmt.Println()
				fmt.Println("******************************************")
				fmt.Println("*                                        *")
				fmt.Println("*       Metode Pencarian                 *")
				fmt.Println("*      ==============================    *")
				fmt.Println("*      1. Cari dengan judul              *")
				fmt.Println("*      2. Cari dengan nama pengarang     *")
				fmt.Println("*                                        *")
				fmt.Println("******************************************")
				fmt.Print("Silahkan pilih metode : ")
				fmt.Scan(&pilihCari)
				if pilihCari == 1 {
					fmt.Print("Masukkan judul buku : ")
					fmt.Scan(&cariJudul)
					fmt.Println()
					cariJudulBuku(&buku, sumBuku, cariJudul)

				} else if pilihCari == 2 {
					fmt.Print("Masukkan nama pengarang buku : ")
					fmt.Scan(&cariPengarang)
					fmt.Println()
					cariPengarangBuku(&buku, sumBuku, cariPengarang)
				}
			}

		}

		if pilih == 3 {
			if sumBuku == 0 {
				fmt.Println("**Tidak ada buku di perpustakaan**")
			} else {
				fmt.Print("Masukkan judul buku yang ingin dihapus : ")
				fmt.Scan(&inputHapus)
				fmt.Println()
				hapusData(&buku, &sumBuku, inputHapus)
			}
		}

		if pilih == 4 {
			if sumBuku == 0 {
				fmt.Println("**Tidak ada buku di perpustakaan**")
			} else {
				fmt.Print("Masukkan judul buku yang ingin diedit datanya : ")
				fmt.Scan(&inputEdit)
				fmt.Println()
				editData(&buku, &sumBuku, inputEdit)
			}
		}

		if pilih == 5 {
			fmt.Println()
			fmt.Println("********************************************")
			fmt.Println("*                                          *")
			fmt.Println("*       Format Tampilan                    *")
			fmt.Println("*      ==============================      *")
			fmt.Println("*      1. Tampilan Biasa                   *")
			fmt.Println("*      2. Terurut berdasarkan judul        *")
			fmt.Println("*      3. Terurut berdasarkan pengarang    *")
			fmt.Println("*      4. Data Peminjam                    *")
			fmt.Println("*                                          *")
			fmt.Println("********************************************")
			fmt.Print("Silahkan pilih format tampilan : ")
			fmt.Scan(&pilihTampilan)

			if pilihTampilan == 1 {
				tampilDataBiasa(&buku, &sumBuku)
			} else if pilihTampilan == 2 {
				tampilDataJudul(&buku, &sumBuku)
			} else if pilihTampilan == 3 {
				tampilDataPengarang(&buku, &sumBuku)
			} else if pilihTampilan == 4 {
				tampilDataPinjam(&buku, &peminjam, &sumPinjam)
			}

		}

		if pilih == 6 {
			fmt.Print("Masukkan judul buku yang ingin dipinjam : ")
			fmt.Scan(&inputPinjam)
			pinjamBuku(&buku, &peminjam, sumBuku, &sumPinjam, inputPinjam)
		}

		if pilih == 7 {
			fmt.Print("Masukkan judul buku yang ingin dikembalikan : ")
			fmt.Scan(&inputKembali)
			kembaliBuku(&buku, &peminjam, sumBuku, &sumPinjam, inputKembali)
		}
	}
}
