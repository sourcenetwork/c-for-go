package translator

import (
	"bytes"
	"fmt"
	"go/token"
)

type CTypeKind int

const (
	TypeKind CTypeKind = iota
	PlainTypeKind
	VoidTypeKind
	DefineKind
	StructKind
	OpaqueStructKind
	UnionKind
	FunctionKind
	EnumKind
)

type CType interface {
	GetBase() string
	GetTag() string
	SetRaw(x string)
	CGoName() string
	GetArrays() string
	GetVarArrays() uint8
	GetPointers() uint8
	SetPointers(uint8)
	AddArray(uint64)
	//
	IsConst() bool
	IsOpaque() bool
	IsComplete() bool
	Kind() CTypeKind
	String() string
	Copy() CType
}

type (
	Value interface{}
)

type CDecl struct {
	Spec       CType
	Name       string
	Value      Value
	Expression string
	IsStatic   bool
	IsTypedef  bool
	IsDefine   bool
	Pos        token.Pos
	Src        string
}

func (c *CDecl) Kind() CTypeKind {
	return c.Spec.Kind()
}

func (c *CDecl) IsOpaque() bool {
	return c.Spec.IsOpaque()
}

func (c *CDecl) IsConst() bool {
	return c.Spec.IsConst()
}

func (c *CDecl) SetPointers(n uint8) {
	c.Spec.SetPointers(n)
}

func (c *CDecl) AddArray(size uint64) {
	c.Spec.AddArray(size)
}

func (c CDecl) String() string {
	buf := new(bytes.Buffer)
	switch {
	case len(c.Name) > 0:
		fmt.Fprintf(buf, "%s %s", c.Spec, c.Name)
	default:
		buf.WriteString(c.Spec.String())
	}
	if len(c.Expression) > 0 {
		fmt.Fprintf(buf, " = %s", string(c.Expression))
	}
	return buf.String()
}