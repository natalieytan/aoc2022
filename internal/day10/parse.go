package day10

import (
	"errors"
	"strconv"
	"strings"
)

func PrepareData(input []byte) ([]Operation, error) {
	data := string(input)
	rawOperations := strings.Split(data, "\n")
	operations := make([]Operation, len(rawOperations))
	for i, rawOperation := range rawOperations {
		operation, err := parseOperation(rawOperation)
		if err != nil {
			return operations, err
		}
		operations[i] = operation
	}

	return operations, nil
}

func parseOperation(rawOperation string) (Operation, error) {
	if rawOperation == "noop" {
		return Operation{
			operationType: Noop,
			value:         0,
		}, nil
	} else {
		operationAndValue := strings.Split(rawOperation, " ")
		if len(operationAndValue) != 2 && operationAndValue[0] != "addx" {
			return Operation{}, errors.New("invalid operation")
		}
		value, err := strconv.Atoi(operationAndValue[1])
		if err != nil {
			return Operation{}, errors.New("invalid operation")
		}
		return Operation{
			operationType: Addx,
			value:         value,
		}, nil
	}
}
