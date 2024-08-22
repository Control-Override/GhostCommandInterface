// command.go
package command_interface

// Command interface that all commands will implement.
type Command interface {
	Execute(args []string)
}
