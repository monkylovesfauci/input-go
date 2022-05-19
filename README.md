# Very EASY & SIMPLE to use Input Libary for GO

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
  age, errAge := input.Int("Enter Age: ")
  if errAge != nil {
    panic(errAge)
  }

  fmt.Printf("Hello %v!\nYou are %v years old!\n", name, age)
}

```
