package model

type ItemMaster struct {
	ID int `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Price int `gorm:"not null"`
	BigCategory string `gorm:"not null"`
	SmallCategory string `gorm:"not null"`
	Grills []GrillPossible `gorm:"foreignKey:ItemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Children []ItemMaster `gorm:"-:all"`
}

type ItemPossibleByID struct{
	ID int `gorm:"primaryKey"`
	ItemID int `gorm:"primaryKey"`
	ChildID int `gorm:"primaryKey"`
}

type ItemPossibleByCategory struct{
	ID int `gorm:"primaryKey"`
	
}

type GrillPossible struct{
	ID int `gorm:"primaryKey"`
	ItemID int `gorm:"primaryKey"`
	MasterID int `gorm:"not null"`
}

type GrillMaster struct{
	ID int `gorm:"primaryKey"`
	Type string `gorm:"not null"`
	Name string `gorm:"not null"`
	Grills []GrillPossible `gorm:"foreignKey:MasterID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}