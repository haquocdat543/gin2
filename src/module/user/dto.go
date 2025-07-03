package user

type CreateUserDTO struct {
	Name  string `json:"name" binding:"required,min=3"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"required,gte=0,lte=130"`
}

