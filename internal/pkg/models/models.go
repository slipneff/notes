package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	Id        uuid.UUID  `gorm:"primaryKey;type:uuid"`
	Header    string     `gorm:"column:header"`
	Content   string     `gorm:"column:content"`
	CreatedAt *time.Time `gorm:"column:CREATED_AT;autoCreateTime"`
}

func (p *Note) BeforeCreate(tx *gorm.DB) (err error) {
	if p.Id == uuid.Nil {
		p.Id = uuid.New()
	}

	return
}
