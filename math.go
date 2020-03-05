package xlib

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Atoi coverts string to int type
func Atoi(str string) int {
	// string to int
	i, err := strconv.Atoi(str)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}

// Atof coverts string to float type
func Atof(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	return f
}

// RandomIntSlice generate slice of random Int [minimum, maximum]
func RandomIntSlice(length int, minimum int, maximum int) []int {
	intSlice := make([]int, 0, length)
	span := maximum - minimum + 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		intSlice = append(intSlice, rand.Intn(span)+minimum)
	}
	return intSlice
}

// RandomInt generate random Int
func RandomInt(maximum int) int {
	rand.Seed(time.Now().UnixNano())
	randInt := rand.Intn(maximum)
	fmt.Printf("funcRandomInt -> randInt:%v", randInt)
	return randInt
}

// RandomSentence produce string of char of ascii 32 ~ 126
func RandomSentence(wordNum int) string {
	var ret string

	for i := 0; i < wordNum; i++ {
		wordLength := RandomInt(8) + 1
		word := RandomString(wordLength)
		if ret == "" {
			ret = word
		} else {
			ret += " " + word
		}
	}
	return ret
}

// RandomString produce string of char of ascii 32 ~ 126
func RandomString(length int) string {
	ascIIa := 97
	var chars []rune

	intSlice := RandomIntSlice(length, 0, 25)

	for _, v := range intSlice {
		chars = append(chars, int32(ascIIa+v))
	}
	return string(chars)
}

// RandomDNA generate random DNA sequence
func RandomDNA(length int) string {
	var DNACode = [4]byte{'A', 'C', 'T', 'G'}

	intSlice := RandomIntSlice(length, 0, 4)
	var seq = make([]byte, 0, length)
	for _, v := range intSlice {
		seq = append(seq, DNACode[v])
	}
	return string(seq)
}

// MaxInt return the max number
func MaxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
