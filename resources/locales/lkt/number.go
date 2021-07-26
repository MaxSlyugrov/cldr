package lkt

import "github.com/MaxSlyugrov/cldr"

var (
	symbols = cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"}
	formats = cldr.NumberFormats{}
)
