package stdout

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/bughomenoise/bsgender/seed"
	"github.com/mdp/qrterminal/v3"
)

const WORD_LIST_LAST_INDEX uint = 2047

func indexToStr4c(index uint) (string, error) {
	var str4c string

	if index > WORD_LIST_LAST_INDEX {
		return str4c, errors.New(fmt.Sprintf("indexToStr4c: index range error -> get %d , expect 0 - %d", index, WORD_LIST_LAST_INDEX))
	}

	strIndex := strconv.Itoa(int(index))

	switch len(strIndex) {
	case 4:
		str4c = fmt.Sprintf("%s", strIndex)
	case 3:
		str4c = fmt.Sprintf("0%s", strIndex)
	case 2:
		str4c = fmt.Sprintf("00%s", strIndex)
	case 1:
		str4c = fmt.Sprintf("000%s", strIndex)
	default:
		return str4c, errors.New("indexToStr4c: switch case error")
	}

	return str4c, nil
}

func indexListToString(indexList []uint) (string, error) {
	const w12 = 12
	const w24 = 24

	var wordlistString string

	length := len(indexList)
	if !(length == w12 || length == w24) {
		return wordlistString, errors.New(fmt.Sprintf("indexListToString: wordlist length error -> get %d, expect %d or %d", length, w12, w24))
	}

	for _, index := range indexList {
		str4c, err := indexToStr4c(index)
		if err != nil {
			return wordlistString, err
		}
		wordlistString = fmt.Sprintf("%s%s", wordlistString, str4c)
	}

	return wordlistString, nil
}

func PrintSeedSignerQRCode(aSeed seed.Seed) {
	const line = "######## QR for Seedsigner ########"
	fmt.Println(line)

	wordlistString, err := indexListToString(aSeed.GetIndexes())
	if err != nil {
		fmt.Printf("can't print qrcode. Error: %s\n", err.Error())
	}

	qrterminal.GenerateHalfBlock(wordlistString, qrterminal.L, os.Stdout)

	fmt.Println(line)
}
