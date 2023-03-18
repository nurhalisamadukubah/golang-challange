package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i)
	}

	for j := 0; j < 11; j++ {
		if j == 5 {
			// var char = []rune{}
			// char = append(char, 'С', 'А', 'Ш', 'А', 'Р', 'В', 'О')
			// position := 0
			// for _, value := range char {
			// 	fmt.Printf("character %U '%c' starts at byte position %d\n", value, value, position)
			// 	position += 2
			// }
			for position, value := range "САШАРВО" {
				fmt.Printf("character %U '%c' starts at byte position %d\n", value, value, position)
			}
		} else {
			fmt.Println("Nilai j = ", j)
		}
	}
}
