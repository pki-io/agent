// ThreatSpec package main
package main

import (
	"github.com/jawher/mow.cli"
	"os"
)

// Global CLI options
var version *bool
var logLevel *string
var logging *string

func main() {

	cmd := cli.App("pki.io.d", "pki.io node agent")

	// Global options
	logLevel = cmd.StringOpt("l log-level", "info", "log level")
	logging = cmd.StringOpt("logging", "", "alternative logging configuration")

	cmd.Before = func() {
		initLogging(*logLevel, *logging)
	}
	cmd.After = func() {
		logger.Close()
	}

	cmd.Command("init", "Initialize an organization", initCmd)
	cmd.Command("run", "Run node tasks", runCmd)

	cmd.Run(os.Args)
}
