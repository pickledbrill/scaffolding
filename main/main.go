package main

import (
	"fmt"

	scaffolding "github.com/pickledbrill/scaffolding/scaffolding"
)

func main() {
	fmt.Println("Starting")
	scaffolding.InitializeConfig()
	scaffolding.AuthenticateUser()
}
