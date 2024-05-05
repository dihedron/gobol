package expressions

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
)

var (
	ErrInvalidOperands = errors.New("invalid number of operands")
)

type Operand interface {
	Name() string
	Evaluate(ctx context.Context) (bool, error)
}

var (
	True  = boolean{true}
	False = boolean{false}
)

type boolean struct {
	value bool
}

func (b boolean) Name() string {
	if b.value {
		return "True"
	}
	return "False"
}

func (b boolean) Evaluate(ctx context.Context) (bool, error) {
	return b.value, nil
}

type Operator struct {
	Operand
	lock     sync.RWMutex
	Operands []Operand
}

func (o *Operator) Add(operands ...Operand) error {
	o.lock.Lock()
	defer o.lock.Unlock()
	if o.Operands == nil {
		o.Operands = operands
	} else {
		o.Operands = append(o.Operands, operands...)
	}
	return nil
}

func (o *Operator) String() string {
	o.lock.RLock()
	defer o.lock.RUnlock()
	names := []string{}
	for _, operand := range o.Operands {
		names = append(names, operand.Name())
	}
	return fmt.Sprintf("%s(%s)", o.Name(), strings.Join(names, ", "))
}

type And struct {
	Operator
}

func (*And) Name() string {
	return "And"
}

func (a *And) Evaluate(ctx context.Context) (bool, error) {
	a.lock.RLock()
	defer a.lock.RUnlock()
	if len(a.Operands) == 0 {
		return false, ErrInvalidOperands
	}
	for _, operand := range a.Operands {
		result, err := operand.Evaluate(ctx)
		if err != nil {
			return false, err
		}
		if !result {
			return false, nil
		}
	}
	return true, nil
}

type Or struct {
	Operator
}

func (*Or) Name() string {
	return "Or"
}

func (o *Or) Evaluate(ctx context.Context) (bool, error) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	if len(o.Operands) == 0 {
		return false, ErrInvalidOperands
	}
	for _, operand := range o.Operands {
		result, err := operand.Evaluate(ctx)
		if err != nil {
			return false, err
		}
		if result {
			return true, nil
		}
	}
	return false, nil
}

type Not struct {
	Operator
}

func (*Not) Name() string {
	return "Not"
}

func (n *Not) Evaluate(ctx context.Context) (bool, error) {
	n.lock.RLock()
	defer n.lock.RUnlock()
	if len(n.Operands) != 1 {
		return false, ErrInvalidOperands
	}
	result, err := n.Operands[0].Evaluate(ctx)
	if err != nil {
		return false, err
	}
	return !result, nil
}

type Xor struct {
	Operator
}

func (*Xor) Name() string {
	return "Xor"
}

func (x *Xor) Evaluate(ctx context.Context) (bool, error) {
	x.lock.RLock()
	defer x.lock.RUnlock()
	if len(x.Operands) != 2 {
		return false, ErrInvalidOperands
	}
	r0, err := x.Operands[0].Evaluate(ctx)
	if err != nil {
		return false, err
	}
	r1, err := x.Operands[1].Evaluate(ctx)
	if err != nil {
		return false, err
	}
	return r0 != r1, nil
}
