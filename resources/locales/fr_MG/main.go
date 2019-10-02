package fr_MG

import "github.com/rannoch/cldr"

var Locale = &cldr.Locale{
	Locale: "fr_MG",
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
