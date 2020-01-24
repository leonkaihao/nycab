package cli

import (
	"os"
)

type quit struct {
}

func (c *quit) Handle(args []string) (resp string, err error) {
	os.Exit(0)
	return "Byebye!", nil
}

// NewQuitCmd ...
func NewQuitCmd() *Cmdline {
	return &Cmdline{
		Cmd:         "quit",
		Example:     "quit",
		Description: "quit current application",
		Handler:     &quit{},
	}
}
