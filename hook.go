package reflag

import "reflect"

var (
	IntFieldHook = Hooks{
		{(*DurationField)(nil).Fit, NewDurationField},
	}
	BoolFieldHook   = Hooks{}
	StringFieldHook = Hooks{}
	SliceFieldHook  = Hooks{}
)

type Hooks []*Hook

func (hs Hooks) Select(t reflect.Type) func(t *Field) Fielder {
	for _, h := range hs {
		if h.Fit(t) {
			return h.New
		}
	}
	return nil
}

type Hook struct {
	Fit func(t reflect.Type) bool
	New func(t *Field) Fielder
}
