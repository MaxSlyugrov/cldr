package sw_UG

import "github.com/rannoch/cldr"

var Locale = &cldr.Locale{
	Locale: "sw_UG",
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
