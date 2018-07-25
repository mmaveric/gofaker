package gofaker

import (
	"math/rand"
	"time"
)

// Seed random. Setting seed to 0 will use time.Now().UnixNano()
func Seed(seed int64) {
	if seed == 0 {
		rand.Seed(time.Now().UTC().UnixNano())
	} else {
		rand.Seed(seed)
	}
}
func SetLocales(locales ...string) {
	defaultGenerator.locales = locales
}

func Word() string {
	return defaultGenerator.Word()
}
func Words(number int) []string {
	return defaultGenerator.Words(number)
}
func Sentence(nbWords int) string {
	return defaultGenerator.Sentence(nbWords)
}
func Sentences(nb int) []string {
	return defaultGenerator.Sentences(nb)
}
func Paragraph(nbSentences int) string {
	return defaultGenerator.Paragraph(nbSentences)
}
func Paragraphs(nb int) []string {
	return defaultGenerator.Paragraphs(nb)
}
func Text(maxNbChars int) string {
	return defaultGenerator.Text(maxNbChars)
}
func Country() string {
	return defaultGenerator.Country()
}
func Generate(tmpl string) string {
	return defaultGenerator.Generate(tmpl)
}

func Numerify(tmpl string) string {
	if tmpl == "" {
		return tmpl
	}
	byteTmpl := []byte(tmpl)
	hash := []byte("#")[0]
	percent := []byte("%")[0]
	numbers := []byte("0123456789")
	for i := 0; i < len(byteTmpl); i++ {
		if byteTmpl[i] == hash {
			byteTmpl[i] = numbers[rand.Intn(9)]
		} else if byteTmpl[i] == percent {
			byteTmpl[i] = numbers[rand.Intn(8)+1]
		}
	}

	return string(byteTmpl)
}

func Stringify(tmpl string) string {
	if tmpl == "" {
		return tmpl
	}
	byteTmpl := []byte(tmpl)
	question := []byte("?")[0]
	letters := []byte("abcdefghijklmnopqrstuvwxyz")
	for i := 0; i < len(byteTmpl); i++ {
		if byteTmpl[i] == question {
			byteTmpl[i] = letters[rand.Intn(26)]
		}
	}

	return string(byteTmpl)
}
