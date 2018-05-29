package provider

var kk_KZ map[string]map[string][]string = map[string]map[string][]string{
	"address": map[string][]string{
		"city_suffix":     []string{"қаласы"},
		"region_suffix":   []string{"облысы"},
		"street_suffix":   []string{"көшесі", "даңғылы"},
		"building_number": []string{"%##"},
		"postcode":        []string{"0#####"},
		"country":         []string{"Қазақстан", "Ресей"},
		"region":          []string{"Алматы", "Ақтау", "Ақтөбе", "Астана", "Атырау", "Байқоңыр", "Қарағанды", "Көкшетау", "Қостанай", "Қызылорда", "Маңғыстау", "Павлодар", "Петропавл", "Талдықорған", "Тараз", "Орал", "Өскемен", "Шымкент"},
		"city":            []string{"Алматы", "Ақтау", "Ақтөбе", "Астана", "Атырау", "Байқоңыр", "Қарағанды", "Көкшетау", "Қостанай", "Қызылорда", "Маңғыстау", "Павлодар", "Петропавл", "Талдықорған", "Тараз", "Орал", "Өскемен", "Шымкент"},
		"street":          []string{"Абай", "Гоголь", "Кенесары", "Бейбітшілік", "Достық", "Бұқар жырау"},
		"address":         []string{"{{address.postcode}}, {{address.region}} {{address.region_suffix}}, {{address.city}} {{address.city_suffix}}, {{address.street}} {{address.street_suffix}}, {{address.building_number}}"},
		"street_address":  []string{"{{address.street}} {{address.street_suffix}}, {{address.building_number}}"},
	},
	"color": map[string][]string{
		"color_name": []string{"қара", "қою қызыл", "жасыл", "қара көк", "сарғыш түс", "күлгін", "көк", "көк", "күміс", "сұр", "сары", "қызылкүрең түс", "теңіз толқыны түс", "ақ"},
	},
	"payment": map[string][]string{
		"bank": []string{"Қазкоммерцбанк", "Халық Банкі"},
	},
}
