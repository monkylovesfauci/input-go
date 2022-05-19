# Easy to use Input Libary for GO

## Example Usage

```go
package main

import (
  "fmt"

  "github.com/monkylovesfauci/go-input"
)

func main() {
  name := input.Print("Enter Name: ")
  fmt.Printf("Hello %v!", name)
}

```
