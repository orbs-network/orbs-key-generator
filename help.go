package main

import (
	"fmt"
	"os"
	"strings"
)

func commandVersion(requiredOptions []string) {
	log("orbs-key-generator version v%s", ORBS_KEY_GENERATOR_VERSION)
}

func commandShowHelp(requiredOptions []string) {
	fmt.Fprintf(os.Stderr, "Usage:\n\n")
	fmt.Fprintf(os.Stderr, "orbs-key-generator COMMAND\n\n")

	fmt.Fprintf(os.Stderr, "Commands:\n\n")
	sortedCommands := sortCommands()
	for _, name := range sortedCommands {
		cmd := commands[name]
		fmt.Fprintf(os.Stderr, "  %s %s %s\n", name, strings.Repeat(" ", 15-len(name)), cmd.desc)
		if cmd.args != "" {
			fmt.Fprintf(os.Stderr, "  %s  options: %s\n", strings.Repeat(" ", 15), cmd.args)
		}
		if cmd.example != "" {
			fmt.Fprintf(os.Stderr, "  %s  example: %s\n", strings.Repeat(" ", 15), cmd.example)
		}
		if cmd.example2 != "" {
			fmt.Fprintf(os.Stderr, "  %s           %s\n", strings.Repeat(" ", 15), cmd.example2)
		}
		fmt.Fprintf(os.Stderr, "\n")
	}

	os.Exit(2)
}

func sortCommands() []string {
	res := make([]string, len(commands))
	for name, cmd := range commands {
		res[cmd.sort] = name
	}
	return res
}
