package model

type ItemMaster struct {
	ID int `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Grills []GrillPossible `gorm:"foreignKey:ItemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Children []ItemMaster `gorm:"-:all"`
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