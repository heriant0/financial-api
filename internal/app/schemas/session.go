package schemas

type LoginRequest struct {
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required,min=8,alphanum" json:"password"`
}

type LoginResponse struct {
	Message      string `json:"message"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenRequest struct {
	RefreshToken string
	UserId       int
}

type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}
