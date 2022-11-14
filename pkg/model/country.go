package model

type Country struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name"`
	ShortName     string `json:"short_name,omitempty"`
	Continent     string `json:"continent"`
	IsOperational bool   `json:"is_operational"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updateAt"`
}
