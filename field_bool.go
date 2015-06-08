package reflag

import (
	"flag"
	"reflect"
)

func init() {
	AddTypeField(NewBoolField, reflect.Bool)
}

type BoolField struct {
	f      *Field
	defval bool
}

func NewBoolField(f *Field) (Fielder, error) {
	defval := true
	if f.DefVal == "" {
		defval = true
	}

	return &BoolField{
		f:      f,
		defval: defval,
	}, nil
}

func (b *BoolField) BindFlag(fs *flag.FlagSet) {
	fs.BoolVar(b.f.Instance().(*bool), b.f.FlagName(), b.defval, b.f.Usage)
}

func (b *BoolField) Default() interface{} {
	return b.defval
}
