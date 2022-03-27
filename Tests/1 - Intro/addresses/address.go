package addresses

import (
	"strings"
)

// AddressType verifies the type of a given address in string format
func AddressType(address string) string {
	validTypes := []string{"rua", "avenida", "estrada", "rodovia"}
	lowerCaseAddress := strings.ToLower(address)
	firstWord := strings.Split(lowerCaseAddress, " ")[0]

	addressHasValidType := false

	for _, addressType := range validTypes {
		if addressType == firstWord {
			addressHasValidType = true
		}
	}

	if addressHasValidType {
		return strings.Title(firstWord)
	}
	return "Invalid type"
}
