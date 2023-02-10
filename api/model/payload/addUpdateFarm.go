package payload

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AddUpdateFarm struct {
	Title       string     `json:"title" validate:"required"`
	Description string     `json:"description"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func (a AddUpdateFarm) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Title, validation.Required),
	)
}
