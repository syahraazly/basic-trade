package variant

import "basic_trade/admin"

type GetVariantByUUIDInput struct {
	UUID string `uri:"uuid" validate:"required"`
}

type VariantInput struct {
	VariantName string `json:"variant_name" form:"variant_name" validate:"required"`
	Quantity    int    `json:"quantity" form:"quantity" validate:"required"`
	ProductUUID string `json:"product_id" form:"product_id" validate:"required"`
	Admin       admin.Admin
}

