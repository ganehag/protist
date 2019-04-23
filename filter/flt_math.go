package filter

import (
	"encoding/json"
	"errors"
)

/*
 * Const
 */

type ConstParams struct {
	Value float64 `json:"value"`
}

type Const struct {
	Fcore
}

func (s Const) Eval() error {
	return nil
}

func (s *Const) ParseConfig(cfg string) error {
	var args ConstParams
	if err := json.Unmarshal([]byte(cfg), &args); err != nil {
		return err
	}

	fv := new(FloatValue)
	fv.Value = make([]float64, 1)
	fv.Value[0] = args.Value
	s.Retval = fv

	return nil
}

/*
 * Add
 */
type Add struct {
	Fcore
}

func (s *Add) Eval() error {
	if len(s.Source) < 2 {
		return errors.New("less than two sources")
	}

	var r float64
	rv := s.Source[0].GetRetval()
	if rv == nil || rv.IsNull() {
		return errors.New("first source contains nil")

	}
	r = rv.GetFloat()

	for _, item := range s.Source[1:] {
		rv = item.GetRetval()

		if rv == nil || rv.IsNull() {
			return errors.New("consecutive source contains nil")
		} else {
			r += rv.GetFloat()
		}
	}

	fv := new(FloatValue)
	fv.Value = make([]float64, 1)
	fv.Value[0] = r
	s.Retval = fv

	return nil
}

func (s *Add) ParseConfig(cfg string) error {
	return nil
}

/*
 * Sub
 */
type Sub struct {
	Fcore
}

func (s *Sub) Eval() error {
	if len(s.Source) < 2 {
		return errors.New("less than two sources")
	}

	var r float64
	rv := s.Source[0].GetRetval()
	if rv == nil || rv.IsNull() {
		return errors.New("first source contains nil")

	}
	r = rv.GetFloat()

	for _, item := range s.Source[1:] {
		rv = item.GetRetval()

		if rv == nil || rv.IsNull() {
			return errors.New("consecutive source contains nil")
		} else {
			r -= rv.GetFloat()
		}
	}

	fv := new(FloatValue)
	fv.Value = make([]float64, 1)
	fv.Value[0] = r
	s.Retval = fv

	return nil
}

func (s *Sub) ParseConfig(cfg string) error {
	return nil
}

/*
 * Mul
 */
type Mul struct {
	Fcore
}

func (s *Mul) Eval() error {
	if len(s.Source) < 2 {
		return errors.New("less than two sources")
	}

	var r float64
	rv := s.Source[0].GetRetval()
	if rv == nil || rv.IsNull() {
		return errors.New("first source contains nil")

	}
	r = rv.GetFloat()

	for _, item := range s.Source[1:] {
		rv = item.GetRetval()

		if rv == nil || rv.IsNull() {
			return errors.New("consecutive source contains nil")
		} else {
			r *= rv.GetFloat()
		}
	}

	fv := new(FloatValue)
	fv.Value = make([]float64, 1)
	fv.Value[0] = r
	s.Retval = fv

	return nil
}

func (s *Mul) ParseConfig(cfg string) error {
	return nil
}

/*
 * Div
 */
type Div struct {
	Fcore
}

func (s *Div) Eval() error {
	if len(s.Source) < 2 {
		return errors.New("less than two sources")
	}

	var r float64
	rv := s.Source[0].GetRetval()
	if rv == nil || rv.IsNull() {
		return errors.New("first source contains nil")

	}
	r = rv.GetFloat()

	for _, item := range s.Source[1:] {
		rv = item.GetRetval()

		if rv == nil || rv.IsNull() {
			return errors.New("consecutive source contains nil")
		} else {
			r /= rv.GetFloat()
		}
	}

	fv := new(FloatValue)
	fv.Value = make([]float64, 1)
	fv.Value[0] = r
	s.Retval = fv

	return nil
}

func (s *Div) ParseConfig(cfg string) error {
	return nil
}
