gol
===

Simple go logging replacement

## Getting Started


~~~ go
package main

import "github.com/mikebeyer/gol"

func main() {
  log := gol.ClassicLogger()

  log.Errorf(">> %s <<", "Error :O")
  log.Infof("%s!", "Hello Log!")
}
~~~

```
2014-12-24T06:42:04Z [ERROR] :: Error :O!
2014-12-24T06:42:04Z  [INFO] :: Hello Log!!