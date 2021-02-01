package vm

import (
	"fmt"

	"github.com/anywhereQL/anywhereQL/common/result"
	"github.com/anywhereQL/anywhereQL/common/value"
	"github.com/anywhereQL/anywhereQL/runtime/vm/function"
)

type OpeType int

const (
	_ OpeType = iota
	PUSH
	POP
	ADD
	SUB
	MUL
	DIV
	MOD
	STORE
	CALL
)

func (o OpeType) String() string {
	switch o {
	case PUSH:
		return "PUSH"
	case POP:
		return "POP"
	case ADD:
		return "ADD"
	case SUB:
		return "SUB"
	case MUL:
		return "MUL"
	case DIV:
		return "DIV"
	case MOD:
		return "MOD"
	case CALL:
		return "CALL"
	case STORE:
		return "STORE"
	default:
		return "Unknwon Operation"
	}
}

type VMCode struct {
	Operator OpeType
	Operand1 value.Value
	Operand2 value.Value
}

func (c VMCode) String() string {
	s := ""
	s = fmt.Sprintf("%s", c.Operator)

	if c.Operand1.Type != value.NA {
		switch c.Operand1.Type {
		case value.INTEGER:
			s = fmt.Sprintf("%s %d", s, c.Operand1.Int)
		case value.FLOAT:
			s = fmt.Sprintf("%s %f", s, c.Operand1.Float)
		case value.STRING:
			s = fmt.Sprintf("%s %s", s, c.Operand1.String)
		}
	}

	if c.Operand2.Type != value.NA {
		switch c.Operand2.Type {
		case value.INTEGER:
			s = fmt.Sprintf("%s %d", s, c.Operand2.Int)
		case value.FLOAT:
			s = fmt.Sprintf("%s %f", s, c.Operand2.Float)
		case value.STRING:
			s = fmt.Sprintf("%s %s", s, c.Operand2.String)
		}
	}

	return s
}

func Run(codes []VMCode) ([]result.Value, error) {
	s := newStack()
	cols := []result.Value{}

	for _, code := range codes {
		switch code.Operator {
		case PUSH:
			s.push(code.Operand1)
		case ADD:
			ope2, err := s.pop()
			if err != nil {
				return []result.Value{}, err
			}
			ope1, err := s.pop()
			if err != nil {
				return []result.Value{}, err
			}
			if ope1.Type == value.INTEGER && ope2.Type == value.INTEGER {
				v := value.Value{
					Type: value.INTEGER,
					Int:  ope1.Int + ope2.Int,
				}
				s.push(v)
			} else if ope1.Type == value.FLOAT && ope2.Type == value.FLOAT {
				v := value.Value{
					Type:  value.FLOAT,
					Float: ope1.Float + ope2.Float,
				}
				s.push(v)
			} else if ope1.Type == value.FLOAT && ope2.Type == value.INTEGER {
				v := value.Value{
					Type:  value.FLOAT,
					Float: ope1.Float + float64(ope2.Int),
				}
				s.push(v)
			} else if ope1.Type == value.INTEGER && ope2.Type == value.FLOAT {
				v := value.Value{
					Type:  value.FLOAT,
					Float: float64(ope1.Int) + ope2.Float,
				}
				s.push(v)
			} else {
				return []result.Value{}, fmt.Errorf("Unknown Operation: %s + %s", ope1.Type, ope2.Type)
			}

		case SUB:
			ope2, err := s.pop()
			if err != nil {
				return []result.Value{}, err
			}
			ope1, err := s.pop()
			if err != nil {
				return []result.Value{}, err
			}
			if ope1.Type == value.INTEGER && ope2.Type == value.INTEGER {
				v := value.Value{
					Type: value.INTEGER,
					Int:  ope1.Int - ope2.Int,
				}
				s.push(v)
			} else if ope1.Type == value.FLOAT && ope2.Type == value.FLOAT {
				v := value.Value{
					Type:  value.FLOAT,
					Float: ope1.Float - ope2.Float,
				}
				s.push(v)
			} else if ope1.Type == value.FLOAT && ope2.Type == value.INTEGER {
				v := value.Value{
					Type:  value.FLOAT,
					Float: ope1.Float - float64(ope2.Int),
				}
				s.push(v)
			} else if ope1.Type == value.INTEGER && ope2.Type == value.FLOAT {
				v := value.Value{
					Type:  value.FLOAT,
					Float: float64(ope1.Int) - ope2.Float,
				}
				s.push(v)
			} else {
				return []result.Value{}, fmt.Errorf("Unknown Operation: %s - %s", ope1.Type, ope2.Type)
			}

		case MUL:
			ope2, err := s.pop()
			if err != nil {
				return []result.Value{}, err
			}
			ope1, err := s.pop()
			if err != nil {
				return []result.Value{}, err
			}
			if ope1.Type == value.INTEGER && ope2.Type == value.INTEGER {
				v := value.Value{
					Type: value.INTEGER,
					Int:  ope1.Int * ope2.Int,
				}
				s.push(v)
			} else if ope1.Type == value.FLOAT && ope2.Type == value.FLOAT {
				v := value.Value{
					Type:  value.FLOAT,
					Float: ope1.Float * ope2.Float,
				}
				s.push(v)
			} else if ope1.Type == value.FLOAT && ope2.Type == value.INTEGER {
				v := value.Value{
					Type:  value.FLOAT,
					Float: ope1.Float * float64(ope2.Int),
				}
				s.push(v)
			} else if ope1.Type == value.INTEGER && ope2.Type == value.FLOAT {
				v := value.Value{
					Type:  value.FLOAT,
					Float: float64(ope1.Int) * ope2.Float,
				}
				s.push(v)
			} else {
				return []result.Value{}, fmt.Errorf("Unknown Operation: %s * %s", ope1.Type, ope2.Type)
			}

		case DIV:
			ope2, err := s.pop()
			if err != nil {
				return []result.Value{}, err
			}
			ope1, err := s.pop()
			if err != nil {
				return []result.Value{}, err
			}
			if ope1.Type == value.INTEGER && ope2.Type == value.INTEGER {
				if ope2.Int == 0 {
					return []result.Value{}, fmt.Errorf("Div by 0")
				}
				v := value.Value{
					Type: value.INTEGER,
					Int:  ope1.Int / ope2.Int,
				}
				s.push(v)
			} else if ope1.Type == value.FLOAT && ope2.Type == value.FLOAT {
				if ope2.Float == 0 {
					return []result.Value{}, fmt.Errorf("Div by 0")
				}
				v := value.Value{
					Type:  value.FLOAT,
					Float: ope1.Float / ope2.Float,
				}
				s.push(v)
			} else if ope1.Type == value.FLOAT && ope2.Type == value.INTEGER {
				if ope2.Int == 0 {
					return []result.Value{}, fmt.Errorf("Div by 0")
				}
				v := value.Value{
					Type:  value.FLOAT,
					Float: ope1.Float / float64(ope2.Int),
				}
				s.push(v)
			} else if ope1.Type == value.INTEGER && ope2.Type == value.FLOAT {
				if ope2.Float == 0 {
					return []result.Value{}, fmt.Errorf("Div by 0")
				}
				v := value.Value{
					Type:  value.FLOAT,
					Float: float64(ope1.Int) / ope2.Float,
				}
				s.push(v)
			} else {
				return []result.Value{}, fmt.Errorf("Unknown Operation: %s / %s", ope1.Type, ope2.Type)
			}

		case MOD:
			ope2, err := s.pop()
			if err != nil {
				return []result.Value{}, err
			}
			if ope2.Int == 0 {
				return []result.Value{}, fmt.Errorf("Div by 0")
			}
			ope1, err := s.pop()
			if err != nil {
				return []result.Value{}, err
			}
			if ope1.Type == value.INTEGER && ope2.Type == value.INTEGER {
				v := value.Value{
					Type: value.INTEGER,
					Int:  ope1.Int % ope2.Int,
				}
				s.push(v)
			} else {
				return []result.Value{}, fmt.Errorf("Unknown Operation: %s %% %s", ope1.Type, ope2.Type)
			}
		case CALL:
			args := []value.Value{}

			argsN, err := s.pop()
			if err != nil {
				return []result.Value{}, err
			}
			for i := 0; int64(i) < argsN.Int; i++ {
				v, err := s.pop()
				if err != nil {
					return []result.Value{}, err
				}
				args = append(args, v)
			}

			call := function.LookupFunction(code.Operand1.String)
			if call == nil {
				return []result.Value{}, fmt.Errorf("Function(%s) is not implement", code.Operand1.String)
			}
			r, err := call(args)
			if err != nil {
				return []result.Value{}, err
			}
			var vr value.Value
			switch r.Type {
			case result.Integral:
				vr = value.Value{
					Type: value.INTEGER,
					Int:  r.Integral,
				}
			case result.Float:
				vr = value.Value{
					Type:   value.FLOAT,
					Float:  r.Float,
					PartI:  r.PartI,
					PartF:  r.PartF,
					FDigit: r.FDigit,
				}
			}
			s.push(vr)
		case STORE:
			v, err := s.pop()
			if err != nil {
				return []result.Value{}, err
			}
			switch v.Type {
			case value.INTEGER:
				cols = append(cols, result.Value{Type: result.Integral, Integral: v.Int})
			case value.FLOAT:
				cols = append(cols, result.Value{Type: result.Float, Float: v.Float})
			case value.DECIMAL:
				cols = append(cols, result.Value{Type: result.Decimal, PartI: v.PartI, PartF: v.PartF, FDigit: v.FDigit})
			}
		}
	}
	return cols, nil
}
