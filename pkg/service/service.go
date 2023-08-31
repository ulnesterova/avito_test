package service

import (
	avitotest "github.com/ulnesterova/avito_test"
	"github.com/ulnesterova/avito_test/pkg/repository"
)

type User interface {
	CreateUser(user avitotest.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Segment interface {
	Create(segment avitotest.Segment) (int, error)
	Delete(slug string) error
	GetAll(userId int) ([]string, error)
	AddSegmentToUser(userId int, slug string) error
	DeleteSegmentFromUser(userId int, slug string) error
}

type Service struct {
	User
	Segment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:    NewAuthService(repos.User),
		Segment: NewSegmentService(repos.Segment),
	}
}
