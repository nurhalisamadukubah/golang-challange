package main

import (
	"fmt"
	"os"
	"strconv"
)

type student struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

// deklarasi interface yang mengambil nilai dari slice student
type biodataGenerator interface {
	biodata() []student
}

// deklarasi struct yang mengambil data dari slice nama
type namaBiodata struct {
	nama []string
}

// fungsi untuk mengambil slice nama yang dikembalikan ke namaBiodata dalam bentuk pointer
func newNamaBiodata(nama []string) biodataGenerator {
	return &namaBiodata{nama: nama}
}

// fungsi untuk menerima pointer
func (n *namaBiodata) biodata() []student {
	address := []string{"Kel. Kessilampe", "Lr. Mata Air 2", "Gunung Jati"}
	job := []string{"Back End Dev", "Android Dev", "HR Consultant"}
	reason := []string{"Belajar", "Dapat Sertifikat", "Switch Career"}
	generate := make([]student, 0) //membuat slice untuk menyimpan var s yang sudah memiliki value
	var s student
	for key, value := range n.nama {
		s.nama = value
		s.alamat = address[key]
		s.pekerjaan = job[key]
		s.alasan = reason[key]
		generate = append(generate, s)
	}
	return generate
}

func main() {
	nama := []string{"Nur", "Lily", "Achel"}
	generator := newNamaBiodata(nama) //mengambil pointer
	output := generator.biodata()     //membuat slice student berdasarkan pointer dan memberikan atribut sesuai dengan method biodata
	var arg string = os.Args[1]
	argInt, err := strconv.Atoi(arg) //konversi string menjadi int
	inama := argInt - 1              //kan di argInt slicenya jadi 1,2,3... sementara indeks nama dari 0,

	//kondisi agar argumen yang diterima int
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	// kondisi agar argumen sesuai dengan ID
	if inama < 0 || inama >= len(nama) {
		fmt.Println("Invalid input")
		return
	}

	for i, v := range output {
		if inama == i {
			fmt.Println("ID : ", i+1) //i+1 agar ID dimulai dari 1 karena aslinya i mengikuti indeks slice nama
			fmt.Println("Nama : ", v.nama)
			fmt.Println("Alamat : ", v.alamat)
			fmt.Println("Pekerjaan : ", v.pekerjaan)
			fmt.Println("Alasan : ", v.alasan)

		}

	}

}
