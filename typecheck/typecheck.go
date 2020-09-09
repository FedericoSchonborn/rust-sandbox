package typecheck

import (
	"reflect"
)

type Checker struct {
	value interface{}

	kind CheckKind
	pt   PrimitiveType
}

type CheckKind int

const (
	CheckKindPrimitive CheckKind = iota
)

type PrimitiveType int

const (
	PrimitiveTypeInt PrimitiveType = iota
)

func (pt PrimitiveType) ToReflectType() reflect.Type {
	switch pt {
	case PrimitiveTypeInt:
		return reflect.TypeOf(int(0))
	default:
		panic("unknown primitive type")
	}
}

func NewChecker(value interface{}) *Checker {
	return &Checker{
		value: value,
	}
}

func (c *Checker) Check() error {
	return nil
}
