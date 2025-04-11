package request

import (
	"context"
)

type Register struct {
	Email    string `validate:"required,email" mod:"trim,lcase"`
	FullName string `validate:"required,lte=255" mod:"trim"`
	Password string `validate:"required" mod:"trim"`
}

func (l *Register) Validate() error {
	if err := transformer.Struct(context.Background(), l); err != nil {
		return err
	}

	if err := validator.Struct(l); err != nil {
		return err
	}

	return nil
}
