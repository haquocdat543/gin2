package user

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type CreateUserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func (r CreateUserDTO) Validate() error {
	return validation.ValidateStruct(
		&r,

		validation.Field(
			&r.Name,
			validation.Required.Error("Username is required"),
			validation.Length(3, 20).Error("Username must be between 3 and 20 characters"),
		),

		validation.Field(
			&r.Email,
			validation.Required.Error("Email is required"),
			is.Email.Error("Invalid email format"),
		),

		validation.Field(
			&r.Age,
			validation.Required.Error("Age is required"),
			validation.Min(13).Error("Age must be greater than 13"),
			validation.Max(80).Error("Age must be lesser than 80"),
		),

	)
}
