package nd

import "github.com/MaxSlyugrov/cldr"

var Locale = &cldr.Locale{
	Locale: "nd",
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
