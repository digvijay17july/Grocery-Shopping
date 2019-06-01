package model

type Categories struct{
	Model
	Image string `json:"image"`
	CategoryDescriptionId uint `gorm:"foreignkey:CategoryId"  json:"categoryDescriptionId"`
	Top bool `json:"top"`
	Column int `json:"column"`
	SortOrder int `json:"sort_order"`
	TempShipId int `json:"temp_ship_id"`
}
