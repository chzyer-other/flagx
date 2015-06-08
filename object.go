package reflag

import (
	"flag"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"time"
)

var (
	RegexpArgNumName = regexp.MustCompile(`^\[(\d*)\]$`)
)

var (
	TypeDuration = reflect.TypeOf(time.Duration(0))
)

type Object struct {
	Type     reflect.Type
	Val      reflect.Value
	Opt      []*Field
	Arg      []*Field
	Usage    func()
	isAllArg bool
}

func NewObject(obj interface{}) (*Object, error) {
	v := reflect.ValueOf(obj)
	t := v.Type()
	if t.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("obj must be a struct pointer")
	}
	t = t.Elem()
	v = v.Elem()
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("obj must be a struct pointer")
	}

	o := &Object{
		Type: t,
		Val:  v,
	}

	argValidate := make([]bool, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		field, err := NewField(t.Field(i), v.Field(i))
		if err != nil {
			return nil, err
		}
		if idx, ok := field.ArgIdx(); ok {
			if idx == -1 {
				if o.isAllArg {
					return nil, fmt.Errorf("only can have one [] arg")
				}
				if len(o.Arg) > 0 {
					return nil, fmt.Errorf("only can use [] arg if only have one specified [] arg")
				}
				o.isAllArg = true
			} else if o.isAllArg {
				return nil, fmt.Errorf("can't use [\\d] if have one [] arg")
			}
			if idx >= len(argValidate) {
				return nil, fmt.Errorf("invalid arg index %d", idx)
			}
			argValidate[idx] = true
			o.Arg = append(o.Arg, field)
		} else {
			o.Opt = append(o.Opt, field)
		}
	}

	for idx := range o.Arg {
		if !argValidate[idx] {
			return nil, fmt.Errorf("missing arg idx: %d", idx)
		}
	}
	return o, nil
}

func (o *Object) usage(fs *flag.FlagSet, name string) {
	arg := ""
	for _, f := range o.Arg {
		idx, _ := f.ArgIdx()
		arg += "[" + f.Name
		if idx < 0 {
			arg += " ..."
		}
		arg += "]"
	}

	io.WriteString(os.Stderr, fmt.Sprintf("%s %s\n", name, arg))
	fs.VisitAll(func(f *flag.Flag) {
		format := "  -%s=%s: %s\n"
		fmt.Fprintf(os.Stderr, format, f.Name, f.DefValue, f.Usage)
	})
}

func (o *Object) Parse() error {
	return o.ParseFlag(os.Args[0], flag.PanicOnError, os.Args[1:])
}

func (o *Object) ParseFlag(name string, eh flag.ErrorHandling, args []string) error {
	fs := flag.NewFlagSet(name, eh)
	for _, f := range o.Opt {
		f.BindFlag(fs)
	}
	fs.Usage = func() {
		o.usage(fs, name)
	}
	o.Usage = fs.Usage

	if err := fs.Parse(args); err != nil {
		return err
	}

	for _, f := range o.Arg {
		idx, _ := f.ArgIdx()
		if idx < 0 {
			if err := f.SetArgs(&f.Val, fs); err != nil {
				return err
			}
		} else {
			if err := f.SetArg(&f.Val, fs); err != nil {
				return err
			}
		}
	}

	return nil
}
