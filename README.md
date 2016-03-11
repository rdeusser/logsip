# Logsip [![Build Status](https://travis-ci.org/iamthemuffinman/logsip.svg?branch=master)](https://travis-ci.org/iamthemuffinman/logsip)
A simple, concise, colorful logger for Go.

### How to use
First, read the source code. It's REALLY simple. Here's a screenshot:

![Just a screenshot here, nothin' to see](/screenshot.png?raw=true)

And here's how you use it:

```go
package main

import (
    "os"

    "github.com/iamthemuffinman/logsip"
)

var log = logsip.New(os.Stdout)

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
    
    log.Infoln("Just some info for ya")
    log.Warnln("Some info you might need to know")
    log.Fatalln("You done did something wrong")
    log.Panicln("You REALLY done did something wrong")
}
```

New() will write to anything that satisfies the io.Writer method so get creative!

Logsip also provides a Default option if you just want to use os.Stdout:

```go
package main

import (
    "github.com/iamthemuffinman/logsip"
)

var log = logsip.Default()

func main() {
    log.Info("Just some info for ya")
}
```
Contributions welcome!

### Credits
Thanks to Go for being awesome and thanks to fatih for his color library which you can find [here](https://github.com/fatih/color)!
