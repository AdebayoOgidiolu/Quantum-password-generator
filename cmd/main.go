package main

import (
	"flag"
	"fmt"
	"os"
	"qpass"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <number of characters>\n", os.Args[0])
		os.Exit(1)
	}
	length, err := strconv.Atoi(os.Args[1])
	if err != nil {
		err := fmt.Errorf("Argument (%s) processing failed\n", os.Args[1])
		fmt.Println(err.Error())
	}
	var characterSet string
	flag.StringVar(&characterSet, "characterSet", "unicode", "Password character set")
	flag.Parse()
	password := qpass.NewPassword(length, characterSet)
	fmt.Println(password)

}
