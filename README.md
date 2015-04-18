# go-logger
A simple, concise, colorful logger for Go.

# How to use
First, read the source code. It's REALLY simple. Here's how you use it though:

```go
package main

import (
    "github.com/iamthemuffinman/go-logger" // exported as logger
)

var log = logger.New()

func main() {
    log.Info("Just some info for ya")
    log.Warn("Some info you might need to know")
    log.Fatal("You done did something wrong")
    log.Panic("You REALLY done did something wrong")
    
    lulz := "lulz"
    log.Infof("Just some %s for ya", lulz)
    log.Warnf("You might need to know this %s", lulz)
    log.Fatalf("You done did the %s", lulz)
    log.Panicf("You REALLY done did the %s", lulz)
}
```

# Credits
Thanks to Go for being awesome and thanks to fatih for his/her color library which you can find [here](https://github.com/fatih/color)!
