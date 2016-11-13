package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

type Zext struct {
	name  string
	value value.Value
	cast  types.Type
}

func NewZext(name string, value value.Value, cast types.Type) *Zext {
	return &Zext{name, value, cast}
}

func (i *Zext) String() string {
	return "zext"
}

func (i *Zext) IsTerminator() bool {
	return false
}

func (i *Zext) Type() types.Type {
	return i.cast
}

func (i *Zext) Ident() string {
	return "%" + i.name
}

func (i *Zext) Llvm() string {
	return fmt.Sprintf("%s = zext %s %s to %s",
		i.name,
		i.value.Type().String(),
		i.value.String(),
		i.cast.String())
}