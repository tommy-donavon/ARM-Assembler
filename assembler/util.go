package assembler

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func convertHexTo12BitBinary(hex string) (string, error) {
	hex = hex[2:]
	i, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%012b", i), nil
}
func convertHexTo16BitBinary(hex string) (string, error) {
	hex = hex[2:]
	i, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%016b", i), nil
}
func convertHexTo24BitBinary(hex string) (string, error) {
	hex = hex[2:]
	i, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%024b", i), nil
}

func convertBinaryToHex(binary string) (string, error) {
	bytes := strings.Split(binary, " ")
	hexCodes := []string{}
	for _, val := range bytes {
		ui, err := strconv.ParseUint(val, 2, 64)
		if err != nil {
			return "", err
		}
		hexCodes = append(hexCodes, fmt.Sprintf("%x", ui))
	}
	hex := insertSpaceNth(strings.Join(hexCodes, ""), 2)
	hexJoined := strings.Split(hex, " ")
	final := []string{}
	for i := len(hexJoined) - 1; i >= 0; i-- {
		final = append(final, hexJoined[i])
	}
	return strings.Join(final, ""), nil

}

func insertSpaceNth(s string, n int) string {
	var buffer bytes.Buffer
	n_1 := n - 1
	l_l := len(s) - 1
	for i, rune := range s {
		buffer.WriteRune(rune)
		if i%n == n_1 && i != l_l {
			buffer.WriteRune(' ')
		}
	}
	return buffer.String()
}
