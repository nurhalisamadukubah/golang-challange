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

const job = "IT dev"

type biodataGenerator interface {
	biodata() []student
}

type namaBiodata struct {
	nama []string
}

func newNamaBiodata(nama []string) biodataGenerator {
	return &namaBiodata{nama: nama}
}

func (n *namaBiodata) biodata() []student {
	address := []string{"kessilampe", "mt2", "gunjat"}
	reason := []string{"punya nur", "punya lily", "punya acel"}
	generate := make([]student, 0)
	var s student
	for key, value := range n.nama {
		s.nama = value
		s.alamat = address[key]
		s.pekerjaan = job
		s.alasan = reason[key]
		generate = append(generate, s)
	}
	return generate
}

func main() {
	nama := []string{"nur", "lily", "acel"}
	generator := newNamaBiodata(nama)
	output := generator.biodata()
	var arg string = os.Args[1]
	argInt, _ := strconv.Atoi(arg)
	var inama int = argInt - 1

	for i, v := range output {
		if inama == i {
			fmt.Println("ID : ", i)
			fmt.Println("Nama : ", v.nama)
			fmt.Println("Alamat : ", v.alamat)
			fmt.Println("Pekerjaan : ", v.pekerjaan)
			fmt.Println("alasan : ", v.alasan)

		}

	}

	//fmt.Println(output)
}
