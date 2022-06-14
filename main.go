package main

import (
	"fmt"

	"github.com/xlzd/gotp"
)

func main() {
	secretLength := 4
	data := gotp.RandomSecret(secretLength)
	fmt.Println(data)
}
