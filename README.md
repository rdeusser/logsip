# Logsip
[![GoDoc](https://godoc.org/github.com/iamthemuffinman/logsip?status.svg)](https://godoc.org/github.com/iamthemuffinman/logsip)
[![Build Status](https://travis-ci.org/iamthemuffinman/logsip.svg?branch=master)](https://travis-ci.org/iamthemuffinman/logsip) [![Go Report Card](https://goreportcard.com/badge/github.com/iamthemuffinman/logsip)](https://goreportcard.com/report/github.com/iamthemuffinman/logsip)

A simple, concise, colorful logger for Go.

### How to use
First, read the source code. It's REALLY simple. Here's a screenshot:

![Just a screenshot here, nothin' to see](/screenshot.png?raw=true)

Here's a basic example:

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

```logsip.Default()``` defaults to ```os.Stdout``` but you can write
out to anything that satisfies the ```io.Writer``` interface if you use
```logsip.New()``` instead of ```logsip.Default()```.

```go
package main

import (
    "os"

    "github.com/iamthemuffinman/logsip"
)

// That string at the end is for package identification. It helps with
// debugging. Plus it's nice to see what logs are coming from where without having
// to look at the code.
var log = logsip.New(os.Stdout, "main")

func main() {
    log.Info("Just some info for ya")
    log.Warn("Some info you might need to know, because you might have broken
    somethin'")
    log.Fatal("You done broke somethin'")
    log.Panic("Alright, Jim. What broke?")
}
```

This package behaves almost identically to the standard log package:

```go
package main

import (
    "github.com/iamthemuffinman/logsip"
)

var log = logsip.Default()

func main() {
    stuff := "stuff"

    log.Infof("Here's some %s", stuff)

    log.Fatalln("You broke it, Jim.")
}
```

There's also a Debug mode.

```go
package main

import (
    "github.com/iamthemuffinman/logsip"
)

var log = logsip.Default()

func main() {
    log.DebugMode = true

    log.Debug("You see this dog?")
}
```

Contributions welcome!
