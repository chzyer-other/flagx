package reflag

import (
	"flag"
	"testing"
	"time"
)

func TestType(t *testing.T) {
	type Config struct {
		IsDecode bool          `flag:"d,usage=urldecode enable"`
		Int      int           `flag:",def=5,usage=inta!"`
		Interval time.Duration `flag:",def=60s,usage=pull interval"`
		Cao      string
		Content  string `flag:"[0]"`
	}

	var c Config
	obj, err := NewObject(&c)
	if err != nil {
		t.Fatal(err)
	}
	fs := &FlagConfig{
		Name:          "./hello",
		ErrorHandling: flag.PanicOnError,
		Args:          []string{"-d", "-interval=15s", "hello"},
	}

	if err := obj.ParseFlag(fs); err != nil {
		t.Fatal(err)
	}
	// obj.Usage()
	// logex.Struct(c, obj)
}
