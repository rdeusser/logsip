# go-logger
A simple, concise, colorful logger for Go.

# How to use
First, read the source code. It's REALLY simple. Here's how you use it though:

```go
package main

import (
    "github.com/iamthemuffinman/go-logger" // exported as logger
)

func main() {
    logger.Info("Just some info for ya")
    logger.Warn("Some info you might need to know")
    logger.Fatal("You done did something wrong") // Note: Don't use this with Panic unless separated by some logic
}
```

# Credits
Thanks to Go for being awesome and thanks to fatih for his/her color library which you can find [here](https://github.com/fatih/color)!
