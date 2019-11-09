# Go Terminput

Package terminput provides terminal keyboard input for interactive command-line tools.

## Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/pkg/term"
	"github.com/tj/go-terminput"
)

func main() {
	t, err := term.Open("/dev/tty")
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}

	t.SetRaw()
	defer t.Restore()

	fmt.Printf("Type something, use 'q' to exit.\r\n")

	for {
		e, err := terminput.Read(t)
		if err != nil {
			log.Fatalf("error: %s\n", err)
		}

		if e.Key() == terminput.KeyEscape || e.Rune() == 'q' {
			break
		}

		fmt.Printf("%s — shift=%v ctrl=%v alt=%v meta=%v\r\n", e.String(), e.Shift(), e.Ctrl(), e.Alt(), e.Meta())
	}
}
```

---

[![GoDoc](https://godoc.org/github.com/tj/go-terminput?status.svg)](https://godoc.org/github.com/tj/go-terminput)
![](https://img.shields.io/badge/license-MIT-blue.svg)
![](https://img.shields.io/badge/status-stable-green.svg)

## Sponsors

This project is [sponsored](https://github.com/sponsors/tj) by [CTO.ai](https://cto.ai/), making it easy for development teams to create and share workflow automations without leaving the command line.

[![](https://apex-software.imgix.net/github/sponsors/cto.png)](https://cto.ai/)

