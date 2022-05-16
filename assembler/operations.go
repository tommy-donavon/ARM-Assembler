package assembler

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	operations = map[string]string{
		"AND": "0000",
		"EOR": "0001",
		"SUB": "0010",
		"RSB": "0011",
		"ADD": "0100",
		"ADC": "0101",
		"SBC": "0110",
		"RSC": "0111",
		"TST": "1000",
		"CMP": "1010",
		"CMN": "1011",
		"ORR": "1100",
		"MOV": "1101",
		"BIC": "1110",
		"MVN": "1111",
	}
)

func (arm *Assembler) BuildInstruction(armInstructions []string) (string, error) {
	binInstructions := []string{}
	for i, instruction := range armInstructions {
		if instruction != "" && instruction[0] != '#' {
			if instruction[0] == ':' {
				s := strings.Split(instruction, " ")
				instruction = strings.ReplaceAll(instruction, s[0]+" ", "")
				fmt.Println(instruction)
			}
			segments := strings.Split(instruction, " ")
			switch segments[0] {
			case "MOVW":
				m, err := newMOVInstruction(instruction, true)
				if err != nil {
					return "", err
				}
				binInstructions = append(binInstructions, m.toString())
			case "MOVT":

				m, err := newMOVInstruction(instruction, false)
				if err != nil {
					return "", err
				}
				binInstructions = append(binInstructions, m.toString())
			case "LDR":
				s, err := newStoreInstruction(instruction, true, false)
				if err != nil {
					return "", err
				}
				binInstructions = append(binInstructions, s.toString())
			case "STR":
				s, err := newStoreInstruction(instruction, true, true)
				if err != nil {
					return "", err
				}
				binInstructions = append(binInstructions, s.toString())
			case "B":
				b, err := newBranchInstruction(instruction)
				if err != nil {
					return "", err
				}
				binInstructions = append(binInstructions, b.toString())
			case "BL":
				subCall := checkForSubRoutine(instruction)
				if subCall == "" {
					return "", fmt.Errorf("%s is not valid sub routine call", instruction)
				}
				for j, val := range armInstructions {
					if val != "" && val[0] != '#' {
						vals := strings.Split(val, " ")
						if vals[0] == subCall {
							steps := (j - i) - 2
							n := int64(steps)
							offset := strconv.FormatInt(n, 2)
							offset = fmt.Sprintf("%024s", offset)
							b, err := newBranchLinkInstruction(instruction, offset)
							if err != nil {
								return "", err
							}
							binInstructions = append(binInstructions, b.toString())
						}
					}
				}
			case "BX":
				b, err := newBranchExchangeInstruction(instruction)
				if err != nil {
					return "", err
				}
				binInstructions = append(binInstructions, b.toString())
			default:
				if val, ok := operations[segments[0]]; ok {
					op, err := newOperationInstruction(val, instruction)
					if err != nil {
						return "", err
					}
					binInstructions = append(binInstructions, op.toString())
				} else if !ok {

					return "", fmt.Errorf("invalid operation code")
				}

			}
		}
	}
	hexCodes := []string{}
	for _, val := range binInstructions {
		hex, err := convertBinaryToHex(val)
		if err != nil {
			return "", err
		}
		hexCodes = append(hexCodes, hex)
	}
	return strings.Join(hexCodes, ""), nil
}
