package main

import (
	"fmt"
	"github.com/madasatya6/go-pkg/barcode"
)

func main() {
	unikstr, err := barcode.GenerateImage(nil)

	if err != nil {
		fmt.Printf("Error: %w", err)
		return 
	}

	fmt.Println("nomor unik barcode : ", unikstr)
}


