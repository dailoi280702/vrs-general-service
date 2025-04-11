package request

import "context"

type GetById struct {
	Id int64 `validate:"required,gt=0"`
}

func (r *GetById) Validate() error {
	if err := transformer.Struct(context.Background(), r); err != nil {
		return err
	}

	return validator.Struct(r)
}

type GetByIds struct {
	Ids []int64
}
