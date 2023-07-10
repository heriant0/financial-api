package schemas

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=8,alphanum"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

type UserProfileRequest struct {
	ID int
}

type UserProfielResponse struct {
	ID        int
	Email     string `json:"email"`
	Username  string `json:"username"`
	Usersince string `json:"usersince"`
}
