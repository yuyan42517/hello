# hello

just say hello

## Install

import code

```bash
go get github.com/yuyan42517/hello@latest
```

install cmd

```bash
go install github.com/yuyan42517/hello/cmd/hello@latest
```

## Example

Here's a simple example as follows:

```go
package main

import (
  "fmt"
  "github.com/yuyan42517/hello"
)

func main() {
  result := hello.Hello("jack")
  fmt.Println(result)
}
```