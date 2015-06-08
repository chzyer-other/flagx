package reflag

import (
	"flag"
	"reflect"
	"strconv"
	"time"
)

func init() {
	t := []reflect.Kind{
		reflect.Int,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
	}
	AddTypeField(NewIntField, t...)
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
	fs.IntVar(i.f.Instance().(*int), i.f.FlagName(), int(i.defval), i.f.Usage)
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
