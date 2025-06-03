package user

type UserService interface {
	Register(user User) (User, error)
}

type GormUserService struct {
	repo UserRepo
}

func NewGormUserService(repo UserRepo) *GormUserService {
	return &GormUserService{repo: repo}
}

func (s *GormUserService) Register(newUser User) (User, error) {
	return s.repo.Create(newUser)
}
