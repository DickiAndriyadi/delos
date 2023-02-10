package model

import "time"

type (
	Farm struct {
		ID          uint64     `json:"id" gorm:"primary_key"`
		Title       string     `json:"title" gorm:"column:title"`
		Description string     `json:"description" gorm:"column:description"`
		CreatedAt   time.Time  `json:"createdAt" gorm:"column:created_at"`
		UpdatedAt   time.Time  `json:"updatedAt" gorm:"column:updated_at"`
		DeletedAt   *time.Time `json:"deletedAt" gorm:"column:deleted_at"`
		Ponds       Ponds      `json:"ponds,omitempty"`
	}

	Farms []Farm
)

func (Farm) ModelName() string {
	return "Farm"
}

func (Farm) TableName() string {
	return "farms"
}
