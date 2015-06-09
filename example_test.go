package reflag

import "fmt"

func ExampleIntField() {
	type Config struct {
		Int      int
		Int32    int32 `flag:"other"`
		Int64Def int64 `flag:",def=10"`
		Int8     int8  `flag:",def=54"`
	}
	var c Config
	ParseFlag(&c, &FlagConfig{
		Args: []string{
			"-int=5", "-other=100", "-int64Def=3",
		},
	})
	fmt.Printf("%v, %v, %v, %v", c.Int, c.Int32, c.Int64Def, c.Int8)
	// Output: 5, 100, 3, 54
}
