package messages

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func insert(db *gorm.DB, msg *Message) (*Message, error) {
	db.Create(msg)
	return msg, nil
}

func findAll(db *gorm.DB) ([]*Message, error) {
	var messages []*Message
	if err := db.Find(&messages).Error; err != nil {
		fmt.Printf("find all messages err, %+v", err)
		return []*Message{}, err
	}
	return messages, nil
}
