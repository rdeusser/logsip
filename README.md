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
out to anything that satisfies the ```io.Writer``` interface.

```go
package main

import (
    "os"

    "github.com/iamthemuffinman/logsip"
)

var log = logsip.New(os.Stdout)

func main() {
    log.Info("Just some info for ya")
    log.Warn("Some info you might need to know, because you might have broken
    somethin'")
    log.Fatal("You done broke somethin'")
    log.Panic("Alright, Jim. What broke?")
}
```

This package behaves almost identi:cally to the standard log package:

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

Contributions welcome!

### Credits
Thanks to Go for being awesome and thanks to fatih for his color library which you can find [here](https://github.com/fatih/color)!
