package helper

import "fmt"

// Writes a hardcoded message to the standard output.
// Exported funcions should have a comment above it's definition.
func Write() {
	fmt.Println("Writing from the helper file.")
	Write2()
}
