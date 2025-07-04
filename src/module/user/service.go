package user

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(
		user *User,
	) error

	GetAllUsers() (
		[]User,
		error,
	)

	Login(
		name string,
		passwors string,
	) error
}

type service struct {
	repo Repository
}

func NewService(
	r Repository,
) Service {
	return &service{
		repo: r,
	}
}

func (s *service) CreateUser(
	user *User,
) error {
	return s.repo.Create(
		user,
	)
}

func (s *service) GetAllUsers() (
	[]User,
	error,
) {
	return s.repo.FindAll()
}

func (s *service) Login(
	name string,
	plainPassword string,
) error {
	hashedPassword, err := s.repo.GetUserPassword(name)
	if err != nil {
		// Could be user not found
		return fmt.Errorf("authentication failed: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		// Don't leak if it's a bad password or a hash mismatch
		return fmt.Errorf("invalid credentials")
	}

	return nil // login success
}
