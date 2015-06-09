package reflag

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

var (
	TypeDuration = reflect.TypeOf(time.Duration(0))
)

func init() {
	t := []reflect.Kind{
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
	}
	AddTypeField(NewIntField, t...)
}

type IntSetter struct {
	Val  *reflect.Value
	Kind reflect.Kind
}

func NewIntSetter(val *reflect.Value, kind reflect.Kind, defVal int64) *IntSetter {
	is := &IntSetter{
		Val:  val,
		Kind: kind,
	}
	is.SetInt(defVal)
	return is
}

func (is *IntSetter) Set(s string) error {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	is.SetInt(i)
	return nil
}

func (is *IntSetter) SetInt(i int64) {
	var val interface{}
	switch is.Kind {
	case reflect.Int:
		val = int(i)
	case reflect.Int8:
		val = int8(i)
	case reflect.Int16:
		val = int16(i)
	case reflect.Int32:
		val = int32(i)
	case reflect.Int64:
		val = int64(i)
	}
	is.Val.Set(reflect.ValueOf(val))
}

func (i *IntSetter) String() string {
	return fmt.Sprintf("%v", i.Val)
}

type IntField struct {
	f      *Field
	defval int64
	Max    int64
	Min    int64
}

func NewIntField(f *Field) (Fielder, error) {
	if (*DurationField).Fit(nil, f.Type) {
		return NewDurationField(f)
	}
	var err error
	field := &IntField{f: f}
	if f.DefVal != "" {
		field.defval, err = strconv.ParseInt(f.DefVal, 10, 64)
		if err != nil {
			return nil, err
		}
	}
	return field, nil
}

func (i *IntField) BindFlag(fs *flag.FlagSet) {
	is := NewIntSetter(i.f.Val, i.f.Type.Kind(), i.defval)
	fs.Var(is, i.f.FlagName(), i.f.Usage)
}

func (i *IntField) Default() interface{} {
	return i.defval
}

func (i *IntField) Argable() bool  { return true }
func (i *IntField) Argsable() bool { return false }

type DurationField struct {
	f      *Field
	defval time.Duration
	Max    time.Duration
	Min    time.Duration
}

func NewDurationField(f *Field) (Fielder, error) {
	df := &DurationField{
		f: f,
	}
	var err error
	if f.DefVal != "" {
		df.defval, err = time.ParseDuration(f.DefVal)
		if err != nil {
			return nil, err
		}
	}
	return df, nil
}

func (d *DurationField) BindFlag(fs *flag.FlagSet) {
	fs.DurationVar(d.f.Instance().(*time.Duration), d.f.FlagName(), d.defval, d.f.Usage)
}

func (d *DurationField) Fit(t reflect.Type) bool {
	return t.AssignableTo(TypeDuration)
}

func (d *DurationField) Default() interface{} {
	return d.defval
}
