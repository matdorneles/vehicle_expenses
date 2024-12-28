package models

type Vehicle struct {
	ID        string `json:"id"`
	Maker     string `json:"maker"`
	Model     string `json:"model"`
	Year      string `json:"year"`
	CreatedAt string `json:"created_at"`
}
