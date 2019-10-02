package fr_CH

import "github.com/rannoch/cldr"

var (
	symbols = cldr.Symbols{Decimal: ".", Group: "", Negative: "", Percent: "", PerMille: ""}
	formats = cldr.NumberFormats{Decimal: "", Currency: "¤\u00a0#,##0.00;¤-#,##0.00", Percent: ""}
)
