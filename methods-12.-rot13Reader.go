package main

import (
	//"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13Reader) Encrypt(p []byte) {

	for i := 0; i < len(p); i++ {
		if (p[i] >= 'A' && p[i] < 'N') || (p[i] >= 'a' && p[i] < 'n') {
			p[i] += 13
		} else if (p[i] >= 'M' && p[i] < 'Z') || (p[i] >= 'm' && p[i] < 'z') {
			p[i] -= 13
		}
	}

	return
}

func (rot *rot13Reader) Read(b []byte) (n int, err error) {

	n, err = rot.r.Read(b)

	rot.Encrypt(b)

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
