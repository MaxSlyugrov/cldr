package zh_Hant_MO

import "github.com/MaxSlyugrov/cldr"

var Locale = &cldr.Locale{
	Locale: "zh_Hant_MO",
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
