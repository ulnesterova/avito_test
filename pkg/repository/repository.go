package repository

import (
	"github.com/jmoiron/sqlx"
	avitotest "github.com/ulnesterova/avito_test"
)

type User interface {
	CreateUser(user avitotest.User) (int, error)
	GetUser(username, password string) (avitotest.User, error)
}

type Segment interface {
	Create(segment avitotest.Segment) (int, error)
	Delete(slug string) error
	GetAll(userId int) ([]string, error)
	AddSegmentToUser(userId int, slug string) error
	DeleteSegmentFromUser(userId int, slug string) error
}

type Repository struct {
	User
	Segment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:    NewAuthPostgres(db),
		Segment: NewSegmentPostgres(db),
	}
}
