package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"github.com/AdebayoOgidiolu/Quantum-password-generator/qrand"
	"encoding/binary"
)

func main() {

	type ΨSource struct{}
	
	func (s *ΨSource) Seed(seed int64) {}func (s *ΨSource) Uint64() (value uint64) {
		binary.Read(qrand.Reader, binary.BigEndian, &value)
		return value
	}
	
	func (s *ΨSource) Int63() (value int64) {
		return int64(s.Uint64() & ^uint64(1<<63))
	}
	
	var random = rand.New(&ΨSource{})

	nchars := CheckArg()

	charSet := [][]byte{
		{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'},
		{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '{', '}', '[', ']', '?'},
	}

	charPool := []byte{}
	for _, item := range charSet {
		charPool = append(charPool, item...)
	}

	var passwd []byte
	for i := 0; i < nchars; i++ {
		index := random.Intn(len(charPool))
		passwd = append(passwd, charPool[index])
	}

	fmt.Println(string(passwd))

}

func CheckArg() int {
	if len(os.Args) < 2 {
		fmt.Print("Please specify the number of password characters you want : ")
		charRdr := bufio.NewReader(os.Stdin)
		conv, err := charRdr.ReadString('\n')
		if err != nil {
			fmt.Println("Argument processing failed")
		}
		conv = strings.Trim(conv, "\r\n")
		output, err := strconv.Atoi(conv)

		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}
		return output

	} else {
		output, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Argument (%s) processing failed\n", os.Args[1])
		}
		return output
	}

}
