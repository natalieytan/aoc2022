package day4

import (
	"errors"
	"strconv"
	"strings"
)

type AssignmentGroup struct {
	LowerLimit int
	UpperLimit int
}

type AssignmentPair struct {
	FirstGroup  AssignmentGroup
	SecondGroup AssignmentGroup
}

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
		FirstGroup:  group1,
		SecondGroup: group2,
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
		LowerLimit: lower,
		UpperLimit: upper,
	}, err
}

func TotalFullyContainsPairs(assignmentPairs []AssignmentPair) int {
	total := 0

	for _, assignmentPair := range assignmentPairs {
		if assignmentPair.FullyContains() {
			total += 1
		}
	}

	return total
}

func (a AssignmentPair) FullyContains() bool {
	return a.FirstGroup.FullyContains(a.SecondGroup) || a.SecondGroup.FullyContains(a.FirstGroup)
}

func (a AssignmentGroup) FullyContains(a2 AssignmentGroup) bool {
	return (a.LowerLimit <= a2.LowerLimit) && (a.UpperLimit >= a2.UpperLimit)
}

func TotalOverlapPairs(assignmentPairs []AssignmentPair) int {
	total := 0

	for _, assignmentPair := range assignmentPairs {
		if assignmentPair.Overlaps() {
			total += 1
		}
	}

	return total
}

func (a AssignmentPair) Overlaps() bool {
	return a.FirstGroup.Overlaps(a.SecondGroup) || a.SecondGroup.Overlaps(a.FirstGroup)
}

func (a AssignmentGroup) Overlaps(a2 AssignmentGroup) bool {
	return (a.UpperLimit >= a2.LowerLimit) && (a.LowerLimit <= a2.LowerLimit)
}
