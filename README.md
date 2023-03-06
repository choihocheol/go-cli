```go
import "github.com/choihocheol/go-cli"
```

# Go CLI

Go CLI is a Go library for writing Bash scripts using Golang code.

## Quick guide : Example

```go
package main

import (
	"fmt"

	"github.com/choihocheol/go-cli"
)

func main() {
	command := "ls"
	output, err := gocli.Run(command)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
```

```go
package main

import (
	"fmt"

	"github.com/choihocheol/go-cli"
)

func main() {
	command := `
		sed -i.bak -e \
		"s/^addr = \".*:8000\"/addr = \"0.0.0.0:8080\"/" \
		config.toml
	`
	output, err := gocli.Run(command)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
```
