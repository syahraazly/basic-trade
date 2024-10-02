package variant

type VariantFormatter struct {
	VariantName string `json:"variant_name"`
	Quantity    int    `json:"quantity"`
	ProductID   int    `json:"product_id"`
}

func FormatVariant(variant Variant) VariantFormatter {
	formatter := VariantFormatter{
		VariantName: variant.VariantName,
		Quantity:    variant.Quantity,
		ProductID:   variant.ProductID,
	}
	return formatter
}

func FormatVariants(variants []Variant) []VariantFormatter {
	var formatters []VariantFormatter
	for _, variant := range variants {
		formatter := FormatVariant(variant) 
		formatters = append(formatters, formatter)
	}
	return formatters
}
