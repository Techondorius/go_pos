package model

// import "gorm.io/gorm"

type ItemMaster struct {
	ID int `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	BigCategory string `gorm:"not null"`
	SmallCategory string `gorm:"not null"`
	Grills []GrillPossible `gorm:"foreignKey:ItemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type GrillPossible struct{
	ID int `gorm:"primaryKey"`
	ItemID int `gorm:"not null"`
	MasterID int `gorm:"not null"`
}

type GrillMaster struct{
	ID int `gorm:"primaryKey"`
	Type string `gorm:"not null"`
	Name string `gorm:"not null"`
	Grills []GrillPossible `gorm:"foreignKey:MasterID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}