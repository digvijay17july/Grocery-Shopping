package model

type Product struct {
	Model
	Rating uint `json:"rating"`
	Description string `json:"description"`
	Quantity int `json:"quantity"`
	MetaDescription string `json:"MetaDescription"`
	Name string `json:"name"`
	MetaKeyword string `json:"meta_keyword"`
	Cost int `json:"cost"`
	Currency string `json:"currency"`
}

