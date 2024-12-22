package helper

import (
	"fmt"
	"strconv"
)

func FormatThousand(nominal int32) string {
	nominalInString := strconv.Itoa(int(nominal))
	nominalLen := len(nominalInString)
	nominalThousandFormatted := ""
	firstDigit := nominalLen % 3

	for i := 1; i <= nominalLen; i++ {
		nominalThousandFormatted += string(nominalInString[i-1])
		if nominalLen > 2 && i == int(firstDigit) || ((i-firstDigit)%3 == 0 && i > 2 && i != nominalLen) {
			nominalThousandFormatted += "."
		}
	}

	return fmt.Sprintf("Rp. %s", nominalThousandFormatted)
}
