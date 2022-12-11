package day11

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func PrepareData(input []byte) ([]*Monkey, error) {
	data := string(input)
	rawMonkeys := strings.Split(data, "\n\n")
	monkeys := make([]*Monkey, len(rawMonkeys))
	for i, rawMonkey := range rawMonkeys {
		operation, err := parseMonkey(rawMonkey)
		if err != nil {
			return monkeys, err
		}
		monkeys[i] = operation
	}

	return monkeys, nil
}

func parseMonkey(rawMonkey string) (*Monkey, error) {
	monkeyExp := regexp.MustCompile(`Monkey (\d+):\n  Starting items: (.+)\n. Operation: new = (.+)\n  Test: divisible by (.+)\n.   If true: throw to monkey (\d+)\n    If false: throw to monkey (\d+)`)
	match := monkeyExp.FindStringSubmatch(rawMonkey)
	if len(match) != 7 {
		return &Monkey{}, errors.New("error parsing monkey data")
	}

	num, _ := strconv.Atoi(match[1])
	startingItemsRaw := strings.Split(match[2], ", ")
	startingItems := make([]int, len(startingItemsRaw))
	for i, itemRaw := range startingItemsRaw {
		item, _ := strconv.Atoi(itemRaw)
		startingItems[i] = item
	}
	testDivisibleBy, _ := strconv.Atoi(match[4])
	trueThrow, _ := strconv.Atoi(match[5])
	falseThrow, _ := strconv.Atoi(match[6])

	operation, err := parseOperation(match[3])
	if err != nil {
		return &Monkey{}, err
	}

	return &Monkey{
		number: num,
		items:  startingItems,
		decision: MonkeyDecision{
			divisbleBy:            testDivisibleBy,
			throwToIfDivisible:    trueThrow,
			throwToIfNotDivisible: falseThrow,
		},
		operation: operation,
	}, nil
}

func parseOperation(rawOperation string) (Operation, error) {
	operations := strings.Split(rawOperation, " ")
	if len(operations) != 3 {
		return Operation{}, errors.New("error parsing operation")
	}

	operationType, err := parseOperationType(operations[1])
	if err != nil {
		return Operation{}, err
	}
	val1, val1Old := parseOperationValue(operations[0])
	val2, val2Old := parseOperationValue(operations[2])

	return Operation{
		operationType: operationType,
		value1IsOld:   val1Old,
		value1:        val1,
		value2IsOld:   val2Old,
		value2:        val2,
	}, nil
}

func parseOperationValue(operationValue string) (int, bool) {
	if operationValue == "old" {
		return 0, true
	}
	i, _ := strconv.Atoi(operationValue)
	return i, false
}

func parseOperationType(operation string) (OperationType, error) {
	switch operation {
	case "+":
		return Add, nil
	case "-":
		return Subtract, nil
	case "*":
		return Multiply, nil
	case "/":
		return Divide, nil
	default:
		return UnknownOperation, errors.New("unknown operation")
	}
}
