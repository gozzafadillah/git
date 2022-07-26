package main

import (
	"fmt"
	"sort"
)

func Sum(data []int) int {
	hasil := 0
	for i := 0; i < len(data); i++ {
		hasil += data[i]
	}
	return hasil
}

func CekEven(number int) bool {
	if number%2 != 0 {
		return false
	}
	return true
}

func main() {
	data := []int{1, 3, 5, 7}
	// sort array
	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
	fmt.Println(data)
	// hitung jarak antar data array
	counter := 0
	dataCounter := []int{}
	for j := 1; j <= len(data)-1; j++ {
		counter = data[j] - data[j-1]
		dataCounter = append(dataCounter, counter)
	}
	// menambah semua selisih
	sum := Sum(dataCounter)
	fmt.Println("data counter :", dataCounter)
	fmt.Println("data sum :", sum)

	// Cek total itu genap
	status := CekEven(sum)
	fmt.Println("Even ? :", status)

}
