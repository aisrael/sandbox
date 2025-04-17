package main

import (
	"fmt"
)

const Version = "0.0.1"

// main entrypoint. Run the CLI with os arguments and default dependency injections
func main() {
	fmt.Printf("Version: %s\n", Version)
}
