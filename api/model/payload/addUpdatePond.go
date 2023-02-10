package payload

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AddUpdatePond struct {
	Title       string     `json:"title" validate:"required"`
	FarmID      uint64     `json:"farm_id"`
	Description string     `json:"description"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Req         string
}

func (a AddUpdatePond) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Title, validation.Required),
		validation.Field(&a.FarmID, validation.When(a.Req == "Add", validation.Required)),
	)
}
