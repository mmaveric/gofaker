package main

import (
	"fmt"

	"github.com/mmaveric/gofaker"
)

func main() {
	gofaker.Seed(0)
	gofaker.SetLocales("en_US")
	for i := 0; i < 10; i++ {
		fmt.Println(gofaker.Generate("{{phone_number.phone_number}}"))
	}
}
