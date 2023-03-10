package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,containsany=!@#?,min=6"`
	Name     string `json:"name" binding:"required,min=4,max=100"`
	Age      int    `json:"age" binding:"required,numeric,min=1"`
}