package dto

import "time"

type ShowEmployee struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CabangID  *uint   `json:"cabang_id"`
	CreatedAt time.Time `json:"created_at"`
	Cabang *CabangResponse `json:"cabang"`
}

type CabangResponse struct {
	ID uint `json:"id"`
	Name string `json:"cabang_name"`
	Location string `json:"location"`
}