package cli

import (
	"os"

	"github.com/ccreater222/go-cli/common"
)

func Run() {
	parsers := common.Args()
	command := parsers.Subcommand()
	if command == nil {
		parsers.WriteHelp(os.Stderr)
		return
	}
	command.(common.Command).Execute()
}
