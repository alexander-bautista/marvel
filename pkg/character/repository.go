package character

type CharacterRepository interface {
	GetOne(id int) (*Character, error)
	GetAll() ([]*Character, error)
}
