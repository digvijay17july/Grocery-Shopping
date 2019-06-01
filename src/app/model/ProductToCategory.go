package model

type ProductToCategory struct {
	Model
	Product Product `gorm:"association_foreignkey:Refer"`
	Category Categories  `gorm:"association_foreignkey:Refer"`
}

