package day9

func Part1(motions []Motion) int {
	tailVisited := map[string]bool{}
	headCoordinate := NewCoordinate()
	tailCoordinate := NewCoordinate()

	for _, motion := range motions {
		for i := 0; i < motion.Distance; i++ {
			headCoordinate.MoveInDirection(motion.Direction)
			tailCoordinate.MoveToTouchCoordinate(headCoordinate)
			tailVisited[tailCoordinate.String()] = true
		}
	}

	return len(tailVisited)
}

func Part2(motions []Motion) int {
	tailVisited := map[string]bool{}
	headCoordinate := NewCoordinate()
	t1 := NewCoordinate()
	t2 := NewCoordinate()
	t3 := NewCoordinate()
	t4 := NewCoordinate()
	t5 := NewCoordinate()
	t6 := NewCoordinate()
	t7 := NewCoordinate()
	t8 := NewCoordinate()
	t9 := NewCoordinate()

	for _, motion := range motions {
		for i := 0; i < motion.Distance; i++ {
			headCoordinate.MoveInDirection(motion.Direction)
			t1.MoveToTouchCoordinate(headCoordinate)
			t2.MoveToTouchCoordinate(t1)
			t3.MoveToTouchCoordinate(t2)
			t4.MoveToTouchCoordinate(t3)
			t5.MoveToTouchCoordinate(t4)
			t6.MoveToTouchCoordinate(t5)
			t7.MoveToTouchCoordinate(t6)
			t8.MoveToTouchCoordinate(t7)
			t9.MoveToTouchCoordinate(t8)
			tailVisited[t9.String()] = true
		}
	}

	return len(tailVisited)
}
