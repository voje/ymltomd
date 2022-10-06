package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/voje/ymltomd/internal/testdocs"
	"github.com/voje/ymltomd/internal/ymltomd"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "ymltomd",
		Usage:  "Read .yml, generate a nice .md documentation",
		Action: Run,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func Run(*cli.Context) error {
	y := ymltomd.NewYtm(ymltomd.YtmCfg{})
	y.Read(testdocs.Doc1)
	return nil
}
