package model

// import "gorm.io/gorm"

type ItemMaster struct {
	ID int
	Name string
	BigCategory string
	SmallCategory string
	Grills []GrillPossible `gorm:"foreignKey:ItemID;references:ID"`
}

type GrillPossible struct{
	ID int
	ItemID int
	MasterID int
}

type GrillMaster struct{
	ID int
	Type string
	Name string
	Grills []GrillPossible `gorm:"foreignKey:MasterID;references:ID"`
}