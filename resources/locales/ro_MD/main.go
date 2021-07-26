package ro_MD

import "github.com/MaxSlyugrov/cldr"

var Locale = &cldr.Locale{
	Locale: "ro_MD",
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
