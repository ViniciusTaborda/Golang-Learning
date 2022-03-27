//Unit Testing
package addresses

import (
	"testing"
)

type testCases struct {
	inputAddress              string
	expectedOutputAddressType string
}

func TestAddressType(t *testing.T) {
	// Runs tests in parallel
	// t.Parallel()

	testCasesValues := []testCases{
		{"Rua Marechal", "Rua"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Avenida XV", "Avenida"},
		{"Estrada 116", "Estrada"},
		{"ESTRADA 116", "Estrada"},
		{"", "Invalid type"},
		{"Praça do japão", "Invalid type"},
	}

	for _, testCasesValue := range testCasesValues {
		functionReturn := AddressType(testCasesValue.inputAddress)
		if functionReturn != testCasesValue.expectedOutputAddressType {
			t.Errorf("Expected %s received %s", testCasesValue.expectedOutputAddressType, functionReturn)
		}
	}
}
