package es_GQ

import "github.com/rannoch/cldr"

var Locale = &cldr.Locale{
	Locale: "es_GQ",
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
