package main

import "fmt"

func main() {
	data := "bat"
	lenght := len(data)
	var str = data[lenght-2 : lenght]
	data = Reverse(str)
	data = string(data[0]) + " " + string(data[1])
	fmt.Println(data)
}

func Reverse(str string) string {
	runeString := []rune(str)
	for i, j := 0, len(runeString)-1; i < j; i, j = i+1, j-1 {
		runeString[i], runeString[j] = runeString[j], runeString[i]
	}

	return string(runeString)
}
