package reflag

func Parse(obj interface{}) {
	o, err := NewObject(obj)
	if err != nil {
		panic(err)
	}
	if err = o.Parse(); err != nil {
		panic(err)
	}
}

func ParseFlag(obj interface{}, fc *FlagConfig) {
	o, err := NewObject(obj)
	if err != nil {
		panic(err)
	}
	if err = o.ParseFlag(fc); err != nil {
		panic(err)
	}
}
