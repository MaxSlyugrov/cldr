package ta

import "github.com/MaxSlyugrov/cldr"

var Locale = &cldr.Locale{
	Locale: "ta",
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
