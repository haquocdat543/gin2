package user

type CreateUserDTO struct {
	Name     string `json:"name" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=8,max=20"`
	Email    string `json:"email" binding:"required,email"`
}

type GetUserInfoDTO struct {
	Name string `uri:"name" binding:"required,min=3,max=20"`
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
	Name    string  `json:"name"`
	Dob     *string `json:"dob"`
	Role    *string `json:"role"`
	Address *string `json:"address"`
}

type PutUserDTO struct {
	Name    string `json:"name" binding:"required"`
	Dob     string `json:"dob" binding:"required"`
	Role    string `json:"role" binding:"required"`
	Address string `json:"address" binding:"required"`
}
