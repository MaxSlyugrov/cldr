package fr_SY

import "github.com/MaxSlyugrov/cldr"

var Locale = &cldr.Locale{
	Locale: "fr_SY",
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
