package ms_Latn_SG

import "github.com/MaxSlyugrov/cldr"

var Locale = &cldr.Locale{
	Locale: "ms_Latn_SG",
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
