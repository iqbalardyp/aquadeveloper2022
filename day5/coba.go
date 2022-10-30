package coba

import "fmt"

type Coba []struct {
	nomor int
}

func (c Coba) CetakAngka() {
	fmt.Println(c)
}
