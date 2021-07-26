package es_mx

import "github.com/MaxSlyugrov/cldr"

var Locale = &cldr.Locale{
	Locale: "es",
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
