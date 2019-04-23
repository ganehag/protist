package filter

import (
	"testing"
)

/*
 * Const
 *
 */

func TestOrParseValidConfig(t *testing.T) {
	f := Or{}
	if err := f.ParseConfig(``); err != nil {
		t.Error("Argument not ignored")
	}

	if err := f.ParseConfig(`InValidJSON`); err != nil {
		t.Error("Argument not ignored")
	}
}

func TestOrEval(t *testing.T) {
	f := Or{}

	if f.Eval() == nil {
		t.Error("Eval didn't fail on missing sources")
	}

	c1 := new(Const) // Nil source
	c2 := new(Const) // Nil source

	// Add only one source
	f.AddSource(c1)

	// Requires 2..X sources
	if f.Eval() == nil {
		t.Error("Eval didn't fail on missing sources")
	}

	// Add second source
	f.AddSource(c2)

	// Requires 2..X sources
	if f.Eval() != nil {
		t.Error("Eval failed to execute")
	}

	if f.GetRetval().IsNull() == false {
		t.Error("Eval failed basic logic")
	}

	c3 := new(Const)
	c3.ParseConfig(`{"value": 3}`)

	// Add source with value
	f.AddSource(c3)

	if f.Eval() != nil {
		t.Error("Eval failed to execute")
	}

	if f.GetRetval().GetFloat() != 3 {
		t.Error("Eval failed basic logic")
	}
}
