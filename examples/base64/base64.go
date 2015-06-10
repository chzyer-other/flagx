package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/chzyer/reflag"
)

type Config struct {
	Decode      bool     `flag:"d,usage=decode base64"`
	UrlEncoding bool     `flag:"url,usage=use urlencoding"`
	Content     []string `flag:"[]"`
}

func main() {
	var c Config
	reflag.Parse(&c)

	encoding := base64.StdEncoding
	if c.UrlEncoding {
		encoding = base64.URLEncoding
	}

	for _, content := range c.Content {
		if c.Decode {
			dsc, err := encoding.DecodeString(content)
			if err != nil {
				fmt.Fprintf(os.Stderr, "decoding '%v' error: %v\n", content, err)
				continue
			}
			fmt.Println(string(dsc))
			continue
		}
		fmt.Println(encoding.EncodeToString([]byte(content)))
	}
}
