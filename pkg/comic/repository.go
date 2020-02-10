package comic

type ComicRepository interface {
	GetOne(id int) (*Comic, error)
	GetAll() ([]*Comic, error)
}
