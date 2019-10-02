package de_LI

import "github.com/rannoch/cldr"

var Locale = &cldr.Locale{
	Locale: "de_LI",
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
