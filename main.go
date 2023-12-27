package main

import (
	"fmt"

	"github.com/ritarock/quickquery/cli"
)

func main() {
	if err := cli.Run(); err != nil {
		fmt.Println(err)
	}
}
