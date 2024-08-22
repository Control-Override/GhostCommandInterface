// command.go
package command

// Command interface that all commands will implement.
type Command interface {
	Execute(args []string)
}
