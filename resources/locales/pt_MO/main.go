package pt_MO

import "github.com/MaxSlyugrov/cldr"

var Locale = &cldr.Locale{
	Locale: "pt_MO",
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
