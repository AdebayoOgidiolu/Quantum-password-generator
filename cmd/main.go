package main

import (
	"fmt"
	"os"
	"qpass"
	"strconv"
)

func main() {
	// TODO: future feature! Customisability!
	// p := qpass.NewPasswordGenerator()
	// p.Charset = qpass.TypeableCharacters
	// p.NewPassword(length)
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <number of characters>\n", os.Args[0])
		os.Exit(1)
	}
	length, err := strconv.Atoi(os.Args[1])
	if err != nil {
		err := fmt.Errorf("Argument (%s) processing failed\n", os.Args[1])
		fmt.Println(err.Error())
	}
	password := qpass.NewPassword(length)
	fmt.Println(password)
}
