package web

type TransactionCreateRequest struct {
	UserID           uint    `json:"user_id" validate:"required"`
	OnTheRoad        float64 `json:"on_the_road"`
	AdminFee         float64 `json:"admin_fee"`
	InstalmentAmount float64 `json:"instalment_amount"`
	InterestAmount   float64 `json:"interest_amount"`
	ProductID        uint    `json:"product_id"`
	Period           int     `json:"period"`
	TransactionType  string  `json:"transaction_type"`
}
