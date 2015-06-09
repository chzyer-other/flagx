package reflag

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	KEY_USAGE = "usage"
	KEY_DEF   = "def"
	KEY_OP    = "op"
)

var (
	TypeField = map[reflect.Kind]func(f *Field) (Fielder, error){}
)

func AddTypeField(f func(f *Field) (Fielder, error), kinds ...reflect.Kind) {
	for _, k := range kinds {
		TypeField[k] = f
	}
}

type Fielder interface {
	Default() interface{}
	BindFlag(*flag.FlagSet)
}

type ArgsSetter interface {
	SetArgs(v *reflect.Value, args []string) error
}

type ArgSetter interface {
	SetArg(v *reflect.Value, arg string) error
}

type Arg interface {
	Arg(n int) string
}

type Field struct {
	Name     string       // field name
	Type     reflect.Type // field type
	Usage    string
	flagName string
	DefVal   string
	Val      *reflect.Value
	fielder  Fielder
}

func NewField(t reflect.StructField, val reflect.Value) (*Field, error) {
	f := &Field{
		Name: t.Name,
		Val:  &val,
		Type: t.Type,
	}
	var err error

	if err = f.decodeTag(t.Tag); err != nil {
		return nil, err
	}

	fielderFunc, ok := TypeField[t.Type.Kind()]
	if !ok {
		return nil, fmt.Errorf("not support type: %v", t.Type)
	}
	if ok {
		f.fielder, err = fielderFunc(f)
		if err != nil {
			return nil, err
		}
	}

	return f, nil
}

func (f *Field) ArgIdx() (int, bool) {
	list := RegexpArgNumName.FindAllStringSubmatch(f.flagName, -1)
	if len(list) == 0 {
		return 0, false
	}
	idx, err := strconv.Atoi(list[0][1])
	if err != nil {
		idx = -1
	}
	return idx, true
}

func (f *Field) decodeTag(t reflect.StructTag) error {
	tags := strings.Split(t.Get("flag"), ",")
	f.flagName = tags[0]
	for i := 1; i < len(tags); i++ {
		sp := strings.Split(tags[i], "=")
		if len(sp) != 2 {
			return fmt.Errorf("wrong length")
		}
		switch sp[0] {
		case KEY_USAGE:
			f.Usage = sp[1]
		case KEY_DEF:
			f.DefVal = sp[1]
		}
	}
	return nil
}

func (f *Field) FlagName() string {
	if f.flagName == "" {
		return strings.ToLower(f.Name[:1]) + f.Name[1:]
	}
	return f.flagName
}

func (f *Field) String() string {
	return fmt.Sprintf("&%+v", *f)
}

func (f *Field) Default() interface{} {
	if f.fielder != nil {
		return f.fielder.Default()
	}
	val := f.DefVal
	switch f.Type.Kind() {
	case reflect.Bool:
		if val == "true" {
			return true
		}
		return false
	case reflect.Int, reflect.Int64:

		if f.Type.AssignableTo(TypeDuration) {
			d, err := time.ParseDuration(val)
			if err != nil {
				panic(err)
			}
			return d
		}
		i, _ := strconv.Atoi(val)
		return i
	default:
		return nil
	}
}

func (f *Field) Instance() interface{} {
	return f.Val.Addr().Interface()
}

func (f *Field) BindFlag(fs *flag.FlagSet) {
	f.fielder.BindFlag(fs)
}

func (f *Field) SetArgs(v *reflect.Value, fs *flag.FlagSet) error {
	as, ok := f.fielder.(ArgsSetter)
	if !ok {
		return fmt.Errorf("field %v is not settable args", f.fielder)
	}
	return as.SetArgs(v, fs.Args())
}

func (f *Field) SetArg(v *reflect.Value, fs *flag.FlagSet) error {
	as, ok := f.fielder.(ArgSetter)
	if !ok {
		return fmt.Errorf("field %v is not settable arg", f)
	}
	idx, ok := f.ArgIdx()
	if !ok {
		return fmt.Errorf("field %v is not define to arg", f)
	}
	return as.SetArg(v, fs.Arg(idx))
}
