package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"github.com/pki-io/controller"
)

func initCmd(cmd *cli.Cmd) {
	cmd.Spec = "NAME [OPTIONS]"

	params := controller.NewNodeParams()
	params.Name = cmd.StringArg("NAME", "", "name of node")

	params.OrgId = cmd.StringOpt("org-id", "", "organisation id")
	params.PairingId = cmd.StringOpt("pairing-id", "", "pairing id")
	params.PairingKey = cmd.StringOpt("pairing-key", "", "pairing key")

	cmd.Action = func() {
		app := NewNodeApp()
		logger.Info("initialising new node")

		cont, err := controller.NewNode(app.env)
		if err != nil {
			app.Fatal(err)
		}

		container, err := cont.Init(params)
		if err != nil {
			app.Fatal(err)
		}

		fmt.Println(container.Dump())
	}
}

func runCmd(cmd *cli.Cmd) {

	params := controller.NewNodeParams()

	cmd.Action = func() {
		app := NewNodeApp()
		logger.Info("initialising new node")

		cont, err := controller.NewNode(app.env)
		if err != nil {
			app.Fatal(err)
		}

		if err := cont.Run(params); err != nil {
			app.Fatal(err)
		}
	}
}
