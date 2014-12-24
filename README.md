gol
===

Simple go logging replacement

[ ![Codeship Status for mikebeyer/gofig](https://codeship.io/projects/e89699b0-6d67-0132-e135-261dee642691/status)](https://codeship.io/projects/54221)

~~~ go
package main

import "github.com/mikebeyer/gol"

func main() {
  log := gol.ClassicLogger()

  log.Errorf("%s", "Error :O")
  log.Tracef("%s", "You won't see this!")
  log.Infof("%s!", "Hello Log!")
}
~~~

Output:
```
2014-12-24T06:46:54Z [ERROR] :: Error :O
2014-12-24T06:46:54Z  [INFO] :: Hello Log!!
```