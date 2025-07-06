package user

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type Service interface {
	CreateUser(user *User) error

	GetUser(name string) (User, error)

	GetUsers() ([]User, error)

	CheckUserExist(name string) bool

	Login(name string, passwors string) error

	UpdateUserPassword(name string, newPassword string) error

	DeleteUser(name string) error

	UpdateUser(user *User) error

	PatchDeleteUser(user *User, fields map[string]interface{}) error
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

func (s *service) GetUser(name string) (User, error) {
	return s.repo.GetUser(name)
}

func (s *service) GetUsers() (
	[]User,
	error,
) {
	return s.repo.GetUsers()
}

func (s *service) CheckUserExist(name string) bool {

	return s.repo.CheckUserExist(name)

}

func (s *service) Login(
	name string,
	plainPassword string,
) error {
	hashedPassword, err := s.repo.GetUserPassword(name)
	if err != nil {

		if strings.Contains(
			err.Error(),
			"record not found",
		) {

			// Could be user not found
			return fmt.Errorf("User Not Found")

		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		// Don't leak if it's a bad password or a hash mismatch
		return fmt.Errorf(ErrInvalidPassword)
	}

	return nil // login success
}

func (s *service) UpdateUserPassword(name string, newPassword string) error {

	err := s.repo.UpdateUserPassword(name, newPassword)
	if err != nil {
		// Could be user not found
		return fmt.Errorf("Update password failed: %w", err)
	}

	return nil
}

func (s *service) DeleteUser(name string) error {

	err := s.repo.DeleteUser(name)
	if err != nil {
		// Could be user not found
		return fmt.Errorf("Update password failed: %w", err)
	}

	return nil
}

func (s *service) UpdateUser(user *User) error {

	err := s.repo.UpdateUser(user)
	if err != nil {
		// Could be user not found
		return fmt.Errorf("Update password failed: %w", err)
	}

	return nil
}

func (s *service) PatchDeleteUser(user *User, fields map[string]interface{}) error {

	err := s.repo.PatchDeleteUser(user, fields)
	if err != nil {
		// Could be user not found
		return fmt.Errorf("Patch delete failed: %w", err)
	}

	return nil
}
