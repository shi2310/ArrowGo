package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Records ...
type Records struct {
	gorm.Model
	Sn          string `gorm:"size(255)"` // Device.Sn
	IsOver      bool
	IsTemp      bool
	Photo       string `gorm:"type:text"`
	Temperature float64
	UnlockTime  time.Time
	UnlockType  int
	UserID      int
}
