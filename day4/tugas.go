package main

import (
	"fmt"
	"strings"
)

func cetakBintang(n int) {
	if n%2 == 1 {
		for i := 1; i <= n; i++ {
			fmt.Println(strings.Repeat("*", i))
		}
		return
	} else {
		for i := 1; i <= n; i++ {
			fmt.Println(strings.Repeat("*", n))
		}
		return
	}
}

func main() {
	var n int

	fmt.Println("Masukkan angka:")
	fmt.Scanln(&n)

	defer cetakBintang(n)
}
