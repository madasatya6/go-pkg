package main

import (
	"fmt"
	"github.com/madasatya6/go-pkg/barcode"
)

func main() {
	unikstr, err := barcode.GenerateImage(barcode.Config{
		Directory: "assets/barcode/",
		Extension: ".jpg",
		Key: "golang",
		Width: 200,
		Height: 200,
	})

	if err != nil {
		fmt.Printf("Error: %w", err)
		return 
	}

	fmt.Println("nomor unik barcode : ", unikstr)
}


