package main

import (
	"github.com/jawher/mow.cli"
	"github.com/pki-io/controller"
)

type NodeApp struct {
	env *controller.Environment
}

func NewNodeApp() *NodeApp {
	app := new(NodeApp)
	app.env = controller.NewEnvironment(logger)
	return app
}

func (app *NodeApp) Exit() {
	logger.Flush()
	cli.Exit(0)
}

func (app *NodeApp) Fatal(err error) {
	logger.Critical(err)
	cli.Exit(1)
}
