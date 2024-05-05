package expressions

import (
	"context"
	"testing"
)

func TestAndName(t *testing.T) {
	a := And{}
	got := a.Name()
	if got != "And" {
		t.Fatalf("expected And, got %q", got)
	}
	t.Logf("got %q as expected", got)
}

func TestAndEvaluate(t *testing.T) {

	type test struct {
		operands []Operand
		want     bool
		error    bool
	}

	tests := map[string]test{
		"And(T,T)->T": {
			operands: []Operand{
				True,
				True,
			},
			want: true,
		},
		"And(T,F)->F": {
			operands: []Operand{
				True,
				False,
			},
			want: false,
		},
		"And(F,T)->F": {
			operands: []Operand{
				False,
				True,
			},
			want: false,
		},
		"And(F,F)->F": {
			operands: []Operand{
				False,
				False,
			},
			want: false,
		},
		"And()->?": {
			want:  false,
			error: true,
		},
	}

	for name, run := range tests {
		a := And{}
		a.Add(run.operands...)
		got, err := a.Evaluate(context.Background())
		if !run.error && err != nil {
			t.Fatalf("%s should not fail with err: %v", name, err)
		} else if run.error && err == nil {
			t.Fatalf("%s should have failed", name)
		}
		if got != run.want {
			t.Fatalf("%s: expected %t, got %t", name, run.want, got)
		}
		t.Logf("%s successful", name)
	}
}

func TestOrName(t *testing.T) {
	a := Or{}
	got := a.Name()
	if got != "Or" {
		t.Fatalf("expected Or, got %q", got)
	}
	t.Logf("got %q as expected", got)
}

func TestOrEvaluate(t *testing.T) {

	type test struct {
		operands []Operand
		want     bool
		error    bool
	}

	tests := map[string]test{
		"Or(T,T)->T": {
			operands: []Operand{
				True,
				True,
			},
			want: true,
		},
		"Or(T,F)->T": {
			operands: []Operand{
				True,
				False,
			},
			want: true,
		},
		"Or(F,T)->T": {
			operands: []Operand{
				False,
				True,
			},
			want: true,
		},
		"Or(F,F)->F": {
			operands: []Operand{
				False,
				False,
			},
			want: false,
		},
		"Or()->?": {
			want:  false,
			error: true,
		},
	}

	for name, run := range tests {
		a := Or{}
		a.Add(run.operands...)
		got, err := a.Evaluate(context.Background())
		if !run.error && err != nil {
			t.Fatalf("%s should not fail with err: %v", name, err)
		} else if run.error && err == nil {
			t.Fatalf("%s should have failed", name)
		}
		if got != run.want {
			t.Fatalf("%s: expected %t, got %t", name, run.want, got)
		}
		t.Logf("%s successful", name)
	}
}

func TestNotName(t *testing.T) {
	a := Not{}
	got := a.Name()
	if got != "Not" {
		t.Fatalf("expected Not, got %q", got)
	}
	t.Logf("got %q as expected", got)
}

func TestNotEvaluate(t *testing.T) {

	type test struct {
		operands []Operand
		want     bool
		error    bool
	}

	tests := map[string]test{
		"Not(T)->F": {
			operands: []Operand{
				True,
			},
			want: false,
		},
		"Not(F)->T": {
			operands: []Operand{
				False,
			},
			want: true,
		},
		"Not(F,T)->?": {
			operands: []Operand{
				False,
				True,
			},
			want:  false,
			error: true,
		},
		"Not()->?": {
			operands: []Operand{},
			want:     false,
			error:    true,
		},
	}

	for name, run := range tests {
		a := Not{}
		a.Add(run.operands...)
		got, err := a.Evaluate(context.Background())
		if !run.error && err != nil {
			t.Fatalf("%s should not fail with err: %v", name, err)
		} else if run.error && err == nil {
			t.Fatalf("%s should have failed", name)
		}
		if got != run.want {
			t.Fatalf("%s: expected %t, got %t", name, run.want, got)
		}
		t.Logf("%s successful", name)
	}
}

func TestXorName(t *testing.T) {
	a := Xor{}
	got := a.Name()
	if got != "Xor" {
		t.Fatalf("expected Or, got %q", got)
	}
	t.Logf("got %q as expected", got)
}

func TestXorEvaluate(t *testing.T) {

	type test struct {
		operands []Operand
		want     bool
		error    bool
	}

	tests := map[string]test{
		"Xor(T,T)->T": {
			operands: []Operand{
				True,
				True,
			},
			want: false,
		},
		"Xor(T,F)->T": {
			operands: []Operand{
				True,
				False,
			},
			want: true,
		},
		"Xor(F,T)->T": {
			operands: []Operand{
				False,
				True,
			},
			want: true,
		},
		"Xor(F,F)->F": {
			operands: []Operand{
				False,
				False,
			},
			want: false,
		},
		"Xor(T,T,T)->?": {
			operands: []Operand{
				True,
				True,
				True,
			},
			want:  false,
			error: true,
		},
		"Xor()->?": {
			want:  false,
			error: true,
		},
	}

	for name, run := range tests {
		a := Xor{}
		a.Add(run.operands...)
		got, err := a.Evaluate(context.Background())
		if !run.error && err != nil {
			t.Fatalf("%s should not fail with err: %v", name, err)
		} else if run.error && err == nil {
			t.Fatalf("%s should have failed", name)
		}
		if got != run.want {
			t.Fatalf("%s: expected %t, got %t", name, run.want, got)
		}
		t.Logf("%s successful", name)
	}
}
