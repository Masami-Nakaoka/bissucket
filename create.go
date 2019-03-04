package main

import (
	"fmt"
	"github.com/urfave/cli"
)

func Create(c *cli.Context) error {
	// Flagからデータを取って
	// bitbucketに送信
	title := c.String("t")
	priority := c.String("p")
	kind := c.String("k")
	raw := c.String("raw")
	fmt.Println(title)
	return nil
}
