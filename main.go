package main

import (
	"fmt"

	"github.com/gosimple/slug"
)

func main() {
	text := slug.Make("E-wallet Dana")
	fmt.Println(text)
}
