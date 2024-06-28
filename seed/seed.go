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
	for c := range str {
		b += byte(c)
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
	indexList []int
}

func (s Seed) GetWordList() []string {
	return s.wordList
}

func (s Seed) GetIndexList() []int {
	return s.indexList
}

func GetWordList() []string {
	return wordlists.English
}

func entropyToSeed(entropy []byte) (Seed, error) {
	seedStr, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return Seed{}, err
	}
	rx := regexp.MustCompile("\\s+")
	seedWord := rx.Split(seedStr, -1)
	// var iResult []int
	wl := GetWordList()
	var seedIndex []int
	for _, v := range seedWord {
		ix, ok := slices.BinarySearchFunc(wl, v, func(a, b string) int { return cmp.Compare(a, b) })
		if ok {
			seedIndex = append(seedIndex, ix)
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
