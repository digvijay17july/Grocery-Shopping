package model

type CategoryDescription struct {
	Model
CategoryId uint `json:"category_id"`
Description string `json:"description"`

	MetaDescription string `json:"MetaDescription"`
	Name string `json:"name"`
	MetaKeyword string `json:"meta_keyword"`
}
