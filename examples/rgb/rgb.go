package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chzyer/flagx"
)

type Config struct {
	Red   int `flag:"[0]"`
	Green int `flag:"[1]"`
	Blue  int `flag:"[2]"`
}

func main() {
	var c Config
	flagx.Parse(&c)
	if c.Green == 0 {
		c.Green = c.Red
	}
	if c.Blue == 0 {
		c.Blue = c.Red
	}

	f := func(r string) string {
		if len(r) < 2 {
			r = strings.Repeat("0", 2-len(r)) + r
		}
		return r
	}

	r := strconv.FormatInt(int64(c.Red), 16)
	g := strconv.FormatInt(int64(c.Green), 16)
	b := strconv.FormatInt(int64(c.Blue), 16)
	fmt.Printf("#%v%v%v\n", f(r), f(g), f(b))
}
