package user

import (
	"gorm.io/gorm"
)

var fields = []string{
	"email",
	"dob",
	"role",
	"address",
}

type Repository interface {
	Create(user *User) error

	GetUser(name string) (User, error)

	GetUsers() ([]User, error)

	CheckUserExist(name string) bool

	GetUserPassword(name string) (string, error)

	UpdateUserPassword(name string, newPassword string) error

	DeleteUser(name string) error

	UpdateUser(user *User) error

	PatchDeleteUser(user *User, fields map[string]interface{}) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(
	db *gorm.DB,
) Repository {
	return &repository{
		db,
	}
}

func (r *repository) Create(
	user *User,
) error {
	return r.db.Create(
		user,
	).Error
}

func (r *repository) GetUser(name string) (User, error) {
	var user User

	err := r.db.Select(fields).Where("name = ?", name).Find(&user).Error
	return user, err
}

func (r *repository) GetUsers() (
	[]User,
	error,
) {
	var users []User
	err := r.db.Select(fields).Find(
		&users,
	).Error
	return users, err
}

func (r *repository) CheckUserExist(name string) bool {
	var user User
	err := r.db.First(&user, "name = ?", name).Error
	if err != nil {
		return false
	}
	return true
}

func (r *repository) GetUserPassword(name string) (string, error) {
	var user User
	err := r.db.First(&user, "name = ?", name).Error
	if err != nil {
		return "", err
	}
	return user.Password, nil
}

func (r *repository) UpdateUserPassword(name string, newPassword string) error {
	var user User

	err := r.db.First(&user, "name = ?", name).Error
	if err != nil {
		return err
	}

	user.Password = newPassword
	r.db.Save(&user)

	return nil
}

func (r *repository) DeleteUser(name string) error {
	var user User

	err := r.db.Unscoped().Delete(&user, "name = ?", name).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateUser(user *User) error {

	var existUser User

	err := r.
		db.
		Model(&existUser).
		Where("name = ?", user.Name).
		Updates(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) PatchDeleteUser(user *User, fields map[string]interface{}) error {

	var existUser User

	err := r.
		db.
		Model(&existUser).
		Where("name = ?", user.Name).
		Updates(fields).Error
	if err != nil {
		return err
	}

	return nil
}
