package product

import "basic_trade/variant"

type ProductFormatter struct {
	ID       string                     `json:"id"`
	Name     string                     `json:"name"`
	ImageURL string                     `json:"image_url"`
	Variants []variant.VariantFormatter `json:"variants"`
}

func FormatProduct(product Product) ProductFormatter {
	var variantsFormatted []variant.VariantFormatter
	if len(product.Variants) > 0 {
		variantsFormatted = variant.FormatVariants(product.Variants)
	} else {
		variantsFormatted = []variant.VariantFormatter{}
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
