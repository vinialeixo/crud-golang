package request

type UserRequest struct {
	ID       string `json:"id" `
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*"`
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required,min=4,max=100"`
	Age      int8   `json:"age" binding:"required,min=1,max=140"`
}

type UserUpdateRequest struct {
	Name string `json:"name" binding:"omitempty,min=4,max=100"`
	Age  int8   `json:"age" binding:"omitempty,min=1,max=140"`
}
