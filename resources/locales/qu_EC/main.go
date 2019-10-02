package qu_EC

import "github.com/rannoch/cldr"

var Locale = &cldr.Locale{
	Locale: "qu_EC",
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
