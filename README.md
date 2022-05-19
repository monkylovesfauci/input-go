# Easy to use Input Libary for GO

## Installation

```bash
$ cd my-project
$ go get https://github.com/monkylovesfauci/input-go
```

## Example Usage

```go
package main

import (
  "fmt"

  "github.com/monkylovesfauci/input-go"
)

func main() {
  name := input.Print("Enter Name: ")
  fmt.Printf("Hello %v!", name)
}

```
