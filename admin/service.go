package admin

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input RegisterInput) (Admin, error)
	Login(input LoginInput) (Admin, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Register(input RegisterInput) (Admin, error) {
	admin := Admin{
		Name:  input.Name,
		Email: input.Email,
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return admin, err
	}

	admin.Password = string(passwordHash)
	newAdmin, err := s.repository.Save(&admin)
	if err != nil {
		return *newAdmin, err
	}

	return *newAdmin, nil

}

func (s *service) Login(input LoginInput) (Admin, error) {
	email := input.Email
	password := input.Password

	admin, err := s.repository.FindByEmail(email)
	if err != nil {
		return Admin{}, err
	}

	if admin.ID == 0 {
		return Admin{}, errors.New("Admin not found on that email")
	}

	log.Printf("Stored Hashed Password: %s", admin.Password)
	log.Printf("Input Password: %s", password)

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err != nil {
		return Admin{}, err
	}

	return *admin, nil
}
