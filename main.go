package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pkg/term"
)

func getch() []byte {
	t, _ := term.Open("/dev/tty")
	term.RawMode(t)
	bytes := make([]byte, 3)
	numRead, err := t.Read(bytes)
	t.Restore()
	t.Close()
	if err != nil {
		return nil
	}
	return bytes[0:numRead]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, 64)
	exit := false

	go func() {
		for {
			c := getch()
			switch {
			case bytes.Equal(c, []byte{3}):
				exit = true
				return
			case bytes.Equal(c, []byte{13}): // left
				fmt.Println("=====================================================")
			}
		}
	}()

	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		if exit {
			return
		}

		// process buf
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		fmt.Println(string(buf))
	}
}
