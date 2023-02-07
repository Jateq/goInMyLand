package main

type Command interface {
	Execute() ([]byte, error)
	ValidateInput() bool
}
type CommandExecutor struct{}

func (c CommandExecutor) Execute(command Command) {
	if command.ValidateInput() {
		command.Execute()
	}
}

type FooComand struct {
	args []string
}
