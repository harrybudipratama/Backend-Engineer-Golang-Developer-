package main

import (
	"errors"
	"fmt"
)

func main() {
	// Masukkan inputan bilangan
	input, err := getUserInput("Masukkan bilangan: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	result := factorial(input)
	fmt.Printf("Faktorial dari %d adalah: %d\n", input, result)

	fmt.Print("Tekan Enter untuk keluar...")
	fmt.Scanln()
}

func getUserInput(input string) (int, error) {
	var userInput int
	fmt.Print(input)
	fmt.Scanln(&userInput)

	if userInput <= 0 {
		return 0, errors.New("inputan faktorial tidak terdefinisi untuk bilangan negatif dan text")
	}

	return userInput, nil
}

// Fungsi rekursif untuk menghitung faktorial
func factorial(n int) int {
	if n == 0 {
		return 1 // faktorial dari 0 adalah 1
	}
	return n * factorial(n-1) // Rekursi
}
