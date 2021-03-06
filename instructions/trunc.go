package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

type Trunc struct {
	block value.Value
	name  string
	value value.Value
	cast  types.Type
}

func NewTrunc(block value.Value, name string, value value.Value, cast types.Type) *Trunc {
	return &Trunc{block, name, value, cast}
}

func (i *Trunc) Block() value.Value {
	return i.block
}

func (i *Trunc) IsTerminator() bool {
	return false
}

func (i *Trunc) Type() types.Type {
	return i.cast
}

func (i *Trunc) Ident() string {
	return "%" + i.name
}

func (i *Trunc) Llvm() string {
	return fmt.Sprintf("%%%s = trunc %s %s to %s",
		i.name,
		i.value.Type().String(),
		i.value.Ident(),
		i.cast.String())
}
