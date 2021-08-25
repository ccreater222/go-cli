package main

import (
	"github.com/ccreater222/go-cli/common"
	"os"
)

func Run()  {
	parsers := common.Args()
	command := parsers.Subcommand()
	if command == nil{
		parsers.WriteHelp(os.Stderr)
	}
	command.(common.Command).Execute()
}