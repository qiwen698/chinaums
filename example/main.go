package main

import "bytes"

func main() {
	b := bytes.Buffer{}
	b.WriteString("234")
	b.Bytes()
}
