package assembler

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
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
			fmt.Println(binary)
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
func checkForCondition(s string) string {
	re := regexp.MustCompile("^.*[{](.*)[}].*$")
	match := re.FindStringSubmatch(s)
	if len(match) > 1 {
		return match[1]
	}
	return "AL"
}

func checkForRegisters(s string) []string {
	re := regexp.MustCompile("(R[0-9]{1,2})")
	match := re.FindAllStringSubmatch(s, -1)
	outRegisters := []string{}
	for _, m := range match {
		outRegisters = append(outRegisters, m[1])
	}

	return outRegisters
}

func checkForMetaStore(s string) string {
	re := regexp.MustCompile("^.*[(](UW|PW|PU)[)].*$")
	match := re.FindStringSubmatch(s)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func checkForImmediateValue(s string) string {
	re := regexp.MustCompile("(0x[A-Fa-f0-9]+)")
	match := re.FindStringSubmatch(s)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func checkForSubRoutine(s string) string {
	re := regexp.MustCompile("^BL (:[A-Z]+)$")
	match := re.FindStringSubmatch(s)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func checkForSBit(s string) bool {
	re := regexp.MustCompile("^.*(S).*$")
	match := re.FindStringSubmatch(s)
	return len(match) > 1
}

func checkForWriteBack(s string) bool {
	re := regexp.MustCompile("^.*[(](EA|FD|FA)[)].*$")
	match := re.FindStringSubmatch(s)
	return len(match) > 1
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
