package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/bitfield/qrand"
)

func main() {

	var random = rand.New(&qrand.Source{})

	var nchars int

	if len(os.Args) < 2 {
		err := fmt.Errorf("Password length is %d or not specified", nchars)
		fmt.Println(err.Error())
		fmt.Println("Usage : PasswordGenerator.go <no. of characters>")
		os.Exit(1)
	}

	nchars, err := strconv.Atoi(os.Args[1])
	if err != nil {
		err := fmt.Errorf("Argument (%s) processing failed\n", os.Args[1])
		fmt.Println(err.Error())
	}
	charSet := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

	var passwd []byte
	for i := 0; i < nchars; i++ {
		index := random.Intn(len(charSet))
		passwd = append(passwd, charSet[index])
	}
	fmt.Println(string(passwd))
}
