package model

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	CreatedAt int64 `gorm:"column:created_at;default:0;NOT NULL"`
	UpdatedAt int64 `gorm:"column:updated_at;default:0;NOT NULL"`
	DeletedAt int64 `gorm:"column:deleted_at;default:0;NOT NULL"`
}

func (record *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if record.CreatedAt == 0 {
		record.CreatedAt = time.Now().Unix()
	}
	if record.UpdatedAt == 0 {
		record.UpdatedAt = time.Now().Unix()
	}
	return
}

func (record *Base) BeforeUpdate(tx *gorm.DB) (err error) {
	record.UpdatedAt = time.Now().Unix()
	return
}
