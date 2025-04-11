package request

import (
	"context"
)

type Register struct {
	Email    string `validate:"required,email" mod:"trim,lcase"`
	FullName string `validate:"required,lte=255" mod:"trim"`
	Password string `validate:"required" mod:"trim"`
}

func (r *Register) Validate() error {
	if err := transformer.Struct(context.Background(), r); err != nil {
		return err
	}

	return validator.Struct(r)
}
