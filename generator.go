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
	return g.Generate(getRandomValue(g.randomLocale(), []string{"address", "country"}), "")
}
func (g Generator) Word() string {
	return g.Generate("{{text.word}}", "")
}
func (g Generator) Words(number int) []string {
	words := make([]string, 0)
	locs := g.locales
	g.locales = []string{g.randomLocale()}
	for i := 0; i < number; i++ {
		words = append(words, g.Word())
	}
	g.locales = locs
	return words
}

func (g Generator) Sentence(nbWords int) string {
	words := g.Words(nbWords)
	words[0] = strings.Title(words[0])
	sentence := strings.Join(words, " ")
	sentence += "."
	return sentence
}
func (g Generator) Sentences(nb int) []string {
	sentences := make([]string, 0)
	locs := g.locales
	g.locales = []string{g.randomLocale()}
	for i := 0; i < nb; i++ {
		sentences = append(sentences, g.Sentence(3+rand.Intn(5)))
	}
	g.locales = locs
	return sentences
}
func (g Generator) Paragraph(nbSentences int) string {
	return strings.Join(g.Sentences(nbSentences), " ")
}
func (g Generator) Paragraphs(nb int) []string {
	paragraphs := make([]string, 0)
	locs := g.locales
	g.locales = []string{g.randomLocale()}
	for i := 0; i < nb; i++ {
		paragraphs = append(paragraphs, g.Paragraph(3+rand.Intn(5)))
	}
	g.locales = locs
	return paragraphs
}
func (g Generator) Text(parargaphs int) string {
	locs := g.locales
	g.locales = []string{g.randomLocale()}
	text := strings.Join(g.Paragraphs(parargaphs), "\n")
	g.locales = locs
	return text
}
func (g Generator) Generate(tmpl string, loc string) string {
	if strings.Contains(tmpl, "{{") && strings.Contains(tmpl, "}}") {
		begin := strings.Index(tmpl, "{{")
		end := strings.Index(tmpl[begin:], "}}")
		toReplace := tmpl[begin+2 : begin+end]
		cats := strings.Split(toReplace, ".")
		if loc == "" {
			loc = g.randomLocale()
		}
		replacer := getRandomValue(loc, cats)
		tmpl = strings.Replace(tmpl, "{{"+toReplace+"}}", replacer, 1)
		tmpl = g.Generate(tmpl, loc)
	}
	tmpl = Numerify(tmpl)
	tmpl = Stringify(tmpl)
	return tmpl
}
