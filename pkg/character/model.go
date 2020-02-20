package character

type Character struct {
	Id          int
	Name        string
	Description string
	Year        int
}

func (c *Character) Scream(what string) string {
	return "!!!! Scream " + what
}
