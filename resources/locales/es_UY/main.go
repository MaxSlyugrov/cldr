package es_UY

import "github.com/MaxSlyugrov/cldr"

var Locale = &cldr.Locale{
	Locale: "es_UY",
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
