package cli

// Handler Interface ...
type Handler interface {
	Handle(args []string) (resp string, err error)
}

// Cmdline struct ...
type Cmdline struct {
	Cmd         string
	Example     string
	Description string
	Handler     Handler
}
