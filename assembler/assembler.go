package assembler

import (
	"fmt"
)

type (
	Assembler struct{}

	operationInstruction struct {
		conditionCode       string
		iBit                bool
		operationCode       string
		sBit                bool
		operandRegister     string
		destinationRegister string
		operandTwo          string
	}
	movInstruction struct {
		conditionCode       string
		destinationRegister string
		immediateValue      string
		isMOVW              bool
	}
	storeInstruction struct {
		conditionCode       string
		baseRegister        string
		destinationRegister string
		isImmediate         bool
		offset              string
		store               bool
		writeBack           bool
	}
)

func newMOVInstruction(instruction string, isMOVW bool) (*movInstruction, error) {
	conditionCode := checkForCondition(instruction)
	registers := checkForRegisters(instruction)
	immValue := checkForImmediateValue(instruction)

	if conditionCode == "" || len(registers) != 1 || immValue == "" {
		return nil, fmt.Errorf("%s is an invalid MOV(T/W) operation", instruction)
	}

	immValueBin, err := convertHexTo16BitBinary(immValue)
	if err != nil {
		return nil, err
	}
	codCodeBin, err := convertConditionCode(conditionCode)
	if err != nil {
		return nil, err
	}
	destRegBin, err := convertRegisters(registers[0])
	if err != nil {
		return nil, err
	}
	return &movInstruction{
		conditionCode:       codCodeBin,
		destinationRegister: destRegBin,
		immediateValue:      immValueBin,
		isMOVW:              isMOVW,
	}, nil
}

func newStoreInstruction(instruction string, isImmediate, isStore bool) (*storeInstruction, error) {
	offset := checkForImmediateValue(instruction)
	if offset == "" {
		offset = "0x0"
	}
	conditionCode := checkForCondition(instruction)
	registers := checkForRegisters(instruction)
	writeBack := checkForWriteBack(instruction)
	if len(registers) != 2 {
		return nil, fmt.Errorf("%s is not a valid Load/Store command", instruction)
	}

	offsetBin, err := convertHexTo12BitBinary(offset)
	if err != nil {
		return nil, err
	}
	codCodeBin, err := convertConditionCode(conditionCode)
	if err != nil {
		return nil, err
	}
	destRegBin, err := convertRegisters(registers[0])
	if err != nil {
		return nil, err
	}
	baseRegBin, err := convertRegisters(registers[1])
	if err != nil {
		return nil, err
	}

	return &storeInstruction{
		conditionCode:       codCodeBin,
		baseRegister:        baseRegBin,
		destinationRegister: destRegBin,
		isImmediate:         isImmediate,
		offset:              offsetBin,
		store:               isStore,
		writeBack:           writeBack,
	}, nil
}

func newOperationInstruction(opcode, instruction string) (*operationInstruction, error) {
	registers := checkForRegisters(instruction)
	conditionCode := checkForCondition(instruction)
	sBit := checkForSBit(instruction)
	immValue := checkForImmediateValue(instruction)
	iBit := false

	destinationReg, err := convertRegisters(registers[0])
	if err != nil {
		return nil, err
	}
	rnRegester, err := convertRegisters(registers[1])
	if err != nil {
		return nil, err
	}
	condCode, err := convertConditionCode(conditionCode)
	if err != nil {
		return nil, err
	}
	if immValue == "" && len(registers) == 3 {
		immValue = registers[2]
		immValue, err = convertRegisters(immValue)
		if err != nil {
			return nil, err
		}
	} else {
		iBit = true
		immValue, err = convertHexTo12BitBinary(immValue)
		if err != nil {
			return nil, err
		}
	}
	return &operationInstruction{
		conditionCode:       condCode,
		iBit:                iBit,
		operationCode:       opcode,
		sBit:                sBit,
		operandRegister:     rnRegester,
		destinationRegister: destinationReg,
		operandTwo:          immValue,
	}, nil
}

func (s *storeInstruction) toString() string {
	iBit := "1"
	lBit := "1"
	pBit := "1"
	uBit := "0"
	wBit := "0"
	if s.isImmediate {
		iBit = "0"
	}
	if s.store {
		lBit = "0"
		uBit = "1"
		pBit = "0"
	}
	if s.writeBack {
		wBit = "1"
	}

	instruction := fmt.Sprintf("%s01%s%s%s0%s%s%s%s%s", s.conditionCode, iBit, pBit, uBit, wBit, lBit, s.baseRegister, s.destinationRegister, s.offset)
	return insertSpaceNth(instruction, 4)
}

func (op *operationInstruction) toString() string {
	iBitC := "0"
	sBitC := "0"
	if op.iBit {
		iBitC = "1"
	}
	if op.sBit {
		sBitC = "1"
	}
	instruction := fmt.Sprintf("%s00%s%s%s%s%s%s", op.conditionCode, iBitC, op.operationCode, sBitC, op.operandRegister, op.destinationRegister, op.operandTwo)
	return insertSpaceNth(instruction, 4)
}

func (m *movInstruction) toString() string {
	imm4 := m.immediateValue[0:4]
	imm12 := m.immediateValue[4:]
	instruction := ""
	if m.isMOVW {
		instruction = fmt.Sprintf("%s00110000%s%s%s", m.conditionCode, imm4, m.destinationRegister, imm12)
	} else {
		instruction = fmt.Sprintf("%s00110100%s%s%s", m.conditionCode, imm4, m.destinationRegister, imm12)
	}
	return insertSpaceNth(instruction, 4)
}
