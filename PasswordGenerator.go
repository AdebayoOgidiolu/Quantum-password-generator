package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/bitfield/qrand"
)

type ΨSource struct{}

func (s *ΨSource) Seed(seed int64) {}

func (s *ΨSource) Uint64() (value uint64) {
	binary.Read(qrand.Reader, binary.BigEndian, &value)
	return value
}

func (s *ΨSource) Int63() (value int64) {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func main() {

	var random = rand.New(&ΨSource{})

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
