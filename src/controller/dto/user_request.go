package dto

type UserRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Age      int8   `json:"age"`
}
