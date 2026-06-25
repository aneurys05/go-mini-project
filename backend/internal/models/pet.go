package models

type Pet struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Species     string `json:"species"`
	Breed       string `json:"breed"`
	Age         int    `json:"age"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
}
