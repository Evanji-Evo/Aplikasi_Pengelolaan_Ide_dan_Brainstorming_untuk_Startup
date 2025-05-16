package main

import "fmt"
import "os"

const NMAX int = 9999

type Ide struct {
	nama_ide string
	rating int
	jml_rating int 
}

type kumpulanIde [NMAX]Ide

func main(){
	var dataIde kumpulanIde
	var input, judul int
	judul = 0
	
	for {
		menuUtama(&input,&judul,&dataIde)
		fmt.Println()
	}
}

func menuUtama(input *int, judul *int, dataIde *kumpulanIde){
	fmt.Println(" *** ------------------------------------------- ***")
	fmt.Println("           APLIKASI PENGELOLA IDE STARTUP            ")
	fmt.Println("         EVO ABIMAYU & PUTRI RAHAYU DAMAYANTI         ")
	fmt.Println("             ALGORITMA DAN PEMROGRAMAN               ")
	fmt.Println(" *** ------------------------------------------- ***")
	fmt.Printf("1. Tampilkan isi data Ide\n2. Tambahkan data Ide\n3. Ubah data Ide\n4. Hapus data Ide\n5. Berikan Rating Ide\n6. Cari data Ide Spesifik\n7. Hasil Ranking Voting Ide\n8. Keluar Aplikasi\n")
	fmt.Print("Pilihan : ")
	fmt.Scan(&(*input))
	switch *input {
		case 1 : 
			CetakdataIde(judul,dataIde)
		case 2 :
			PenambahanIde(judul, dataIde)
		case 3 :
			PengubahanIde(judul, dataIde)
		case 4 :
			PenghapusanIde(judul, dataIde)
		case 5 :
			ratingIde(judul ,dataIde)
		case 6 :
			cariIde(judul, dataIde)
		case 7 :
			HasilIde(judul, dataIde)
		case 8 :
			Keluar()
	}
}

func CetakdataIde(judul *int, dataIde *kumpulanIde){
	var i int
	for i = 0 ; i < *judul ; i++{
		fmt.Printf("%d. %s \nRating : %d\n",i+1,dataIde[i].nama_ide, dataIde[i].rating)
	}
}

func PenambahanIde(judul *int, dataIde *kumpulanIde){
	var i , jumlah_ideSekarang,jumlah_ideBaru int
	fmt.Print("Jumlah ide yang ingin di tambahkan : ")
	fmt.Scan(&jumlah_ideBaru)
	jumlah_ideBaru = jumlah_ideBaru + *judul
	jumlah_ideSekarang = *judul
	fmt.Println("Gunakan uderscore atau _ untuk spasi")
	for i = jumlah_ideSekarang; i < jumlah_ideBaru ; i++{
		fmt.Print("Masukan Ide : ")
		fmt.Scan(&dataIde[i].nama_ide)
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

func ratingIde(judul * int, dataIde *kumpulanIde){
	var bintang int
	var i int

	if *judul == 0 {
		fmt.Println("")
	} else {
		fmt.Println("Berikan rating:")
		for i = 0; i < *judul; i++ {
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
		for i = 0; i < *judul; i++ {
			fmt.Printf("%d. %s  Total Rating: %d dari %d orang\n", i+1, dataIde[i].nama_ide, dataIde[i].rating, dataIde[i].jml_rating)
		}
		fmt.Println()
	}
}

func cariIde(n *int, dataIde *kumpulanIde){
	var data string
	fmt.Print("Nama ide yang sedang dicari : ")
	fmt.Scan(&data)
			
	sequentialSearch(n,dataIde,data)
}

func sequentialSearch(n *int, dataIde *kumpulanIde, data string){
	var i int
	for i = 0 ; i < *n ; i++{
		if (dataIde[i].nama_ide == data ){
			fmt.Printf("Nama Ide : %s\n", dataIde[i].nama_ide)
			fmt.Printf("Dengan Skor Rating : %d\n",dataIde[i].rating)
			fmt.Printf("Dan dengan jumlah pemilih : %d",dataIde[i].jml_rating)
		}
	}
	fmt.Println()
}

func HasilIde(judul *int, A *kumpulanIde){
	var input , i int
	fmt.Printf("Pilih cara sorting : \n1. Selection Sort\n2. Insertion Sort\nPilihan : ")
	fmt.Scan(&input)
	switch input {
		case 1 : 
			SelectionSort(judul,A)
		case 2 :
			insertionSort(judul,A)
	}
	fmt.Println("Ide diurutkan secara descending dari rating tertinggi : ")
	for i = 0 ; i < *judul ; i++{
		fmt.Printf("Nama Ide : %s\n", A[i].nama_ide)
		fmt.Printf("Rating : %d\n\n", A[i].rating)
	}
}

func SelectionSort(judul *int, A *kumpulanIde){
	var i, idx, pass int
	var temp Ide
	
	pass = 1
	for pass < *judul {
		idx = pass - 1
		i = pass
		for i < *judul {
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

func insertionSort(judul *int, A *kumpulanIde){
	var i, j int
	var temp Ide
	for i = 1 ; i < *judul ; i++ {
		j = i 
		temp = A[i]
		for j > 0 && temp.rating > A[i-1].rating {
			A[j] = A[j-1]
			j = j -1
		}
		A[j] = temp
	}
}

func Keluar(){
	os.Exit(0)
}