package user

type Service interface {
	CreateUser(
		user *User,
	) error

	GetAllUsers() (
		[]User,
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
	user *User,
) error {
	return s.repo.Create(
		user,
	)
}

func (s *service) GetAllUsers() (
	[]User,
	error,
) {
	return s.repo.FindAll()
}
