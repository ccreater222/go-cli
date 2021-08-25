# go-cli
a easy command line project struct based on go-arg

## quickstart

### example for go-cli

```

commands/
    test.go
main.go

```
push.go:

```go
package main

import (
	cli "github.com/ccreater222/go-cli"
	_ "example.com/commands"
)

func main()  {
	cli.Run()
}

```


push.go

```go
package commands

import (
	"fmt"
	"github.com/ccreater222/go-cli/common"
)

type PushCmd struct {

	Remote      string `arg:"positional"`
	Branch      string `arg:"positional"`
	SetUpstream bool   `arg:"-u"`
}

func (receiver PushCmd) Execute()  {
	fmt.Println(receiver.Remote)

}

func init()  {
	common.RegisterCmd("Scan",PushCmd{},"arg:\"subcommand:push\"")
}
```

### translate to go-arg

```go

func push(arg Push){
    fmt.Println(arg.Remote)
}

type PushCmd struct {
	Remote      string `arg:"positional"`
	Branch      string `arg:"positional"`
	SetUpstream bool   `arg:"-u"`
}
var args struct {
	Push     *PushCmd     `arg:"subcommand:push"`
	Quiet    bool         `arg:"-q"` // this flag is global to all subcommands
}

parsers = arg.MustParse(&args)
command := parsers.Subcommand()
if command == nil {
    parsers.WriteHelp(os.Stderr)
    return
}
if parsers.SubcommandName()=="push"{
    push(command)
}
```