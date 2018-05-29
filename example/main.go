package main

import (
	"fmt"

	"github.com/mmaveric/gofaker"
)

func main() {
	gofaker.Seed(0)
	gofaker.SetLocales("kk_KZ")
	fmt.Println(gofaker.Generate("{{payment.bank}}"))
}
