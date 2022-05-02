package assembler

import (
	"fmt"
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

func (arm *Assembler) BuildInstruction(armInstruction string) (string, error) {
	segs := strings.Split(armInstruction, ", ")
	var first []string
	if len(segs) == 1 {
		first = strings.Split(armInstruction, " ")
	} else {
		first = strings.Split(segs[0], " ")
	}
	opCode := ""
	if len(first[0]) > 3 {
		opCode = first[0][0:3]
	} else {
		opCode = "B"
	}

	switch opCode {
	case "MOV":
		if first[0][3] == byte('W') {
			return arm.buildMOV(opCode, first, segs, true)
		} else if first[0][3] == byte('T') {
			return arm.buildMOV(opCode, first, segs, false)
		}
	case "LDR":
		return arm.buildStore(opCode, first, segs, false)
	case "STR":
		return arm.buildStore(opCode, first, segs, true)
	case "B":
		return arm.buildBranch(first[0][1:], first[1])
	default:
		return arm.buildOperationInstruction(opCode, first, segs)

	}

	return "", fmt.Errorf("something went wrong")
}

func (arm *Assembler) buildBranch(condCode, offset string) (string, error) {
	b, err := newBranchInstruction(condCode, offset)
	if err != nil {
		return "", err
	}
	hex, err := convertBinaryToHex(b.toString())
	return hex, err
}

func (arm *Assembler) buildMOV(opcode string, first, segs []string, isMOVW bool) (string, error) {
	codCode := first[0][len(first[0])-2:]
	destinationReg := first[1]
	immValue := segs[len(segs)-1]
	m, err := newMOVInstruction(codCode, destinationReg, immValue, isMOVW)
	if err != nil {
		return "", err
	}
	hex, err := convertBinaryToHex(m.toString())
	return hex, err
}

func (arm *Assembler) buildStore(opcode string, first, segs []string, isStore bool) (string, error) {
	codCode := first[0][len(first[0])-2:]
	s, err := newStoreInstruction(codCode, first[1], segs[1], "0x0", true, isStore)
	if err != nil {
		return "", err
	}
	hex, err := convertBinaryToHex(s.toString())

	return hex, err
}

func (arm *Assembler) buildOperationInstruction(opCode string, first, segs []string) (string, error) {
	if val, ok := operations[opCode]; ok {
		codCode := first[0][len(first[0])-2:]
		sBit := false
		iBit := false
		destinationReg := ""
		if len(first) == 3 {
			sBit = true
			destinationReg = first[2]
		} else {
			destinationReg = first[1]
		}
		destinationReg, err := convertRegisters(destinationReg)
		if err != nil {
			return "", err
		}
		rnRegester := segs[1]
		rnRegester, err = convertRegisters(rnRegester)
		if err != nil {
			return "", err
		}
		immValue := segs[2]
		condCode, err := convertConditionCode(codCode)
		if err != nil {
			return "", err
		}
		if string(immValue[0]) == "0" {
			iBit = true
			immValue, err = convertHexTo12BitBinary(immValue)
			if err != nil {
				return "", err
			}
		} else {
			immValue, err = convertRegisters(immValue)
			if err != nil {
				return "", err
			}
		}
		opInstruction := &operationInstruction{
			conditionCode:       condCode,
			iBit:                iBit,
			operationCode:       val,
			sBit:                sBit,
			operandRegister:     rnRegester,
			destinationRegister: destinationReg,
			operandTwo:          immValue,
		}
		hex, err := convertBinaryToHex(opInstruction.toString())
		return hex, err
	}
	return "", fmt.Errorf("not a valid operation")
}
