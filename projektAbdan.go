package main

import "fmt"

type tempatWisata struct{
	Nama string
	Kategori string
	Jarak float64
	Biaya int
	Fasilitas string
}

const NMAX = 5
type tabWisata[NMAX]tempatWisata

var tw tabWisata
var nw int = 0

//menu utama
func main(){
	menu_pilihan()
	
}

func menu_pilihan(){
	var pilih int
	for {
		fmt.Println("== APLIKASI PARIWISATA == ")
		fmt.Println("1. Tambah Tempat Wisata (Admin)")
		fmt.Println("2. Edit Tempat Wisata (Admin)")
		fmt.Println("3. Hapus Tempat Wisata (Admin)")
		fmt.Println("4. Lihat Semua Tempat Wisata")
		fmt.Println("5. Cari Tempat Wisata")
		fmt.Println("6. Filter berdasarkan Kategori")
		fmt.Println("7. Urutkan Tempat Wisata")
		fmt.Println("8. Keluar")
		fmt.Println("Pilih menu: ")
		
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			tambahTempat(&tw, &nw)
		case 2:
			if nw == 0 {
				fmt.Println("DATA KOSONG")
			} else {
				fmt.Println("\n= DAFTAT TEMPAT WISATA =")
				tampilkanData(tw, nw)
				var x string
				fmt.Print("Masukkan nama tempat yang ingin diedit: ")
				fmt.Scan(&x)
				editTempat(&tw, nw, x)
			}
		case 3:
			if nw == 0 {
				fmt.Println("DATA KOSONG")
			} else {
				fmt.Println("\n= DAFTAT TEMPAT WISATA =")
				tampilkanData(tw, nw)
				var x string
				fmt.Print("Masukkan nama tempat yang ingin dihapus: ", x)
				fmt.Scan(&x)
				hapusTempat(&tw, &nw, x)
			}
		case 4:
			if nw == 0 {
				fmt.Println("DATA KOSONG")
			} else {
				fmt.Println("\n= DAFTAT TEMPAT WISATA =")
				tampilkanData(tw, nw)
			}
		case 5:
			if nw == 0 {
				fmt.Println("DATA KOSONG")
			} else {
				var x string
				fmt.Print("Masukkan nama tempat yang ingin dicari: ")
				fmt.Scan(&x)
				cariTempat(tw, nw, x)
			}
		case 6:
			if nw == 0 {
				fmt.Println("DATA KOSONG")
			} else {
				var x string
				fmt.Print("Masukkan kategori yang ingin dicari: ", x)
				fmt.Scan(&x)
				filterKategori(tw, nw, x)
			}
		case 7:
			urutkanTempat(&tw, nw)
		case 8:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
		default:
			fmt.Println("Pilihan tidak valid!")
		}
		
		if pilih == 8 {
			break
		}
	}	
}

func tambahTempat(tw *tabWisata, n *int){
	if *n < NMAX {
		fmt.Print("Tambahkan tempat wisata: ")
		fmt.Scan(&tw[*n].Nama)
		fmt.Print("Kategori (Pantai/Gunung/Taman dll): ")
		fmt.Scan(&tw[*n].Kategori)
		fmt.Print("Jarak dari lokasi anda (km): ")
		fmt.Scan(&tw[*n].Jarak)
		fmt.Print("Biaya masuk (Rp): ")
		fmt.Scan(&tw[*n].Biaya)
		fmt.Print("Fasilitas yang tersedia: ")
		fmt.Scan(&tw[*n].Fasilitas)
		*n++
		fmt.Println("Data berhasil ditambahkan!")
	}else {
		fmt.Println("Data penuh")
	}
	
}

func editTempat(tw *tabWisata, n int, x string){
	var i int
	i = seqSearch(*tw, n, x)
	if i != -1 {
		fmt.Print("Nama Baru: ")
		fmt.Scan(&tw[i].Nama)
		fmt.Print("Kategori: ")
		fmt.Scan(&tw[i].Kategori)
		fmt.Print("Jarak Baru(km): ")
		fmt.Scan(&tw[i].Jarak)
		fmt.Print("Biaya Baru(Rp): ")
		fmt.Scan(&tw[i].Biaya)
		fmt.Print("Fasilitas yang tersedia: ")
		fmt.Scan(&tw[i].Fasilitas)
		fmt.Println("Data berhasil diedit!")
	} else {
		fmt.Println("Data tidak ditemukan!")
	}	
}

func seqSearch(tw tabWisata, n int, x string)int{
	var i, idx int
	
	idx = -1
	i = 0
	for i < n && idx == -1 {
		if tw[i].Nama == x {
			idx = i
		}
		i++
	}
	return idx
}

func binSearch(tw tabWisata, n int, x string) int {
	var left, right, mid int

	left = 0
	right = n - 1
	for left <= right {
		mid = (left + right) / 2
		if tw[mid].Nama == x {
			return mid
		} else if x < tw[mid].Nama {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}


func hapusTempat(tw *tabWisata, n *int, x string){
	var i, k int
	k = binSearch(*tw, *n, x)
	if k != -1 {
		for i = k+1; i < *n; i++{
			tw[i-1] = tw[i]
		}
		*n--
		fmt.Println("Data berhasil dihapus!" )
	} else {
		fmt.Println("Data tidak ditemukan!")
	}	
}

func tampilkanData(tw tabWisata, n int ){
	var i int
	
	fmt.Printf("%-4s %-20s %-12s %-10s %-10s %-20s\n", "No", "Nama", "Kategori", "Jarak", "Biaya", "Fasilitas")
	
	for i = 0; i < n; i++ {
		fmt.Printf("%-4d %-20s %-12s %-10.2f %-10d %-20s\n",
			i+1,
			tw[i].Nama,
			tw[i].Kategori,
			tw[i].Jarak,
			tw[i].Biaya,
			tw[i].Fasilitas)
	}
}

func cariTempat(tw tabWisata, n int, x string){
	var i int
	i = seqSearch(tw, n, x)
	if i != -1 {
        fmt.Println("\n== Data Ditemukan ==")
        fmt.Println("Nama      :", tw[i].Nama)
        fmt.Println("Kategori  :", tw[i].Kategori)
        fmt.Printf("Jarak     : %.2f km\n", tw[i].Jarak)
        fmt.Printf("Biaya     : Rp %d\n", tw[i].Biaya)
        fmt.Println("Fasilitas :", tw[i].Fasilitas)
    } else {
        fmt.Println("Tempat wisata tidak ditemukan.")
    }
}


func filterKategori(tw tabWisata, n int, x string) {
	var i, nomor int
	var ketemu bool
	ketemu = false
	nomor = 1

	// Header tabel
	fmt.Printf("%-4s %-20s %-10s %-10s %-20s\n", "No", "Nama", "Jarak", "Biaya", "Fasilitas")

	for i = 0; i < n; i++ {
		if tw[i].Kategori == x {
			fmt.Printf("%-4d %-20s %-10.2f %-10d %-20s\n",
				nomor,
				tw[i].Nama,
				tw[i].Jarak,
				tw[i].Biaya,
				tw[i].Fasilitas)
			nomor++
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("Tidak ada tempat wisata dengan kategori tersebut")
	}
}

func urutkanTempat(tw *tabWisata, n int) {
    var pilihan, urutan int
    
    fmt.Println("1. Urutkan berdasarkan Biaya")
    fmt.Println("2. Urutkan berdasarkan Jarak")
    fmt.Print("Pilih: ")
    fmt.Scan(&pilihan)
    
    if pilihan == 1 || pilihan == 2 {
        fmt.Println("\n1. Ascending (Kecil ke Besar)")
        fmt.Println("2. Descending (Besar ke Kecil)")
        fmt.Print("Pilih urutan: ")
        fmt.Scan(&urutan)
    }
    
    switch pilihan {
    case 1:
        if urutan == 1 {
            selectionSortBiayaAsc(tw, n)
            fmt.Println("\nData telah diurutkan berdasarkan Biaya (Ascending)")
        } else if urutan == 2 {
            insertionSortBiayaDesc(tw, n)
            fmt.Println("\nData telah diurutkan berdasarkan Biaya (Descending)")
        } else {
            fmt.Println("Pilihan urutan tidak valid!")
		}
        
    case 2:
        if urutan == 1 {
            selectionSortJarakAsc(tw, n)
            fmt.Println("Data telah diurutkan berdasarkan Jarak (Ascending)")
        } else if urutan == 2 {
            insertionSortJarakDesc(tw, n)
            fmt.Println("Data telah diurutkan berdasarkan Jarak (Descending)")
        } else {
            fmt.Println("Pilihan urutan tidak valid!")
            
        }
    default:
        fmt.Println("Pilihan tidak valid!")
        
    }
    
    tampilkanData(*tw, n)
}

func selectionSortBiayaAsc(tw *tabWisata, n int) {
    var pass, i, idx int
    var temp tempatWisata
    
    pass = 1
    for pass < n {
        idx = pass - 1
        i = pass
        for i < n {
            if tw[i].Biaya < tw[idx].Biaya {
                idx = i
            }
            i++
        }
        temp = tw[pass-1]
        tw[pass-1] = tw[idx]
        tw[idx] = temp
        pass++
    }
}

func insertionSortBiayaDesc(tw *tabWisata, n int) {
    var i, pass int
    var temp tempatWisata

    pass = 1
    for pass <= n - 1 {
        temp = tw[pass]
        i = pass
        for i > 0 && temp.Biaya > tw[i-1].Biaya {
            tw[i] = tw[i-1]
            i--
        }
        tw[i] = temp
        pass++
    }
}

func selectionSortJarakAsc(tw *tabWisata, n int) {
    var pass, i, idx int
    var temp tempatWisata
    
    pass = 1
    for pass < n {
        idx = pass - 1
        i = pass
        for i < n {
            if tw[i].Jarak < tw[idx].Jarak {
                idx = i
            }
            i++
        }
        temp = tw[pass-1]
        tw[pass-1] = tw[idx]
        tw[idx] = temp
        pass++
    }
}

func insertionSortJarakDesc(tw *tabWisata, n int) {
    var i, pass int
    var temp tempatWisata

    pass = 1
    for pass <= n - 1 {
        temp = tw[pass]
        i = pass
        for i > 0 && temp.Jarak > tw[i-1].Jarak {
            tw[i] = tw[i-1]
            i--
        }
        tw[i] = temp
        pass++
    }
}
