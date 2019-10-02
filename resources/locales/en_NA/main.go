package en_NA

import "github.com/rannoch/cldr"

var Locale = &cldr.Locale{
	Locale: "en_NA",
	Number: cldr.Number{
		Symbols:    symbols,
		Formats:    formats,
		Currencies: currencies,
	},
	Calendar:   calendar,
	PluralRule: pluralRule,
}

func init() {
	cldr.RegisterLocale(Locale)
}
