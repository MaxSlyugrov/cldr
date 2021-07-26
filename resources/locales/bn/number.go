package bn

import "github.com/MaxSlyugrov/cldr"

var (
	symbols = cldr.Symbols{}
	formats = cldr.NumberFormats{Decimal: "#,##,##0.###", Currency: "#,##,##0.00¤", Percent: "#,##,##0%"}
)
