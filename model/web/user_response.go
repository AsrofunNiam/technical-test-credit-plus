package web

type UserResponse struct {
	// Fields
	ID           uint    `json:"id"`
	FullName     string  `json:"full_name"`
	LegalName    string  `json:"legal_name"`
	PlaceOfBirth string  `json:"place_of_birth"`
	DateOfBirth  string  `json:"date_of_birth"`
	Salary       float64 `json:"salary"`
	Role         string  `json:"role"`
}

type UserLoginRequest struct {
	AccessToken string `json:"access_token"`
}
