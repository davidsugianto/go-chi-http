package user

type userRepo interface {
}

type UseCase struct {
	userRepo userRepo
}

func New(user userRepo) *UseCase {
	return &UseCase{
		userRepo: user,
	}
}
