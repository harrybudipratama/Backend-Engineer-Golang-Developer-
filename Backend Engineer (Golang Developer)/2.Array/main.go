package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	// Masukkan jumlah elemenn array
	input, err := getUserInput("Masukkan jumlah elemen array: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	arr := make([]int, input)
	fmt.Println("Masukkan elemen array:")

	for i := 0; i < input; i++ {
		// Masukkan value perelemen
		element, err := getUserInputElemen("Elemen ke-%d", i+1)
		if err != nil {
			fmt.Println(err)
			return
		}
		arr[i] = element
	}

	result := findMax(arr)
	fmt.Println("Nilai maksimum:", result)

	fmt.Print("Tekan Enter untuk keluar...")
	fmt.Scanln()

}

func getUserInput(input string) (int, error) {
	var userInput int
	fmt.Print(input)
	fmt.Scanln(&userInput)

	if userInput <= 0 {
		return 0, errors.New("inputan harus berupa angka tidak boleh text atau angka negatif")
	}

	return userInput, nil
}

func getUserInputElemen(prompt string, index int) (int, error) {
	var input string
	fmt.Printf(prompt, index)
	fmt.Print(": ")
	fmt.Scanln(&input)

	// Mengonversi input ke integer
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("inputan harus berupa angka")
	}

	return num, nil
}

func findMax(arr []int) int {
	max := arr[0]

	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}
