package pa_Arab

import "github.com/MaxSlyugrov/cldr"

var Locale = &cldr.Locale{
	Locale: "pa_Arab",
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
