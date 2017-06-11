package main

import "fmt"

var buildVersion string = "0.0.0"
var buildIteration string = "0"

// TODO: Add docopt.
// TODO: Add logging.
// TODO: Add subcommand structure.
// TODO: RPC plugin system - https://github.com/hashicorp/go-plugin

func main() {
    fmt.Printf("Hello, world! from Version %s-%s\n", buildVersion, buildIteration)
}