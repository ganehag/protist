package filter

import (
	"errors"
	"fmt"
)

type ReturnValue interface {
	GetFloat() float64
	GetFloats() []float64
	GetBool() bool
	GetBools() []bool
	GetInt() int64
	GetInts() []int64
	IsNull() bool
}

type Filter interface {
	AddSource(Filter) error
	GetRetval() ReturnValue
	Eval() error
	ParseConfig(string) error

	SetInstance(int64)
}

type FloatValue struct {
	Value []float64
}

func (f *FloatValue) IsNull() bool {
	if len(f.Value) > 0 {
		return false
	}
	return true
}

func (f *FloatValue) GetFloat() float64 {
	if f.IsNull() == false {
		return f.Value[0]
	}
	return 0
}

func (f *FloatValue) GetFloats() []float64 {
	if f.IsNull() == false {
		return f.Value
	}
	return make([]float64, 0)
}

func (f *FloatValue) GetBool() bool {
	if f.IsNull() == true {
		return false
	} else if f.Value[0] != 1 {
		return false
	}
	return true
}

func (f *FloatValue) GetBools() []bool {
	if f.IsNull() == true {
		return make([]bool, 0)
	}
	b := make([]bool, len(f.Value))
	for i, v := range f.Value {
		if v == 1 {
			b[i] = true
		} else {
			b[i] = false
		}
	}
	return b
}

func (f *FloatValue) GetInt() int64 {
	if f.IsNull() == false {
		return int64(f.Value[0])
	}
	return 0
}

func (f *FloatValue) GetInts() []int64 {
	if f.IsNull() == true {
		return make([]int64, 0)
	}
	ivals := make([]int64, len(f.Value))
	for i, v := range f.Value {
		ivals[i] = int64(v)
	}
	return ivals
}

type Fcore struct {
	Instance int64
	Source   []Filter
	Retval   ReturnValue
}

func (c *Fcore) AddSource(f Filter) error {
	if f == nil {
		return errors.New("supplied argument is nil")

	}
	c.Source = append(c.Source, f)

	return nil
}

func (c *Fcore) GetRetval() ReturnValue {
	return c.Retval
}

func (c *Fcore) PrintSources() {
	for _, item := range c.Source {
		fmt.Printf("%+v\n", item)
	}
}
func (c *Fcore) SetInstance(i int64) {
	c.Instance = i
}

func Factory(instance int64, t string, cfg string) (Filter, error) {
	var f Filter

	switch t {
	case "const":
		f = new(Const)
		break
	case "add":
		f = new(Add)
		break
	case "sub":
		f = new(Sub)
		break
	case "mul":
		f = new(Mul)
		break
	case "div":
		f = new(Div)
		break

	case "or":
		f = new(Or)
		break

	case "setmem":
		f = new(Setmem)
		break
	case "getmem":
		f = new(Getmem)
		break
	default:
		return nil, errors.New("invalid type")
	}

	if f == nil {
		return nil, errors.New("unable to create filter")
	}

	f.SetInstance(instance)

	if err := f.ParseConfig(cfg); err != nil {
		return nil, err
	}

	return f, nil
}
