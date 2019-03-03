package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const ORBS_KEY_GENERATOR_VERSION = "0.1.0"

type command struct {
	desc            string
	args            string
	example         string
	example2        string
	handler         func([]string)
	sort            int
	requiredOptions []string
}

var commands = map[string]*command{
	"node": {
		desc:            "generate a new node address and private key",
		example:         "orbs-key-generator node",
		handler:         commandGenerateNodeKey,
		sort:            0,
		requiredOptions: nil,
	},
	"client": {
		desc:            "generate a new client address and public/private key pair",
		example:         "orbs-key-generator client",
		handler:         commandGenerateClientKey,
		sort:            1,
		requiredOptions: nil,
	},
	"version": {
		desc:            "print orbs-key-generator version",
		handler:         commandVersion,
		sort:            2,
		requiredOptions: nil,
	},
	"help": {
		desc:            "print this help screen",
		sort:            3,
		requiredOptions: nil,
	},
}

func main() {
	flag.Usage = func() { commandShowHelp(nil) }
	commands["help"].handler = commandShowHelp

	if len(os.Args) <= 1 {
		commandShowHelp(nil)
	}
	cmdName := os.Args[1]
	cmd, found := commands[cmdName]
	if !found {
		die("Command '%s' not found, run 'orbs-key-generator help' to see available commands.", cmdName)
	}

	requiredOptions := []string{}
	if len(cmd.requiredOptions) > 0 {
		if len(os.Args) < 2+len(cmd.requiredOptions) {
			die("Command '%s' is missing required arguments %v.", cmdName, cmd.requiredOptions)
		}
		requiredOptions = os.Args[2 : 2+len(cmd.requiredOptions)]
		for i, requiredOption := range requiredOptions {
			if strings.HasPrefix(requiredOption, "-") {
				die("Command '%s' argument %d should be %s.", cmdName, i+1, cmd.requiredOptions[i])
			}
		}
	}

	os.Args = os.Args[2+len(cmd.requiredOptions)-1:]
	flag.Parse()

	cmd.handler(requiredOptions)
}

func die(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "ERROR:\n  ")
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintf(os.Stderr, "\n\n")
	os.Exit(2)
}

func log(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, format, args...)
	fmt.Fprintf(os.Stdout, "\n")
}
