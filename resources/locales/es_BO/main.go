package es_BO

import "github.com/rannoch/cldr"

var Locale = &cldr.Locale{
	Locale: "es_BO",
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
