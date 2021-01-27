package main

import (
	"flag"
	"fmt"
	"os"
	"qpass"
	"strconv"
)

var usage = fmt.Sprintf("Usage: %s [-u] <number of characters>\n\n-u Use all Unicode characters, not just typeable characters.\n", os.Args[0])

func main() {
	unicode := flag.Bool("u", false, "Use all Unicode characters")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, usage)
		os.Exit(1)
	}
	length, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, usage)
		os.Exit(1)
	}
	g := qpass.NewGenerator()
	if *unicode {
		g.CharacterSet = qpass.Unicode
	}
	password := g.NewPassword(length)
	fmt.Println(password)
}
