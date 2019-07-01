package service

import (
	"github.com/villers/api/datamodel"
	"github.com/villers/api/repository"
)

type UserService interface {
	GetByID(id int64) (datamodel.User, bool)
	GetAll() map[int64]datamodel.User
	InsertOrUpdate(user datamodel.User) (datamodel.User, error)
	DeleteByID(id int64) bool
}

// NewMovieService returns the default movie service.
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

type userService struct {
	repo repository.UserRepository
}

func (s *userService) GetByID(id int64) (datamodel.User, bool) {
	return s.repo.GetByID(id)
}

func (s *userService) GetAll() map[int64]datamodel.User {
	return s.repo.GetAll()
}

func (s *userService) InsertOrUpdate(user datamodel.User) (datamodel.User, error) {
	return s.repo.InsertOrUpdate(user)
}

func (s *userService) DeleteByID(id int64) bool {
	return s.repo.DeleteByID(id)
}
