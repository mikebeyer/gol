gol
===

Simple go logging replacement

## Getting Started


~~~ go
package main

import "github.com/mikebeyer/gol"

func main() {
  log := gol.ClassicLogger()

  log.Infof("%s!", "Hello Log!")
}
~~~