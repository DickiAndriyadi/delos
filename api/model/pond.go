package model

import "time"

type (
	Pond struct {
		ID          uint64     `json:"id" gorm:"primary_key"`
		FarmID      uint64     `json:"farmID" gorm:"column:farm_id"`
		Title       string     `json:"title" gorm:"column:title"`
		Description string     `json:"description" gorm:"column:description"`
		CreatedAt   time.Time  `json:"createdAt" gorm:"column:created_at"`
		UpdatedAt   time.Time  `json:"updatedAt" gorm:"column:updated_at"`
		DeletedAt   *time.Time `json:"deletedAt" gorm:"column:deleted_at"`
	}

	Ponds []Pond
)

func (Pond) ModelName() string {
	return "Pond"
}

func (Pond) TableName() string {
	return "ponds"
}
