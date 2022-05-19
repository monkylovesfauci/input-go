# Easy to use Input Libary for GO

## Installation

```bash
$ cd my-project
$ go get github.com/monkylovesfauci/input-go@v0.1.2
```

## Example Usage

```go
package main

import (
  "fmt"

  "github.com/monkylovesfauci/input-go"
)

func main() {
  name := input.String("Enter Name: ")
  fmt.Printf("Hello %v!", name)

  age, errAge := input.Int("Enter Age: ")
  if errAge != nil {
    panic(err)
  }
  fmt.Printf("You are %v years old!", age)
}

```
