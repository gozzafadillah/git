package main

import (
	"fmt"
	"regexp"
)

func main() {

	var text = "+62895631948686"
	// var text = "081320503262"
	var regex, _ = regexp.Compile(`[a-z]+`)

	regex.FindStringIndex(text)

	lenght := len(text)
	var str = text[3:lenght]
	fmt.Println("0" + str)
	status := check(str)
	fmt.Println("status :", status)
}

func check(phone string) bool {
	regex, _ := regexp.Compile(`^(0|62|\+62)(8[1-35-9]\d{7,10}|2[124]\d{7,8}|619\d{8}|2(?:1(?:14|500)|2\d{3})\d{3}|61\d{5,8}|(?:2(?:[35][1-4]|6[0-8]|7[1-6]|8\d|9[1-8])|3(?:1|[25][1-8]|3[1-68]|4[1-3]|6[1-3568]|7[0-469]|8\d)|4(?:0[1-589]|1[01347-9]|2[0-36-8]|3[0-24-68]|43|5[1-378]|6[1-5]|7[134]|8[1245])|5(?:1[1-35-9]|2[25-8]|3[124-9]|4[1-3589]|5[1-46]|6[1-8])|6(?:[25]\d|3[1-69]|4[1-6])|7(?:02|[125][1-9]|[36]\d|4[1-8]|7[0-36-9])|9(?:0[12]|1[013-8]|2[0-479]|5[125-8]|6[23679]|7[159]|8[01346]))\d{5,8})`)
	var isMatch = regex.MatchString(phone)
	return isMatch
}
