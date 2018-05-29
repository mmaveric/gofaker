package gofaker

import (
	"math/rand"

	"github.com/mmaveric/gofaker/provider"
)

func getRandomValue(locale string, cat []string) string {
	if !check(locale, cat[0], cat[1]) {
		return ""
	}
	return provider.Locales[locale][cat[0]][cat[1]][rand.Intn(len(provider.Locales[locale][cat[0]][cat[1]]))]
}

func check(locale, category, subcategory string) bool {
	if loc, ok := provider.Locales[locale]; ok {
		if cat, ok := loc[category]; ok {
			if _, ok := cat[subcategory]; ok {
				return true
			}
			return false
		}
		return false
	}
	return false
}
