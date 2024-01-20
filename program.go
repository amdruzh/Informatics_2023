package program

import (
	"fmt"
)

type Program struct {
	Name          string
	YearOfRelease int
	Type          string
}

func (p *Program) GetProgramType() string {
	if p.Name == "Microsoft Office" {
		return "Microsoft"
	} else if p.Name == "Google Chrome" {
		return "Google"
	} else if p.Name == "Apple iTunes" {
		return "Apple"
	} else {
		return "Unknown"
	}
}

func (p *Program) String() string {
	return fmt.Sprintf("%s (%s, %d)", p.Name, p.Type, p.YearOfRelease)
}
