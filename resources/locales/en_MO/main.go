package en_MO

import "github.com/MaxSlyugrov/cldr"

var Locale = &cldr.Locale{
	Locale: "en_MO",
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
