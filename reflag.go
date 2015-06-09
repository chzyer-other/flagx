package reflag

import "os"

func Parse(obj interface{}) {
	o, err := NewObject(obj)
	if err != nil {
		exit(err)
	}
	if err = o.Parse(); err != nil {
		exit(err)
	}
}

func exit(err error) {
	if err != nil {
		println(err.Error())
	}
	os.Exit(1)
}

func ParseFlag(obj interface{}, fc *FlagConfig) {
	o, err := NewObject(obj)
	if err != nil {
		exit(err)
	}
	if err = o.ParseFlag(fc); err != nil {
		exit(err)
	}
}
