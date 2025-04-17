package main

import (
	"fmt"
)

const Version = "v0.0.8"

// main entrypoint. Run the CLI with os arguments and default dependency injections
func main() {
	fmt.Printf("Version: %s\n", Version)
}
