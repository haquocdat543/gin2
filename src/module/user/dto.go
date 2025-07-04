package user

type CreateUserDTO struct {
	Name     string `json:"name" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=8,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Age      int    `json:"age" binding:"required,gte=0,lte=130"`
}
