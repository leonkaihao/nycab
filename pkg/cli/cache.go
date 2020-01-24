package cli

type cache struct {
}

func (c *cache) Handle(args []string) (resp string, err error) {
	return "", nil
}

// NewCacheCmd ...
func NewCacheCmd() *Cmdline {
	return &Cmdline{
		Cmd:         "cache",
		Example:     "cache clean",
		Description: "remove all cache",
		Handler:     &cache{},
	}
}
