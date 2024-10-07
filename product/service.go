package product

import (
	"basic_trade/common"
	"errors"
)

type Service interface {
	GetProducts(adminID int, offset, limit int, search string) ([]common.Product, error)
	GetProductByUUID(input GetProductByUUIDInput) (*common.Product, error)
	CreateProduct(input ProductInput) (*common.Product, error)
	UpdateProduct(inputID GetProductByUUIDInput, input ProductInput) (*common.Product, error)
	DeleteProduct(uuid string, adminID int) (*common.Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetProducts(adminID int, offset, limit int, search string) ([]common.Product, error) {
	if adminID != 0 {
		products, err := s.repository.GetByAdminID(adminID, offset, limit, search)
		if err != nil {
			return products, err
		}

		return products, nil
	}

	products, err := s.repository.GetAll(offset, limit, search)
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *service) GetProductByUUID(input GetProductByUUIDInput) (*common.Product, error) {
	product, err := s.repository.GetByUUID(input.UUID)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *service) CreateProduct(input ProductInput) (*common.Product, error) {
	product := &common.Product{
		Name:     input.Name,
		ImageURL: input.ImageURL,
		AdminID:  input.Admin.ID,
	}

	product, err := s.repository.Save(product)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *service) UpdateProduct(inputID GetProductByUUIDInput, input ProductInput) (*common.Product, error) {
	product, err := s.repository.GetByUUID(inputID.UUID)
	if err != nil {
		return product, err
	}

	if product.AdminID != input.Admin.ID {
		return product, errors.New("unauthorized")
	}

	product.Name = input.Name

	if input.ImageURL != "" {
		product.ImageURL = input.ImageURL
	}

	updatedProduct, err := s.repository.Update(product)
	if err != nil {
		return updatedProduct, err
	}

	return updatedProduct, nil
}

func (s *service) DeleteProduct(uuid string, adminID int) (*common.Product, error) {
	product, err := s.repository.GetByUUID(uuid)
	if err != nil {
		return product, err
	}

	if product.AdminID != adminID {
		return product, errors.New("unauthorized")
	}

	deletedProduct, err := s.repository.Delete(uuid)
	if err != nil {
		return deletedProduct, err
	}

	return deletedProduct, nil
}
