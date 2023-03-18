package main

import (
	"fmt"
	"strconv"
)

func main() {

	var i = 21
	var persen = "%"
	var boolBenar bool = true
	var boolSalah bool = false
	var angka = 21
	var f = 15
	var str = "Я"
	var float = 123.456

	fmt.Println(i)
	fmt.Printf("%T \n", i)
	fmt.Println(persen)
	fmt.Println(boolBenar)
	fmt.Println(boolSalah)
	fmt.Printf("%d\n", angka)
	fmt.Printf("%o\n", angka)
	fmt.Printf("%x\n", f)
	fmt.Printf("%X\n", f)
	asciiStr := strconv.QuoteToASCII(str)
	fmt.Println(asciiStr)
	fmt.Printf("%f\n", float)
	fmt.Printf("%E\n", float)

}
