package es_NI

import "github.com/MaxSlyugrov/cldr"

var Locale = &cldr.Locale{
	Locale: "es_NI",
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
