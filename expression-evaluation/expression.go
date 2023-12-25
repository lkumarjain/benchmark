package expressionevaluation

import (
	"errors"

	"github.com/google/cel-go/interpreter"
)

type Input struct {
	ID       int
	Name     string
	City     string
	Country  string
	Currency string
}

func (i Input) ResolveName(name string) (any, bool) {
	switch name {
	case "ID":
		return i.ID, true
	case "Name":
		return i.Name, true
	case "City":
		return i.City, true
	case "Country":
		return i.Country, true
	case "Currency":
		return i.Currency, true
	default:
		return name, false
	}
}

func (i Input) Parent() interpreter.Activation {
	return nil
}

func (i Input) Get(name string) (interface{}, error) {
	value, found := i.ResolveName(name)

	if !found {
		return nil, errors.New("Error")
	}

	return value, nil
}
