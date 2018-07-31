package main

import (
	"fmt"

	"github.com/mmaveric/gofaker"
)

func main() {
	gofaker.Seed(0)
	gofaker.SetLocales("en_US", "ru_RU", "kk_KZ")
	for i := 0; i < 100; i++ {
		fmt.Println(gofaker.Number(12, true))
	}
}
