package character

import (
	"github.com/pkg/errors"
)

type CharacterService interface {
	GetOne(id int) (*Character, error)
	GetAll() ([]*Character, error)
	Scream(int, string) (string, error)
	Add(*Character) (interface{}, error)
}

var (
	ErrCharacterNotFound = errors.New("Character Not Found")
)

type service struct {
	r CharacterRepository
}

func NewCharacterService(r CharacterRepository) CharacterService {
	return &service{r}
}

func (s *service) GetOne(id int) (*Character, error) {
	char, err := s.r.GetOne(id)

	if err != nil {
		return nil, errors.Wrap(ErrCharacterNotFound, "service.Character.GetOne")
	}

	return char, nil
}

func (s *service) GetAll() ([]*Character, error) {
	return s.r.GetAll()
}

func (s *service) Scream(id int, what string) (string, error) {
	char, err := s.r.GetOne(id)

	if err != nil {
		return "", errors.Wrap(ErrCharacterNotFound, "service.Character.Scream")
	}

	scream := char.Scream(what)

	return scream, nil
}

func (s *service) Add(char *Character) (interface{}, error) {
	result, err := s.r.Add(char)

	if err != nil {
		return nil, err
	}

	return result, err

}
