package main

import "fmt"

func main() {
	input := "selamat malam"

	//menghitung kemunculan karakter
	charCount := make(map[string]int)

	//iterasi karakter 'input' dan menghitung frekuensi kemunculannya
	for _, s := range input {
		fmt.Printf("%c\n", s)
		charCount[string(s)] += 1
	}

	fmt.Println(charCount)
}
