package usecase

type UserUsecase struct {
	repository UserRepository
}

func NewUserUsecase() {
	return &UserUsecase{}
}
