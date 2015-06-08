package reflag

import (
	"flag"
	"reflect"
)

func init() {
	AddTypeField(NewSliceField, reflect.Slice)
}

type SliceField struct {
	f *Field
}

func NewSliceField(f *Field) (Fielder, error) {
	return &SliceField{f: f}, nil
}

func (b *SliceField) BindFlag(fs *flag.FlagSet) {

	// fs.StringVar(ins, b.f.FlagName(), b.f.DefVal, b.f.Usage)
}

func (b *SliceField) Default() interface{} {
	return b.f.DefVal
}
