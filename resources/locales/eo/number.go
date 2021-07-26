package eo

import "github.com/MaxSlyugrov/cldr"

var (
	symbols = cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "−", Percent: "%", PerMille: "‰"}
	formats = cldr.NumberFormats{Decimal: "#,##0.###", Currency: "", Percent: "#,##0%"}
)
