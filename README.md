# anonymity
anonymity is a simple random ip address, useragent, and header generator package for golang.

![goreport](https://goreportcard.com/badge/github.com/zirvaorg/anonymity)
![license](https://badgen.net/github/license/zirvaorg/anonymity)
[![pkg.go.dev](https://pkg.go.dev/badge/github.com/zirvaorg/anonymity)](https://pkg.go.dev/github.com/zirvaorg/anonymity)
![sourcegraph](https://sourcegraph.com/github.com/zirvaorg/anonymity/-/badge.svg)

## Installation
```bash
go get github.com/zirvaorg/anonymity
```

## Usage
```go
package main

import (
    "fmt"
    "github.com/zirvaorg/anonymity"
)

func main() {
    a := anonymity.New()
	
    // Generate random ip address
    ip := a.IP()
    fmt.Println(ip)

    // Generate random user agent
    ua := a.UserAgent()
    fmt.Println(ua)

    // Generate random headers
    headers := a.Header(ip)
    fmt.Println(headers)
}
```