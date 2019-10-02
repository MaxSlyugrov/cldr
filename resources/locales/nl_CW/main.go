package nl_CW

import "github.com/rannoch/cldr"

var Locale = &cldr.Locale{
	Locale: "nl_CW",
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
