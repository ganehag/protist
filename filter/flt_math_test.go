package filter

import (
	"testing"
)

/*
 * Const
 *
 */

func TestConstParseValidConfig(t *testing.T) {
	c := Const{}
	if err := c.ParseConfig(`{"value": 3.14}`); err != nil {
		t.Error("Unable to parse valid config")
	}
}

func TestConstParseInValidConfig(t *testing.T) {
	c := Const{}

	if err := c.ParseConfig(`{"value": "string_not_float"}`); err == nil {
		t.Error("Didn't fail on invalid config")
	}

	if err := c.ParseConfig(`{"value":}`); err == nil {
		t.Error("Didn't fail on invalid config")
	}

	if err := c.ParseConfig(`{"val`); err == nil {
		t.Error("Didn't fail on invalid config")
	}

	if err := c.ParseConfig(``); err == nil {
		t.Error("Didn't fail on invalid config")
	}
}

func TestConstEval(t *testing.T) {
	c := Const{}

	if c.Eval() != nil {
		t.Error("Eval did not return nil")
	}
}

/*
 * Add
 *
 */

func TestAddParseConfig(t *testing.T) {
	f := Add{}
	if err := f.ParseConfig(``); err != nil {
		t.Error("Argument not ignored")
	}

	if err := f.ParseConfig(`InValidJSON`); err != nil {
		t.Error("Argument not ignored")
	}
}

func TestAddEval(t *testing.T) {
	f := Add{}

	if f.Eval() == nil {
		t.Error("Eval didn't fail on missing sources")
	}

	c1 := new(Const)
	c1.ParseConfig(`{"value": 1}`)

	// Add only one source
	f.AddSource(c1)

	// Requires 2..X sources
	if f.Eval() == nil {
		t.Error("Eval didn't fail on missing sources")
	}

	c2 := new(Const)
	c2.ParseConfig(`{"value": 2}`)

	// Add second source
	f.AddSource(c2)

	// Requires 2..X sources
	if f.Eval() != nil {
		t.Error("Eval failed to execute")
	}

	if f.GetRetval().GetFloat() != 3 {
		t.Error("Eval failed basic math")
	}

	c3 := new(Const)
	c3.ParseConfig(`{"value": 3}`)

	// Testing three sources
	f.AddSource(c3)

	if f.Eval() != nil {
		t.Error("Eval failed to execute")
	}

	if f.GetRetval().GetFloat() != 6 {
		t.Error("Eval failed basic math")
	}
}

/*
 * Sub
 *
 */

func TestSubParseConfig(t *testing.T) {
	f := Sub{}
	if err := f.ParseConfig(``); err != nil {
		t.Error("Argument not ignored")
	}

	if err := f.ParseConfig(`InValidJSON`); err != nil {
		t.Error("Argument not ignored")
	}
}

func TestSubEval(t *testing.T) {
	f := Sub{}

	if f.Eval() == nil {
		t.Error("Eval didn't fail on missing sources")
	}

	c1 := new(Const)
	c1.ParseConfig(`{"value": 1}`)

	// Add only one source
	f.AddSource(c1)

	// Requires 2..X sources
	if f.Eval() == nil {
		t.Error("Eval didn't fail on missing sources")
	}

	c2 := new(Const)
	c2.ParseConfig(`{"value": 2}`)

	// Add second source
	f.AddSource(c2)

	// Requires 2..X sources
	if f.Eval() != nil {
		t.Error("Eval failed to execute")
	}

	if f.GetRetval().GetFloat() != -1 {
		t.Error("Eval failed basic math")
	}

	c3 := new(Const)
	c3.ParseConfig(`{"value": 3}`)

	// Testing three sources
	f.AddSource(c3)

	if f.Eval() != nil {
		t.Error("Eval failed to execute")
	}

	if f.GetRetval().GetFloat() != -4 {
		t.Error("Eval failed basic math")
	}
}

/*
 * Mul
 *
 */

func TestMulParseConfig(t *testing.T) {
	f := Mul{}
	if err := f.ParseConfig(``); err != nil {
		t.Error("Argument not ignored")
	}

	if err := f.ParseConfig(`InValidJSON`); err != nil {
		t.Error("Argument not ignored")
	}
}

func TestMulEval(t *testing.T) {
	f := Mul{}

	if f.Eval() == nil {
		t.Error("Eval didn't fail on missing sources")
	}

	c1 := new(Const)
	c1.ParseConfig(`{"value": 1}`)

	// Add only one source
	f.AddSource(c1)

	// Requires 2..X sources
	if f.Eval() == nil {
		t.Error("Eval didn't fail on missing sources")
	}

	c2 := new(Const)
	c2.ParseConfig(`{"value": 2}`)

	// Add second source
	f.AddSource(c2)

	// Requires 2..X sources
	if f.Eval() != nil {
		t.Error("Eval failed to execute")
	}

	if f.GetRetval().GetFloat() != 2 {
		t.Error("Eval failed basic math")
	}

	c3 := new(Const)
	c3.ParseConfig(`{"value": 3}`)

	// Testing three sources
	f.AddSource(c3)

	if f.Eval() != nil {
		t.Error("Eval failed to execute")
	}

	if f.GetRetval().GetFloat() != 6 {
		t.Error("Eval failed basic math")
	}
}

/*
 * Div
 *
 */

func TestDivParseConfig(t *testing.T) {
	f := Div{}
	if err := f.ParseConfig(``); err != nil {
		t.Error("Argument not ignored")
	}

	if err := f.ParseConfig(`InValidJSON`); err != nil {
		t.Error("Argument not ignored")
	}
}

func TestDivEval(t *testing.T) {
	f := Div{}

	if f.Eval() == nil {
		t.Error("Eval didn't fail on missing sources")
	}

	c1 := new(Const)
	c1.ParseConfig(`{"value": 180}`)

	// Add only one source
	f.AddSource(c1)

	// Requires 2..X sources
	if f.Eval() == nil {
		t.Error("Eval didn't fail on missing sources")
	}

	c2 := new(Const)
	c2.ParseConfig(`{"value": 2}`)

	// Add second source
	f.AddSource(c2)

	// Requires 2..X sources
	if f.Eval() != nil {
		t.Error("Eval failed to execute")
	}

	if f.GetRetval().GetFloat() != 90 {
		t.Error("Eval failed basic math")
	}

	c3 := new(Const)
	c3.ParseConfig(`{"value": 3}`)

	// Testing three sources
	f.AddSource(c3)

	if f.Eval() != nil {
		t.Error("Eval failed to execute")
	}

	if f.GetRetval().GetFloat() != 30 {
		t.Error("Eval failed basic math")
	}
}
