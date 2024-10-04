package variant

import (
	"basic_trade/common"
	"basic_trade/product"
	"errors"
)

type Service interface {
	CreateVariant(input VariantInput) (*common.Variant, error)
	UpdateVariant(inputID GetVariantByUUIDInput, input VariantInput) (*common.Variant, error)
	DeleteVariant(uuid string, adminID int) (*common.Variant, error)
	GetVariants(page, limit int, search string) ([]common.Variant, error)
	GetVariantByUUID(input GetVariantByUUIDInput) (*common.Variant, error)
}

type service struct {
	repository        Repository
	productRepository product.Repository
}

func NewService(repository Repository, productRepository product.Repository) *service {
	return &service{repository, productRepository}
}

func (s *service) GetVariants(page, limit int, search string) ([]common.Variant, error) {
	variants, err := s.repository.GetAll(page, limit, search)
	if err != nil {
		return variants, err
	}

	return variants, nil
}

func (s *service) GetVariantByUUID(input GetVariantByUUIDInput) (*common.Variant, error) {
	variant, err := s.repository.GetByUUID(input.UUID)
	if err != nil {
		return nil, err
	}
	return variant, nil
}

func (s *service) CreateVariant(input VariantInput) (*common.Variant, error) {
	product, err := s.productRepository.GetByUUID(input.ProductUUID)
	if err != nil {
		return nil, errors.New("Product ID not found")
	}

	if product.AdminID != input.Admin.ID {
		return nil, errors.New("unauthorized")
	}

	variant := &common.Variant{
		VariantName: input.VariantName,
		Quantity:    input.Quantity,
		ProductID:   product.ID,
	}

	variant, err = s.repository.Save(variant)
	if err != nil {
		return nil, err
	}

	return variant, nil
}

func (s *service) UpdateVariant(inputID GetVariantByUUIDInput, input VariantInput) (*common.Variant, error) {
	variant, err := s.repository.GetByUUID(inputID.UUID)
	if err != nil {
		return variant, err
	}

	productID := variant.ProductID

	product, err := s.repository.GetProductByID(productID)
	if err != nil {
		return nil, errors.New("Product ID not found")
	}

	if product.AdminID != input.Admin.ID {
		return nil, errors.New("unauthorized")
	}

	variant.VariantName = input.VariantName
	variant.Quantity = input.Quantity

	updatedVariant, err := s.repository.Update(variant)
	if err != nil {
		return updatedVariant, err
	}

	return updatedVariant, nil
}

func (s *service) DeleteVariant(uuid string, adminID int) (*common.Variant, error) {
	variant, err := s.repository.GetByUUID(uuid)
	if err != nil {
		return nil, errors.New("Variant not found")
	}

	product, err := s.repository.GetProductByID(variant.ProductID)
	if err != nil {
		return nil, errors.New("Product not found")
	}

	if product.AdminID != adminID {
		return nil, errors.New("unauthorized")
	}

	deletedVariant, err := s.repository.Delete(variant.UUID)
	if err != nil {
		return nil, err
	}

	return deletedVariant, nil
}
