package day4

type AssignmentGroup struct {
	lowerLimit int
	upperLimit int
}

func (a AssignmentGroup) fullyContains(a2 AssignmentGroup) bool {
	return (a.lowerLimit <= a2.lowerLimit) && (a.upperLimit >= a2.upperLimit)
}

func (a AssignmentGroup) overlaps(a2 AssignmentGroup) bool {
	return (a.upperLimit >= a2.lowerLimit) && (a.lowerLimit <= a2.lowerLimit)
}
