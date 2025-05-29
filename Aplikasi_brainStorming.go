package main

import "fmt"

const NMAX int = 999

type Ide struct {
	nomor int
	nama_ide string
	rating int
	jml_rating int
	hari,bulan,tahun int
}

type kumpulanIde [NMAX]Ide

func main(){
	var dataIde kumpulanIde
	var input, judul int
	judul = 0
	
	for input != 8 {
		menuUtama(&input)
		switch input {
		case 1 : 
			CetakdataIde(judul,dataIde)
		case 2 :
			PenambahanIde(&judul, &dataIde)
		case 3 :
			PengubahanIde(&judul, &dataIde)
		case 4 :
			PenghapusanIde(&judul, &dataIde)
		case 5 :
			ratingIde(judul ,&dataIde)
		case 6 :
			cariIde(judul, dataIde)
		case 7 :
			urutanIde(judul, dataIde)
	}
		fmt.Println()
	}
}

func menuUtama(input *int){
	fmt.Println(" *** ------------------------------------------- ***")
	fmt.Println("           APLIKASI PENGELOLA IDE STARTUP            ")
	fmt.Println("         EVO ABIMAYU & PUTRI RAHAYU DAMAYANTI         ")
	fmt.Println("             ALGORITMA DAN PEMROGRAMAN               ")
	fmt.Println(" *** ------------------------------------------- ***")
	fmt.Printf("1. Tampilkan isi data Ide\n2. Tambahkan data Ide\n3. Ubah data Ide\n4. Hapus data Ide\n5. Berikan Rating Ide\n6. Cari data Ide\n7. Pengurutan Ranking Voting Ide\n8. Keluar Aplikasi\n")
	fmt.Print("Pilihan : ")
	fmt.Scan(input)
}

func CetakdataIde(judul int, dataIde kumpulanIde){
	var i int
	fmt.Print("-----------------------------------------------------------")
	fmt.Printf("\n| %s | %-20s | %-10s | %-14s |\n","No" , "Judul Ide", "Rating", "Tanggal")
	fmt.Println("-----------------------------------------------------------")
	for i = 0 ; i < judul ; i++{
		fmt.Printf("| %-2d | %-20s | %-10d | %-1d / %-1d / %-5d |\n",i+1 , dataIde[i].nama_ide, dataIde[i].rating, dataIde[i].hari, dataIde[i].bulan, dataIde[i].tahun)
	}
	fmt.Println("-----------------------------------------------------------")
}

func PenambahanIde(judul *int, dataIde *kumpulanIde){
	var i , jumlah_ideSekarang,jumlah_ideBaru int
	var hari,bulan,tahun int
	
	fmt.Print("Jumlah ide yang ingin di tambahkan : ")
	fmt.Scan(&jumlah_ideBaru)
	for jumlah_ideBaru < 1 {
		fmt.Printf("Tidak bisa menambakan %d ide baru, berikan nilai bilangan asli (1,2,3,..): ",jumlah_ideBaru)
		fmt.Scan(&jumlah_ideBaru)
	}
	
	fmt.Print("Hari (1-31) : ")
	fmt.Scan(&hari)
	for hari < 1 || hari> 31 {
		fmt.Print("Hari tidak valid, masukan hari yang valid (1-32): ")
		fmt.Scan(&hari)
	}
	
	fmt.Print("Bulan (1-12) : ")
	fmt.Scan(&bulan)
	for bulan < 1 || bulan > 12 {
		fmt.Print("Bulan tidak valid, masukan bulan yang valid (1-12 ): ")
		fmt.Scan(&bulan)
	}

	fmt.Print("Tahun : ")
	fmt.Scan(&tahun)
	
	jumlah_ideBaru = jumlah_ideBaru + *judul
	jumlah_ideSekarang = *judul
	
	fmt.Println("Gunakan uderscore atau _ untuk spasi")
	for i = jumlah_ideSekarang; i < jumlah_ideBaru ; i++{
		fmt.Printf("%d. Nama Ide : ",i+1)
		fmt.Scan(&dataIde[i].nama_ide)
		dataIde[i].hari = hari
		dataIde[i].bulan = bulan
		dataIde[i].tahun = tahun
		dataIde[i].nomor = i + 1
	}
	*judul = i
	
	fmt.Println()
}

func PengubahanIde(judul *int, dataIde *kumpulanIde){
	var i, Nomor_Ide int
	for i = 0 ; i < *judul ; i++{
		fmt.Printf("%d. %s\n",i+1,dataIde[i].nama_ide)
	}
	fmt.Print("Nomor Data yang ingin diubah : ")
	fmt.Scan(&Nomor_Ide)
	fmt.Print("Di ubah menjadi : ")
	fmt.Scan(&dataIde[Nomor_Ide-1].nama_ide)
	fmt.Println()
}

func PenghapusanIde(judul *int, dataIde *kumpulanIde){
	var i, Nomor_Ide int
	for i = 0 ; i < *judul ; i++{
		fmt.Printf("%d. %s\n",i+1,dataIde[i].nama_ide)
	}
	fmt.Print("Nomor Data yang ingin dihapus : ")
	fmt.Scan(&Nomor_Ide)
	for i = Nomor_Ide-1 ; i < *judul ; i++{
		dataIde[i] = dataIde[i+1]
	}
	*judul=*judul-1
}

func ratingIde(judul int, dataIde *kumpulanIde){
	var bintang int
	var i int

	if judul == 0 {
		fmt.Println("")
	} else {
		fmt.Println("Berikan rating:")
		for i = 0; i < judul; i++ {
			fmt.Printf("%d. %s", i+1, dataIde[i].nama_ide)
			fmt.Print(" Rating: ")
			fmt.Scan(&bintang)

			if bintang >= 1 && bintang <= 5 {
				dataIde[i].rating += bintang
			}else if (bintang > 5){
				dataIde[i].rating +=  5
				fmt.Println("tidak valid, rating harus kurang dari 5. Karena lebih maka rating 5")
			}else{
				fmt.Println("tidak valid, rating harus lebih dari 1. Karena kurang maka rating 0")
			}
			dataIde[i].jml_rating++
		}

		fmt.Println("\nTotal Rating:")
		for i = 0; i < judul; i++ {
			fmt.Printf("%d. %s  Total Rating: %d dari %d orang\n", i+1, dataIde[i].nama_ide, dataIde[i].rating, dataIde[i].jml_rating)
		}
		fmt.Println()
	}
}

func cariIde(n int, dataIde kumpulanIde){
	var pilihan , i int
	fmt.Printf("Cari berdasarkan\n1. Nomor\n2. Judul ide\nPilihan : ")
	fmt.Scan(&pilihan)
	
	for i = 0 ; i < n ; i++{
		fmt.Printf("%d. %s\n",i+1,dataIde[i].nama_ide)
	}
	switch pilihan {
		case 1 :
			binarySearchNomor(n,dataIde)
		case 2 :
			sequentialSearchJudul(n,dataIde)	
	}
}

func binarySearchNomor(n int, dataIde kumpulanIde){
	var kanan, kiri, tengah int
	var nomor int
	
	kiri = 0
	kanan = n
	fmt.Print("Nomor ide yang di cari : ")
	fmt.Scan(&nomor)
	
	for kiri <= kanan{
		tengah = (kiri+kanan)/2
		if (dataIde[tengah].nomor < nomor){
			kiri = tengah + 1
		}else if (dataIde[tengah].nomor > nomor){
			kanan = tengah - 1
		}else{
			fmt.Printf("Judul Ide : %s\n", dataIde[tengah].nama_ide)
			fmt.Printf("Tanggal : %d/%d/%d\n", dataIde[tengah].hari, dataIde[tengah].bulan, dataIde[tengah].tahun)
			fmt.Printf("Dengan Skor Rating : %d\n",dataIde[tengah].rating)
			fmt.Printf("Dan dengan jumlah pemilih : %d\n",dataIde[tengah].jml_rating)
			kiri = kanan + 1
		}
	}
	
	
	
}

func sequentialSearchJudul(n int, dataIde kumpulanIde){
	var i int
	var data string
	
	fmt.Print("Judul ide yang di cari : ")
	fmt.Scan(&data)
	for i = 0 ; i < n ; i++{
		if (dataIde[i].nama_ide == data ){
			fmt.Printf("Judul Ide : %s\n", dataIde[i].nama_ide)
			fmt.Printf("Tanggal : %d/%d/%d\n", dataIde[i].hari, dataIde[i].bulan, dataIde[i].tahun)
			fmt.Printf("Dengan Skor Rating : %d\n",dataIde[i].rating)
			fmt.Printf("Dan dengan jumlah pemilih : %d",dataIde[i].jml_rating)
		}
	}
	fmt.Println()
}

func urutanIde(judul int, A kumpulanIde){
	var berdasarkan, caraPengurutan int 
	fmt.Printf("Pilih urutkan berdasarkan : \n1. Rating \n2. Tanggal\nPilihan : ")
	fmt.Scan(&berdasarkan)
	
	if berdasarkan == 1 {
		//Menu Pengurutan Berdasarkan rating
		fmt.Printf("Pilih cara pengurutan ranking : \n1. Descending atau Paling Populer \n2. Ascending ataa Paling tidak populer)\nPilihan : ")
		fmt.Scan(&caraPengurutan)
		switch caraPengurutan{
			case 1 :
				SelectionSortRating(judul,&A)
			case 2 : 
				insertionSortRating(judul,&A)
			}
		
	}else{
		//Menu Pengurutan Berdasarkan Tanggal 
		fmt.Printf("Pilih cara pengurutan ranking : \n1. Descending atau Data paling baru \n2. Ascending atau Data paling lama\nPilihan : ")
		fmt.Scan(&caraPengurutan)
		switch caraPengurutan{
			case 1 :
				SelectionSortTanggal(judul,&A)
			case 2 : 
				insertionSortTanggal(judul,&A)
			}
		
	}
	
	CetakdataIde(judul,A)
}

func SelectionSortRating(judul int, A *kumpulanIde){
	var i, idx, pass int
	var temp Ide
	
	pass = 1
	for pass < judul {
		idx = pass - 1
		i = pass
		for i < judul {
			if (A[i].rating > A[idx].rating) {
				idx = i
			}else{
				i = i + 1
			}
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}

func insertionSortRating(judul int, A *kumpulanIde){
	var i, j int
	var temp Ide
	for i = 1 ; i < judul ; i++ {
		j = i 
		temp = A[i]
		for j > 0 && temp.rating < A[i-1].rating {
			A[j] = A[j-1]
			j = j -1
		}
		A[j] = temp
	}
}

func SelectionSortTanggal(judul int, A *kumpulanIde){
	var i, idx, pass int
	var temp Ide
	
	pass = 1
	for pass < judul {
		idx = pass - 1
		i = pass
		for i < judul {
			if (A[i].tahun > A[idx].tahun || A[i].tahun == A[idx].tahun && A[i].bulan < A[idx].bulan || A[i].tahun == A[idx].tahun && A[i].bulan == A[idx].bulan && A[i].hari < A[idx].hari) {
				idx = i
			}else{
				i = i + 1
			}
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}

func insertionSortTanggal(judul int, A *kumpulanIde){
	var i, j int
	var temp Ide
	for i = 1 ; i < judul ; i++ {
		j = i 
		temp = A[i]
		for j > 0 && temp.tahun < A[i-1].tahun || j == 0 && temp.tahun == A[i-1].tahun && temp.bulan > A[i-1].bulan || j == 0 && temp.tahun == A[i-1].tahun && temp.bulan == A[i-1].bulan && temp.hari > A[i-1].hari {
			A[j] = A[j-1]
			j = j -1
		}
		A[j] = temp
	}
}