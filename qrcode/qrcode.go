package qrcode

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bughomenoise/bsgender/seed"
	"github.com/mdp/qrterminal/v3"
)

func PrintSeedSignerQRCode(aSeed seed.Seed) {
	var qrStr string
	for _, v := range aSeed.GetIndexList() {
		iStr := strconv.Itoa(v)
		if v > 999 {
			qrStr = fmt.Sprintf("%s%s", qrStr, iStr)
		} else if v > 99 {
			qrStr = fmt.Sprintf("%s%s%s", qrStr, "0", iStr)
		} else if v > 9 {
			qrStr = fmt.Sprintf("%s%s%s", qrStr, "00", iStr)
		} else {
			qrStr = fmt.Sprintf("%s%s%s", qrStr, "000", iStr)
		}
	}

	fmt.Println("######## QR for Seedsigner ########")
	qrterminal.GenerateHalfBlock(qrStr, qrterminal.L, os.Stdout)
	fmt.Println("######## QR for Seedsigner ########")

}
