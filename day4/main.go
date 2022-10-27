package main

import "fmt"

func main() {
	// Print Hello World
	fmt.Println("Hello World!")

	// Delare Variables
	angka1 := 1
	dec1 := 0.1
	kata1 := "lala"
	bool1 := true

	var (
		// Integer
		angka2 int = 2
		angka3     = 3

		// Float
		dec2 float32 = 0.2
		dec3         = 0.5

		// String
		kata2 string = "po"
		kata3        = "olala"

		// Bool
		bool2 bool = false
	)

	fmt.Printf("Integer: %d %d %d", angka1, angka2, angka3)
	fmt.Printf("\nFloat: %f %.2f %.3f", dec1, dec2, dec3)
	fmt.Printf("\nString: %s %s %s", kata1, kata2, kata3)
	fmt.Printf("\nBool: %t %t", bool1, bool2)

	// Declare multi-variable
	var nama1, nama2, nama3 string = "iqbal", "ardy", "putra"
	fmt.Printf("\n%s %s %s", nama1, nama2, nama3)

	// Constant
	const PI float32 = 3.14
	fmt.Println()
	fmt.Println(PI)

	// Array
	var arr = []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	fmt.Println(len(arr))

	// If..Else..Then
	const r float32 = 7.0
	var area float32
	area = PI * r
	if bool1 {
		fmt.Println(area)
	} else {
		fmt.Print(bool2)
	}

	// Switch..Case
	switch {
	case (angka2 == 2):
		fmt.Println(nama2)
	case (angka2 == 3):
		fmt.Println(nama3)
	}

	// Looping
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Array indeks ke-%d: %d\n", i, arr[i])
	}

	// Underscore
	_, test1 := 123, 900
	fmt.Println(test1)

	var chicken map[string]int
	chicken = map[string]int{
		"januari":  50,
		"februari": 100,
	}
	fmt.Println(chicken["januari"])
}
