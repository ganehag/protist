package filter

import (
	"errors"
)

/*
 * Or
 */

type OrParams struct {
	Name string `json:"name"`
}

type Or struct {
	Fcore
}

func (s *Or) Eval() error {
	s.Retval = new(FloatValue)

	if len(s.Source) < 2 {
		return errors.New("less than two sources")
	}

	for _, item := range s.Source {
		rv := item.GetRetval()

		if rv == nil || rv.IsNull() {
			continue
		}
		s.Retval = rv
		break
	}

	return nil
}

func (s *Or) ParseConfig(cfg string) error {
	return nil
}
