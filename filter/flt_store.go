package filter

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

var memstore map[string]float64

func init() {
	memstore = make(map[string]float64)
}

/*
 * Setmem
 */

type SetmemParams struct {
	Name string `json:"name"`
}

type Setmem struct {
	Fcore
	args SetmemParams
}

func (s *Setmem) Eval() error {
	s.Retval = new(FloatValue)

	if len(s.Source) != 1 {
		return errors.New("'setmem' number of sources must be exactly one")
	}

	var r float64
	rv := s.Source[0].GetRetval()
	if rv == nil || rv.IsNull() {
		return errors.New("first source contains nil")

	}
	r = rv.GetFloat()

	fv := new(FloatValue)
	fv.Value = make([]float64, 1)
	fv.Value[0] = r // return input value
	s.Retval = fv

	var key string
	key = fmt.Sprintf("%d.%s", s.Instance, s.args.Name)

	memstore[key] = r

	log.Printf("DEBUG setmem %s %f\n", s.args.Name, memstore[key])

	return nil
}

func (s *Setmem) ParseConfig(cfg string) error {
	if err := json.Unmarshal([]byte(cfg), &s.args); err != nil {
		return err
	}

	return nil
}

/*
 * Getmem
 */

type GetmemParams struct {
	Name string `json:"name"`
}

type Getmem struct {
	Fcore
	args GetmemParams
}

func (s *Getmem) Eval() error {
	s.Retval = new(FloatValue)

	/*
	 * FIXME:
	 * How should we handle a source with NULL or false? Don't get the value?
	 *
	 */

	if val, ok := memstore[fmt.Sprintf("%d.%s", s.Instance, s.args.Name)]; ok {
		fv := new(FloatValue)
		fv.Value = make([]float64, 1)
		fv.Value[0] = val
		s.Retval = fv
	} else {
		return errors.New("getmem unable to fetch value")
	}

	log.Printf("DEBUG getmem %+v\n", s.Retval)

	return nil
}

func (s *Getmem) ParseConfig(cfg string) error {
	if err := json.Unmarshal([]byte(cfg), &s.args); err != nil {
		return err
	}

	return nil
}
