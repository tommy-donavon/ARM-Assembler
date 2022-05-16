package assembler

import (
	"fmt"
	"strings"
)

var (
	conditions = map[string]string{
		"EQ": "0000",
		"NE": "0001",
		"CS": "0010",
		"CC": "0011",
		"MI": "0100",
		"PL": "0101",
		"VS": "0110",
		"VC": "0111",
		"HI": "1000",
		"LS": "1001",
		"GE": "1010",
		"LT": "1011",
		"GT": "1100",
		"LE": "1101",
		"AL": "1110",
	}
)

func convertConditionCode(code string) (string, error) {
	key := strings.ToUpper(code)
	if val, ok := conditions[key]; ok {
		return val, nil
	}
	return "", fmt.Errorf("%s not valid condition code", code)
}
