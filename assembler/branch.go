package assembler

import "fmt"

type (
	branchInstruction struct {
		conditionCode string
		offset        string
		link          bool
	}
	branchExchangeInstruction struct {
		conditionCode string
	}
)

func newBranchInstruction(instruction string) (*branchInstruction, error) {
	conditionCode := checkForCondition(instruction)
	offset := checkForImmediateValue(instruction)
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
		link:          false,
	}, nil
}
func newBranchLinkInstruction(instruction, offset string) (*branchInstruction, error) {
	conditionCode := checkForCondition(instruction)

	codCodeBin, err := convertConditionCode(conditionCode)
	if err != nil {
		return nil, err
	}
	return &branchInstruction{
		conditionCode: codCodeBin,
		offset:        offset,
		link:          true,
	}, nil
}
func newBranchExchangeInstruction(instruction string) (*branchExchangeInstruction, error) {
	conditionCode := checkForCondition(instruction)
	codCodeBin, err := convertConditionCode(conditionCode)
	if err != nil {
		return nil, err
	}
	return &branchExchangeInstruction{
		conditionCode: codCodeBin,
	}, nil
}

func (b *branchInstruction) toString() string {
	link := "0"
	if b.link {
		link = "1"
	}
	instruction := fmt.Sprintf("%s101%s%s", b.conditionCode, link, b.offset)
	return insertSpaceNth(instruction, 4)
}

func (bx *branchExchangeInstruction) toString() string {
	instruction := fmt.Sprintf("%s0001001011111111111100011110", bx.conditionCode)
	return instruction
}
