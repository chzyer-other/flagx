package flagx

import (
	"fmt"
	"time"
)

func ExampleSliceField() {
	type Config struct {
		Slice []string
	}
	var c Config
	ParseFlag(&c, &FlagConfig{Args: []string{
		"-slice=hello", "-slice=bye",
	}})
	fmt.Printf("%v", c)
	// Output:
	// {[hello bye]}
}

func ExampleStringField() {
	type Config struct {
		Desc string
		Max  string `flag:"max=3"`
	}
	var c Config
	ParseFlag(&c, &FlagConfig{Args: []string{
		"-desc=cao", "-max=13日1jfn",
	}})
	fmt.Printf("%v", c)
	// Output:
	// {cao 13日}
}

func ExampleBoolField() {
	type Config struct {
		Ok   bool
		True bool `flag:"def=true"`
	}
	var c Config
	ParseFlag(&c, &FlagConfig{Args: []string{
		"-ok",
	}})
	fmt.Printf("%v", c)
	// Output: {true true}
}

func ExampleDurationField() {
	type Config struct {
		Second time.Duration
		Minute time.Duration `flag:"def=1m"`
		Min    time.Duration `flag:"min=1h"`
		Max    time.Duration `flag:"max=30s"`
	}
	var c Config
	ParseFlag(&c, &FlagConfig{
		Args: []string{
			"-second=5s", "-min=8s", "-max=10h",
		},
	})
	fmt.Printf("%v", c)
	// Output: {5s 1m0s 1h0m0s 30s}
}

func ExampleIntField() {
	type Config struct {
		Int      int
		Int32    int32 `flag:"other"`
		Int64Def int64 `flag:"def=10"`
		Int8     int8  `flag:"def=54"`
		Int16    int16 `flag:"int16;min=6;def=5;max=7"`
		Max      int   `flag:"max=5"`
		Min      int   `flag:"min=3"`
	}
	var c Config
	ParseFlag(&c, &FlagConfig{
		Args: []string{
			"-int=5", "-other=100", "-int64Def=3",
			"-max=100", "-min=2",
		},
	})
	fmt.Printf("%v", c)
	// Output: {5 100 3 54 6 5 3}
}
