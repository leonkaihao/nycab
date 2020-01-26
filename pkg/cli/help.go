package cli

import "fmt"

type help struct {
	Cmds []*Cmdline
}

func (c *help) Handle(args []string) (resp string, err error) {
	for _, cmd := range c.Cmds {
		resp += fmt.Sprintf("%v\n\tDescription:\t%v\n\tExample:\t%v\n", cmd.Cmd, cmd.Description, cmd.Example)
	}
	return resp, nil
}

// NewHelpCmd ...
func NewHelpCmd(cmds []*Cmdline) *Cmdline {
	return &Cmdline{
		Cmd:         "help",
		Example:     "help",
		Description: "help all commands",
		Handler:     &help{cmds},
	}
}
