package main

import (
	"errors"
	"fmt"
)

func main() {
	// Masukkan inputan tinggi segitiga
	input, err := getUserInput("Masukkan tinggi segitiga: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 1; i <= input; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println() // Pindah ke baris berikutnya setelah mencetak bintang
	}

	fmt.Print("Tekan Enter untuk keluar...")
	fmt.Scanln()
}

func getUserInput(input string) (int, error) {
	var userInput int
	fmt.Print(input)
	fmt.Scanln(&userInput)

	if userInput <= 0 {
		return 0, errors.New("inputan harus berupa angka")
	}

	return userInput, nil
}
