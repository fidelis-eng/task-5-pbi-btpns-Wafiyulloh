package models

type Photo struct {
	Id       string `gorm:"primaryKey;not null" valid:"required" json:"id"`
	Title    string `gorm:"type:text;not null" json:"title"`
	Caption  string `gorm:"type:text;not null" json:"caption"`
	PhotoUrl string `gorm:"type:text;not null" json:"photourl"`
	Userid   string `gorm:"references:Id;not null;" json:"userid"`
}
