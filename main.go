package main

import (
	"c100k/ddot/cmds"
	"c100k/ddot/flags"
	"c100k/ddot/ui"
	"flag"
	"fmt"
	"os"
)

const ARG_OUT_DESC = "The destination file where to write the content of the resources"
const ARG_URI_DESC = "The URI of a resource to use (e.g 'bw://project 1 env', 'file://.env.sample') (one or many)"

const FF_CMD_LOADENV_ENABLED = true
const FF_CMD_PUBENV_ENABLED = false
const FF_CMD_VERSION_ENABLED = true

const VERSION = "0.1.0-beta.2"

func main() {
	var out string
	var uris flags.StringArr

	var commands []*flag.FlagSet

	loadenvCmd := flag.NewFlagSet("loadenv", flag.ExitOnError)
	pubenvCmd := flag.NewFlagSet("pubenv", flag.ExitOnError)
	versionCmd := flag.NewFlagSet("version", flag.ExitOnError)

	if FF_CMD_LOADENV_ENABLED {
		loadenvCmd.StringVar(&out, "out", ".env", ARG_OUT_DESC)
		loadenvCmd.Var(&uris, "uri", ARG_URI_DESC)
		commands = append(commands, loadenvCmd)
	}

	if FF_CMD_PUBENV_ENABLED {
		pubenvCmd.Var(&uris, "uri", ARG_URI_DESC)
		commands = append(commands, pubenvCmd)
	}

	if FF_CMD_VERSION_ENABLED {
		commands = append(commands, versionCmd)
	}

	if len(os.Args) < 2 {
		printHelp(commands)
		os.Exit(1)
	}

	cmd := os.Args[1]
	args := os.Args[2:]
	switch cmd {
	case loadenvCmd.Name():
		loadenvCmd.Parse(args)
		err := cmds.LoadEnv(out, uris)
		if err != nil {
			fail(err)
		}
	case pubenvCmd.Name():
		pubenvCmd.Parse(args)
		err := cmds.PubEnv()
		if err != nil {
			fail(err)
		}
	case versionCmd.Name():
		err := cmds.Version(VERSION)
		if err != nil {
			fail(err)
		}
	default:
		ui.PrintErr(fmt.Errorf("command '%s' unknown", cmd))
		printHelp(commands)
	}
}

func fail(err error) {
	ui.PrintErr(err)
	os.Exit(1)
}

func printHelp(commands []*flag.FlagSet) {
	for _, cmd := range commands {
		cmd.Usage()
	}
}
