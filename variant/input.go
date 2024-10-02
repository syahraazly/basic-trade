package variant

type GetVariantByUUIDInput struct {
	UUID string `json:"uuid" form:"uuid" validate:"required"`
}

type VariantInput struct {
	VariantName string `json:"variant_name" form:"variant_name" validate:"required"`
	Quantity    int    `json:"quantity" form:"quantity" validate:"required"`
	ProductID   int    `json:"product_id" form:"product_id" validate:"required"`
}
