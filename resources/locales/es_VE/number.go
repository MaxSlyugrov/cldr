package es_VE

import "github.com/MaxSlyugrov/cldr"

var (
	symbols = cldr.Symbols{Decimal: ",", Group: ".", Negative: "", Percent: "", PerMille: ""}
	formats = cldr.NumberFormats{Decimal: "", Currency: "¤#,##0.00;¤-#,##0.00", Percent: ""}
)
