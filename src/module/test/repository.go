package test

import "gorm.io/gorm"

type Repository interface {
	Create(
		user *Test,
	) error

	FindAll() (
		[]Test,
		error,
	)
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
	user *Test,
) error {
	return r.db.Create(
		user,
	).Error
}

func (r *repository) FindAll() (
	[]Test,
	error,
) {
	var users []Test
	err := r.db.Find(
		&users,
	).Error
	return users, err
}
