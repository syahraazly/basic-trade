package product

import "basic_trade/admin"

type GetProductByUUIDInput struct {
	UUID string `uri:"uuid" validate:"required"`
}

type ProductInput struct {
	Name     string `json:"name" form:"name" validate:"required"`
	ImageURL string `json:"image_url" form:"image_url" validate:"required,url"`
	Admin    admin.Admin
}
