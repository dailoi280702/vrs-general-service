package auth

type Usecase struct{}

func NewUsecasez() Interface {
	return &Usecase{}
}
