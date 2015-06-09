package reflag

import (
	"flag"
	"fmt"
	"reflect"
)

func init() {
	AddTypeField(NewStringField, reflect.String)
}

type StringField struct {
	f *Field
}

func NewStringField(f *Field) Fielder {
	return &StringField{f: f}
}

func (b *StringField) Init() error {
	return nil
}

func (b *StringField) BindFlag(fs *flag.FlagSet) {
	ins := b.f.Instance().(*string)
	fs.StringVar(ins, b.f.FlagName(), b.f.DefVal, b.f.Usage)
}

func (b *StringField) Default() interface{} {
	return b.f.DefVal
}

func (b *StringField) SetArg(v *reflect.Value, arg string) error {
	if !v.CanSet() {
		return fmt.Errorf("value %v is not settable", v)
	}

	if arg != "" {
		v.Set(reflect.ValueOf(arg))
	}
	return nil
}
