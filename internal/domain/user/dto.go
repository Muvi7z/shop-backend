package user

type UserDTO struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
