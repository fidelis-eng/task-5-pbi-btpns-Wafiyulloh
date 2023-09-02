package models

import (
	"time"
)

type User struct {
	Id        string    `gorm:"primaryKey;not null" json:"id"`
	Username  string    `gorm:"type:varchar;not null"  json:"username"`
	Email     string    `gorm:"type:varchar;not null;unique;" json:"email"`
	Password  string    `gorm:"type:varchar;not null;" json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Photo     Photo     `gorm:"foreignKey:userid;references:Id; constraint:On;Update:CASCADE;OnDelete:SET NULL;"`
}
