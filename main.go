package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"strings"
)

const LettersLower = "qwertyuiopasdfghjklzxcvbnm"
const LettersUpper = "QWERTYUIOPASDFGHJKLZXCVBNM"
const NUMBERS = "1234567890"
const SYMBOLS = `!@#$%^&*()_+=-[]{};:'"\|,./~`

func main() {
	// parse arguments
	lowercase := flag.Bool("l",true,"Include lowercase characters")
	uppercase := flag.Bool("u",true,"Include uppercase characters")
	numbers := flag.Bool("n",true,"Include numbers")
	symbols := flag.Bool("s",false,"Include symbols")
	length := flag.Int("length",16,"The length of the password")
	flag.Parse()
	// generate password
	fmt.Println(GeneratePassword(*length,*lowercase,*uppercase,*numbers,*symbols))
}

// Generates a password
func GeneratePassword(length int, lowercase, uppercase, number, symbol bool) string {
	str := ""
	if lowercase {
		str += LettersLower
	}
	if uppercase {
		str += LettersUpper
	}
	if number {
		str += NUMBERS
	}
	if symbol {
		str += SYMBOLS
	}
	//Memory shit
	var sb strings.Builder
	max := big.NewInt(int64(len(str)))
	sb.Grow(length)
	for i := 0; i < length; i++ {
		index, _ := rand.Int(rand.Reader, max)
		sb.WriteByte(str[index.Int64()])
	}
	return sb.String()
}