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
	}
	branchInstruction struct {
		conditionCode string
		offset        string
	}
)

func newMOVInstruction(conditionCode, destinationRegister, immediateValue string, isMOVW bool) (*movInstruction, error) {
	immValueBin, err := convertHexTo16BitBinary(immediateValue)
	if err != nil {
		return nil, err
	}
	codCodeBin, err := convertConditionCode(conditionCode)
	if err != nil {
		return nil, err
	}
	destRegBin, err := convertRegisters(destinationRegister)
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

func newStoreInstruction(conditionCode, destinationRegister, baseRegister, offset string, isImmediate, isStore bool) (*storeInstruction, error) {
	offsetBin, err := convertHexTo12BitBinary(offset)
	if err != nil {
		return nil, err
	}
	codCodeBin, err := convertConditionCode(conditionCode)
	if err != nil {
		return nil, err
	}
	destRegBin, err := convertRegisters(destinationRegister)
	if err != nil {
		return nil, err
	}
	baseRegBin, err := convertRegisters(baseRegister)
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
	}, nil
}

func newBranchInstruction(conditionCode, offset string) (*branchInstruction, error) {
	codCodeBin, err := convertConditionCode(conditionCode)
	if err != nil {
		return nil, err
	}
	offsetBin, err := convertHexTo24BitBinary(offset)
	if err != nil {
		return nil, err
	}
	return &branchInstruction{
		conditionCode: codCodeBin,
		offset:        offsetBin,
	}, nil
}

func (b *branchInstruction) toString() string {
	instruction := fmt.Sprintf("%s1010%s", b.conditionCode, b.offset)
	return insertSpaceNth(instruction, 4)
}

func (s *storeInstruction) toString() string {
	iBit := "1"
	lBit := "1"
	if s.isImmediate {
		iBit = "0"
	}
	if s.store {
		lBit = "0"
	}
	instruction := fmt.Sprintf("%s01%s1100%s%s%s%s", s.conditionCode, iBit, lBit, s.baseRegister, s.destinationRegister, s.offset)
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
