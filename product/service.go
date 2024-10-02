package product

import (
	"errors"
)

type Service interface {
	GetProducts(adminID int, page, limit int, search string) ([]Product, error)
	GetProductByUUID(input GetProductByUUIDInput) (*Product, error)
	CreateProduct(input ProductInput) (*Product, error)
	UpdateProduct(inputID GetProductByUUIDInput, input ProductInput) (*Product, error)
	DeleteProduct(uuid string, adminID int) (*Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetProducts(adminID int, page, limit int, search string) ([]Product, error) {
	if adminID != 0 {
		products, err := s.repository.GetByAdminID(adminID, page, limit, search)
		if err != nil {
			return products, err
		}

		return products, nil
	}

	products, err := s.repository.GetAll(page, limit, search)
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *service) GetProductByUUID(input GetProductByUUIDInput) (*Product, error) {
	product, err := s.repository.GetByUUID(input.UUID)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *service) CreateProduct(input ProductInput) (*Product, error) {
	product := &Product{
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

func (s *service) UpdateProduct(inputID GetProductByUUIDInput, input ProductInput) (*Product, error) {
	product, err := s.repository.GetByUUID(inputID.UUID)
	if err != nil {
		return product, err
	}

	if product.AdminID != input.Admin.ID {
		return product, errors.New("You don't have permission to update this product")
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

func (s *service) DeleteProduct(uuid string, adminID int) (*Product, error) {
	product, err := s.repository.GetByUUID(uuid)
	if err != nil {
		return product, err
	}

	if product.AdminID != adminID {
		return product, errors.New("You don't have permission to delete this product")
	}

	deletedProduct, err := s.repository.Delete(uuid)
	if err != nil {
		return deletedProduct, err
	}

	return deletedProduct, nil
}
