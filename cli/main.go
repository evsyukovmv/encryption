package main

import (
	"flag"
	"fmt"
	"github.com/evsyukovmv/encryption"
	"strings"
)

var d = flag.Bool("d", false, "decrypt input (encrypt be default)")

func main() {
	flag.Parse()
	string := strings.ToLower(strings.Join(flag.Args(), " "))

	if *d {
		fmt.Println(encryption.Decode(string))
	} else {
		fmt.Println(encryption.Encode(string))
	}
}