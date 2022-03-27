package main

import (
	"fmt"
	"intro-tests/addresses"
)

func main() {
	addressType := addresses.AddressType("Rua paulista")
	fmt.Println(addressType)
}
