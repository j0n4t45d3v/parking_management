package domain

type Vehicle struct {
	Band      string `json:"band"`
	Model     string `json:"model"`
	Color     string `json:"color"`
	Plate     string `json:"plate"`
	Type      string `json:"type"`
	Withdrawn bool   `json:"withdrawn"`
}
