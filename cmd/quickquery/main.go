package main

import (
	"log"
	"os"

	"github.com/ritarock/quickquery/internal/action"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "quickquery"
	app.Usage = "SQL-like query for csv"
	app.Action = action.RunQuey

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
