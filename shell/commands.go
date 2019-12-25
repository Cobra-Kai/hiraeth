package shell

type Commands struct {
	actions map[string]actionHandler
}

var commands Commands

func getCommandAction(name string) (actionHandler, bool) {
	act, ok := commands.actions[name]

	return act, ok
}

func RegisterAction(name string, fn HandlerFunc, args ...string) {
	handler := PromoteHandlerFunc(fn, args...)

	commands.actions[name] = handler
}
