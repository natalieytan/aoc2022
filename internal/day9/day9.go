package day9

func Part1(motions []Motion) int {
	return tailPositionsVisitedWithKnots(motions, 2)
}

func Part2(motions []Motion) int {
	return tailPositionsVisitedWithKnots(motions, 10)
}

func tailPositionsVisitedWithKnots(motions []Motion, knots int) int {
	tailVisited := map[string]bool{}
	rope := make([]*Coordinate, knots)
	for i := range rope {
		rope[i] = NewCoordinate()
	}

	for _, motion := range motions {
		for i := 0; i < motion.Distance; i++ {
			for knotNumber, knot := range rope {
				if knotNumber == 0 {
					knot.MoveInDirection(motion.Direction)
				} else {
					knot.MoveToTouchCoordinate(rope[knotNumber-1])
				}
				if knotNumber == len(rope)-1 {
					tailVisited[knot.String()] = true
				}
			}
		}
	}

	return len(tailVisited)
}
