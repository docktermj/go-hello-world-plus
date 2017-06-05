package main

import "fmt"

var gitVersion string = "0.0.0"
var gitIteration string = "0"

// TODO: Add docopt.
// TODO: Add logging.
// TODO: Add subcommand structure.
// TODO: RPC plugin system - https://github.com/hashicorp/go-plugin

func main() {
    fmt.Printf("Hello, world! from Version %s-%s\n", gitVersion, gitIteration)
}