package cli

import (
	"fmt"
	"strings"

	"github.com/leonkaihao/nycab/pkg/config"
)

// Parser ...
type Parser interface {
	Parse(cmd string) error
}

type parser struct {
	cmds []*Cmdline
}

// Parse current command input
func (p *parser) Parse(cmd string) error {
	args := strings.Fields(cmd)
	if len(args) == 0 {
		return fmt.Errorf("null command")
	}
	for _, cmd := range p.cmds {
		if cmd.Cmd == args[0] {
			resp, err := cmd.Handler.Handle(args[1:])
			if err != nil {
				return err
			}
			fmt.Println(resp)
			return nil
		}
	}
	return fmt.Errorf("unknown command %v", args[0])
}

// NewParser ...
func NewParser(conf *config.Cli) Parser {
	cmds := []*Cmdline{
		NewCountCmd(conf),
		NewCacheCmd(conf),
		NewQuitCmd(),
	}

	helpCmd := NewHelpCmd(cmds)
	cmds = append(cmds, helpCmd)
	return &parser{cmds}
}
