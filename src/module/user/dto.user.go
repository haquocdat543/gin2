package user

type CreateUserDTO struct {
	Name     string `json:"name" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=8,max=20"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginDTO struct {
	Name     string `json:"name" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

type UpdatePasswordDTO struct {
	Name        string `json:"name" binding:"required,min=3,max=20"`
	Password    string `json:"password" binding:"required,min=8,max=20"`
	NewPassword string `json:"new_password" binding:"required,min=8,max=20"`
}

type DeleteUserDTO struct {
	Name     string `json:"name" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

// Other
type PatchUserDTO struct {
	Dob     *string `json:"dob"`
	Role    *string `json:"role"`
	Address *string `json:"address"`
}

type PutUserDTO struct {
	Dob     string `json:"dob" binding:"required"`
	Role    string `json:"role" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type PatchDeleteDTO struct {
	Dob     *bool `json:"dob"`
	Role    *bool `json:"role"`
	Address *bool `json:"address"`
}
