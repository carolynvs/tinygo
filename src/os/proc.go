package os

import (
	"syscall"
)

// Exit the current program with the specified status code.
func Exit(code int) {
	// TODO: This doesn't work, not sure what needs to be implemented to make it work?
	syscall.Exit(code)
}
