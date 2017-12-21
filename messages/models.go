package messages

import (
	"github.com/jinzhu/gorm"
)

type Message struct {
	gorm.Model `json:"-"`
	Name       string `json:"name" gorm:"size:255;not null"`
	Email      string `json:"email" gorm:"size:255;not null"`
	Date       int64  `json:"date" gorm:"not null"`
	Text       string `json:"text" gorm:"not null"`
}
