package seed

import (
	"cmp"
	"regexp"
	"slices"

	"github.com/tyler-smith/go-bip39"
	"github.com/tyler-smith/go-bip39/wordlists"
)

func sum8(str string) byte {
	var b byte
	for char := range str {
		b += byte(char)
	}
	return b
}

func eStringsToBytes(strings []string) []byte {
	var bytes []byte
	for _, str := range strings {
		bytes = append(bytes, sum8(str))
	}
	return bytes
}

type Seed struct {
	wordList  []string
	indexList []uint
}

func (s Seed) GetWords() []string {
	return s.wordList
}

func (s Seed) GetIndexes() []uint {
	return s.indexList
}

func entropyToSeed(entropy []byte) (Seed, error) {
	seedStr, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return Seed{}, err
	}
	rx := regexp.MustCompile("\\s+")
	seedWord := rx.Split(seedStr, -1)
	// var iResult []int
	wordlist := wordlists.English
	var seedIndex []uint
	for _, word := range seedWord {
		ix, ok := slices.BinarySearchFunc(wordlist, word, func(a, b string) int { return cmp.Compare(a, b) })
		if ok {
			seedIndex = append(seedIndex, uint(ix))
		}
	}

	return Seed{
		wordList:  seedWord,
		indexList: seedIndex,
	}, err
}

func StringsTo12W(eStr [16]string) (Seed, error) {
	return entropyToSeed(eStringsToBytes(eStr[:]))
}

func StringsTo24W(eStr [32]string) (Seed, error) {
	return entropyToSeed(eStringsToBytes(eStr[:]))
}

func BytesTo12W(eByte [16]byte) (Seed, error) {
	return entropyToSeed(eByte[:])
}

func ByteTo24W(eByte [32]byte) (Seed, error) {
	return entropyToSeed(eByte[:])
}
