package web

type LimitResponse struct {
	ID          uint         `json:"id"`
	UserID      uint         `json:"user_id"`
	User        UserResponse `json:"user"`
	Tenor       int          `json:"tenor"`
	LimitAmount float64      `json:"limit_amount"`
	Available   bool         `json:"available"`
}
