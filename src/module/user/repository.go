package user

import "gorm.io/gorm"

type Repository interface {
	Create(
		user *User,
	) error

	FindAll() (
		[]User,
		error,
	)

	GetUserPassword(name string) (
		string,
		error,
	)

	UpdateUserPassword(name string, newPassword string) error

	DeleteUser(name string) error
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

func (r *repository) FindAll() (
	[]User,
	error,
) {
	var users []User
	err := r.db.Find(
		&users,
	).Error
	return users, err
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
