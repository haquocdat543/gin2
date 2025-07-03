package test

type Service interface {
	CreateUser(
		user *Test,
	) error

	GetAllUsers() (
		[]Test,
		error,
	)
}

type service struct {
	repo Repository
}

func NewService(
	r Repository,
) Service {
	return &service{
		repo: r,
	}
}

func (s *service) CreateUser(
	user *Test,
) error {
	return s.repo.Create(
		user,
	)
}

func (s *service) GetAllUsers() (
	[]Test,
	error,
) {
	return s.repo.FindAll()
}
