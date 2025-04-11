package request

import "context"

type UpdateVideo struct {
	Id        int64 `validate:"required"`
	Name      string
	Likes     int64
	Views     int64
	Comments  int64
	Shares    int64
	Length    int64
	WatchTime int64
}

func (r *UpdateVideo) Validate() error {
	if err := transformer.Struct(context.Background(), r); err != nil {
		return err
	}

	return validator.Struct(r)
}
