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

func chooseInt64(a, b int64, max bool) int64 {
	if a > b {
		if max {
			return a
		} else {
			return b
		}
	}
	if max {
		return b
	} else {
		return a
	}
}

type IntSetter struct {
	Val      *reflect.Value
	Kind     reflect.Kind
	Max, Min *int64
}

func NewIntSetter(val *reflect.Value, kind reflect.Kind, max, min *int64) *IntSetter {
	is := &IntSetter{
		Val:  val,
		Kind: kind,
		Max:  max,
		Min:  min,
	}
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
	if is.Max != nil {
		i = chooseInt64(i, *is.Max, false)
	}
	if is.Min != nil {
		i = chooseInt64(i, *is.Min, true)
	}

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
	Max    *int64
	Min    *int64
}

func NewIntField(f *Field) Fielder {
	if (*DurationField).Fit(nil, f.Type) {
		return NewDurationField(f)
	}
	return &IntField{f: f}
}
func (i *IntField) Init() (err error) {
	if i.f.DefVal != "" {
		i.defval, err = strconv.ParseInt(i.f.DefVal, 10, 64)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *IntField) BindOpt(key, value string) error {
	switch key {
	case KEY_MAX, KEY_MIN:
		m, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		if key == KEY_MAX {
			i.Max = &m
		} else {
			i.Min = &m
		}
		return nil
	}
	return ErrOptMissHandler.Format(key, i.f.Name)
}

func (i *IntField) BindFlag(fs *flag.FlagSet) {
	is := NewIntSetter(i.f.Val, i.f.Type.Kind(), i.Max, i.Min)
	is.SetInt(i.defval)
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
	Max    *time.Duration
	Min    *time.Duration
}

func NewDurationField(f *Field) Fielder {
	df := &DurationField{
		f: f,
	}
	return df
}

func (d *DurationField) Init() (err error) {
	if d.f.DefVal != "" {
		d.defval, err = time.ParseDuration(d.f.DefVal)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *DurationField) BindOpt(key, value string) error {
	switch key {
	case KEY_MAX, KEY_MIN:
		v, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		if key == KEY_MAX {
			d.Max = &v
		} else {
			d.Min = &v
		}
	default:
		return ErrOptMissHandler.Format(key, d.f.Name)
	}
	return nil
}

func (d *DurationField) BindFlag(fs *flag.FlagSet) {
	ins := d.f.Instance().(*time.Duration)
	fs.DurationVar(ins, d.f.FlagName(), d.defval, d.f.Usage)
}

func (d *DurationField) Fit(t reflect.Type) bool {
	return t.AssignableTo(TypeDuration)
}

func (d *DurationField) Default() interface{} {
	return d.defval
}

func (d *DurationField) AfterParse() error {
	ptr := d.f.Instance().(*time.Duration)
	if d.Min != nil {
		if *ptr < *d.Min {
			*ptr = *d.Min
		}
	}
	if d.Max != nil {
		if *ptr > *d.Max {
			*ptr = *d.Max
		}
	}
	return nil
}
