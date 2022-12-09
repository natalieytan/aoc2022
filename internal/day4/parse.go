package day4

import (
	"errors"
	"strconv"
	"strings"
)

func PrepareData(input []byte) ([]AssignmentPair, error) {
	data := string(input)
	rawPairs := strings.Split(data, "\n")
	assignmentPairs := make([]AssignmentPair, len(rawPairs))
	for i, rawPair := range rawPairs {
		assignmentPair, err := parseAssignmentPair(rawPair)
		if err != nil {
			return assignmentPairs, err
		}
		assignmentPairs[i] = assignmentPair
	}
	return assignmentPairs, nil
}

func parseAssignmentPair(rawPair string) (AssignmentPair, error) {
	rawGroups := strings.Split(rawPair, ",")
	if len(rawGroups) != 2 {
		return AssignmentPair{}, errors.New("bad data")
	}
	group1, err := parseGroup(rawGroups[0])
	if err != nil {
		return AssignmentPair{}, errors.New("bad data")
	}
	group2, err := parseGroup(rawGroups[1])
	if err != nil {
		return AssignmentPair{}, errors.New("bad data")
	}
	return AssignmentPair{
		firstGroup:  group1,
		secondGroup: group2,
	}, err
}

func parseGroup(rawGroup string) (AssignmentGroup, error) {
	rawLimits := strings.Split(rawGroup, "-")
	if len(rawLimits) != 2 {
		return AssignmentGroup{}, errors.New("bad data")
	}
	lower, err := strconv.Atoi(rawLimits[0])
	if err != nil {
		return AssignmentGroup{}, errors.New("bad data")
	}
	upper, err := strconv.Atoi(rawLimits[1])
	if err != nil {
		return AssignmentGroup{}, errors.New("bad data")
	}
	return AssignmentGroup{
		lowerLimit: lower,
		upperLimit: upper,
	}, err
}
