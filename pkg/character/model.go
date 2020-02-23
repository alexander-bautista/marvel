package character

import "fmt"

type Character struct {
	Id          int
	Name        string
	Description string
	Year        int
}

func (c *Character) Scream(what string) string {
	return fmt.Sprintf("%s scream %s", c.Name, what)
}
