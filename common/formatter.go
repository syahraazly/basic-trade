package common

type ProductFormatter struct {
	ID       string             `json:"id"`
	Name     string             `json:"name"`
	ImageURL string             `json:"image_url"`
	Variants []VariantFormatter `json:"variants"`
}

func FormatProduct(product Product) ProductFormatter {
	var variantsFormatted []VariantFormatter
	if len(product.Variants) > 0 {
		variantsFormatted = FormatVariants(product.Variants)
	} else {
		variantsFormatted = []VariantFormatter{}
	}

	formatter := ProductFormatter{
		ID:       product.UUID,
		Name:     product.Name,
		ImageURL: product.ImageURL,
		Variants: variantsFormatted,
	}
	return formatter
}

func FormatProducts(products []Product) []ProductFormatter {
	var formatters []ProductFormatter
	for _, product := range products {
		formatter := FormatProduct(product)
		formatters = append(formatters, formatter)
	}
	return formatters
}

func FormatProductDetail(product Product) ProductFormatter {
	formatter := FormatProduct(product)
	return formatter
}

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
