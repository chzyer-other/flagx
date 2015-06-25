package main

import (
	"fmt"
	"github.com/chzyer/flagx"
	"net/url"
)

func ParseUrl(u string) {
	U, err := url.Parse(u)
	if err != nil {
		println("parse string '%s' error: %v", u, err)
		return
	}
	query := U.Query()
	U.RawQuery = ""
	println(U.String() + ":")
	for k, v := range query {
		if len(v) == 1 {
			println(" ", k, "=", v[0])
		} else {
			println(" ", k, "=", fmt.Sprintf("%v", v))
		}
	}
}

type Config struct {
	Urls []string `flag:"[]"`
}

func main() {
	var c Config
	flagx.Parse(&c)
	for _, o := range c.Urls {
		ParseUrl(o)
	}
}
