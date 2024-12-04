package web

type TokenResponse struct {
	// Fields
	ID           uint   `json:"id"`
	FullName     string `json:"user_name"`
	LegalName    string `json:"legal_name"`
	Role         string `json:"role"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
