package user

import "golang.org/x/crypto/bcrypt"

type Service struct {
	repo Repo
}

func NewService(repo Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(req CreateUserRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	// FIXME: Error?
	if err != nil {
		return err
	}

	user := NewUser(NewUserOpts{
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: hashedPassword,
	})

	err = s.repo.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetByUsername(username string) (*User, error) {
	return s.repo.GetByUsername(username)
}
