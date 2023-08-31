package service

import (
	avitotest "github.com/ulnesterova/avito_test"
	"github.com/ulnesterova/avito_test/pkg/repository"
)

type SegmentService struct {
	repo repository.Segment
}

func NewSegmentService(repo repository.Segment) *SegmentService {
	return &SegmentService{repo: repo}
}

func (s *SegmentService) Create(segment avitotest.Segment) (int, error) {
	return s.repo.Create(segment)
}

func (s *SegmentService) Delete(slug string) error {
	return s.repo.Delete(slug)
}

func (s *SegmentService) GetAll(userId int) ([]string, error) {
	return s.repo.GetAll(userId)
}

func (s *SegmentService) AddSegmentToUser(userId int, slug string) error {
	return s.repo.AddSegmentToUser(userId, slug)
}

func (s *SegmentService) DeleteSegmentFromUser(userId int, slug string) error {
	return s.repo.DeleteSegmentFromUser(userId, slug)
}
