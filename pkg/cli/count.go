package cli

type count struct {
}

func (c *count) Handle(args []string) (resp string, err error) {
	return "", nil
}

// NewCountCmd ...
func NewCountCmd() *Cmdline {
	return &Cmdline{
		Cmd:         "count",
		Example:     "count medallion=medallion1,medallion2,medallion... date=2013-01-12 refresh=true",
		Description: "count medallions pickup in a day ",
		Handler:     &count{},
	}
}
