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

    // output:
    // 78.82.211.197
    // ::::::::::::::::::::::::::

    // Generate random user agent
    ua := a.UserAgent()
    fmt.Println(ua)

    // output:
    // Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; AS; rv:11.0) like Gecko
    // ::::::::::::::::::::::::::

    // Generate random headers
    headers := a.Header(ip)
    fmt.Println(headers)

    // output:
    // X-Originating-IP		78.82.211.197
    // X-Remote-IP		78.82.211.197
    // X-Remote-Addr		78.82.211.197
    // X-Client-IP		78.82.211.197
    // X-Host			78.82.211.197
    // X-Forwarded-Host		78.82.211.197
    // X-Forwarded-For		78.82.211.197
    // X-Forwarded-Server	78.82.211.197
    // ::::::::::::::::::::::::::
}
```
