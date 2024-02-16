package adapters

import (
	"gorm.io/gorm"
)

// Secondary adapter send to DB
type GormBookStore struct {
	db *gorm.DB
}
