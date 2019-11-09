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
