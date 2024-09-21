package main

import (
	"fmt"
	"strings"
)

func main() {
	// Masukkan inputan
	var input string
	fmt.Print("Input: ")
	fmt.Scanln(&input)

	if isPalindrome(input) {
		fmt.Println("Output: true")
	} else {
		fmt.Println("Output: false")
	}

	fmt.Print("Tekan Enter untuk keluar...")
	fmt.Scanln()
}

func isPalindrome(text string) bool {

	// Menghapus spasi dan mengubah semua huruf menjadi huruf kecil
	text = strings.ToLower(strings.ReplaceAll(text, " ", ""))

	// Menghitung panjang string
	n := len(text)

	// Memeriksa apakah string sama jika dibaca dari depan dan belakang
	for i := 0; i < n/2; i++ {
		if text[i] != text[n-1-i] {
			return false
		}
	}
	return true
}
