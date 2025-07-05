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
