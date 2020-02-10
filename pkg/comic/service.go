package comic

import (
	"github.com/pkg/errors"
)

type ComicService interface {
	GetOne(id int) (*Comic, error)
	GetAll() ([]*Comic, error)
	CalculateTaxes(id int) (float32, error)
}

var (
	ErrComicNotFound = errors.New("Comic Not Found")
)

type service struct {
	r ComicRepository
}

func NewComicService(r ComicRepository) ComicService {
	return &service{r}
}

func (s *service) GetOne(id int) (*Comic, error) {
	comic, err := s.r.GetOne(id)

	if err != nil {
		return nil, errors.Wrap(ErrComicNotFound, "service.Comic.GetOne")
	}

	return comic, nil

}

func (s *service) GetAll() ([]*Comic, error) {
	return s.r.GetAll()
}

func (s *service) CalculateTaxes(id int) (float32, error) {
	comic, err := s.r.GetOne(id)

	if err != nil {
		return 0, errors.Wrap(ErrComicNotFound, "Service.Comic.CalculateTaxes")
	}

	return comic.EstimatedTaxes(), nil
}
