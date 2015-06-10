package reflag

import "os"

func Parse(obj interface{}) {
	o, err := NewObject(obj)
	if err != nil {
		Exit(o, err)
	}
	if err = o.Parse(); err != nil {
		Exit(o, err)
	}
}

func Exit(obj *Object, err error) {
	if err != nil {
		println(err.Error())
	}
	if err == ErrUsage {
		println("\nusage:")
		obj.Usage()
	}
	os.Exit(1)
}

func ParseFlag(obj interface{}, fc *FlagConfig) {
	o, err := NewObject(obj)
	if err != nil {
		Exit(o, err)
	}
	if err = o.ParseFlag(fc); err != nil {
		Exit(o, err)
	}
}
