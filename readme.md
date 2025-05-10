Demonstration of creating go module.

# How to install
to install the module, open terminal and type command bellow
```
go get github.com/pradiptaagus/go-greeting
```

Example of the usage like bellow
```
package main

import (
	"fmt"

	go_greeting "github.com/pradiptaagus/go-greeting/v2"
)

func main() {
	fmt.Println(go_greeting.SayHello("John Doe"))
}

```

# Features
This module has following features

- `SayHello(name string) string`
this method accept an argument `name` and will return `Hello + name`. For example if the argument is `John Doe`, it will be `Hello John Doe`