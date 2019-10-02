package ca_ES_VALENCIA

import "github.com/rannoch/cldr"

var Locale = &cldr.Locale{
	Locale: "ca_ES_VALENCIA",
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
