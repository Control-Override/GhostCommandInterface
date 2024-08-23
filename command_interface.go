// command.go
package command_interface

// Command interface that all commands will implement.
type Command interface {
	Execute(args []string)
    Author() string
    Alias() string
    Help(isShort bool, args []string) string
}
