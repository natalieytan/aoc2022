package day4

type AssignmentPair struct {
	firstGroup  AssignmentGroup
	secondGroup AssignmentGroup
}

func (a AssignmentPair) fullyContains() bool {
	return a.firstGroup.fullyContains(a.secondGroup) || a.secondGroup.fullyContains(a.firstGroup)
}

func (a AssignmentPair) overlaps() bool {
	return a.firstGroup.overlaps(a.secondGroup) || a.secondGroup.overlaps(a.firstGroup)
}
