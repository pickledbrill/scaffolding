package main

import (
	"flag"

	"github.com/pickledbrill/scaffolding/scaffolding"
)

func main() {
	templatePtr := flag.String("template", "default", "Specify the project template used to create the repository.")
	visibilityPtr := flag.Bool("private", false, "Specify the project is public repository or private repository.")
	namePtr := flag.String("name", "", "Repository name.")

	flag.Parse()

	scaffolding.InitializeConfig()
	if err := scaffolding.VerifyFlags(*templatePtr, *namePtr); err != nil {
		panic(err)
	}

	scaffolding.Process(*templatePtr, *namePtr, *visibilityPtr)
}
