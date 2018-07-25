package gofaker

import (
	"math/rand"
	"strings"
)

type Generator struct {
	locales []string
}

var defaultGenerator Generator = Generator{[]string{"en_US"}}

func NewGenerator(locales ...string) Generator {
	if len(locales) == 0 {
		return defaultGenerator
	}
	return Generator{locales: locales}
}

func (g Generator) randomLocale() string {
	return g.locales[rand.Intn(len(g.locales))]
}

func (g Generator) Int(min, max int) int {
	if min == max {
		return min
	}
	if min >= max {
		return rand.Intn(min-max) + max
	}
	return rand.Intn(max-min) + min
}

func (g Generator) Country() string {
	return g.Generate(getRandomValue(g.randomLocale(), []string{"address", "country"}))
}
func (g Generator) Word() string {
	return g.Generate(getRandomValue(g.randomLocale(), []string{"lorem", "word"}))
}
func (g Generator) Words(number int) []string {
	words := make([]string, 0)
	for i := 0; i < number; i++ {
		words = append(words, g.Word())
	}
	return words
}

func (g Generator) Sentence(nbWords int) string      { return "" }           //todo
func (g Generator) Sentences(nb int) []string        { return []string{""} } //todo
func (g Generator) Paragraph(nbSentences int) string { return "" }           //todo
func (g Generator) Paragraphs(nb int) []string       { return []string{""} } //todo
func (g Generator) Text(maxNbChars int) string       { return "" }           //todo
func (g Generator) Generate(tmpl string) string {
	if strings.Contains(tmpl, "{{") && strings.Contains(tmpl, "}}") {
		begin := strings.Index(tmpl, "{{")
		end := strings.Index(tmpl[begin:], "}}")
		toReplace := tmpl[begin+2 : begin+end]
		cats := strings.Split(toReplace, ".")
		replacer := getRandomValue(g.randomLocale(), cats)
		tmpl = strings.Replace(tmpl, "{{"+toReplace+"}}", replacer, 1)
		tmpl = g.Generate(tmpl)
	}
	tmpl = Numerify(tmpl)
	tmpl = Stringify(tmpl)
	return tmpl
}
