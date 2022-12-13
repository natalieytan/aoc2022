package day13

type PacketPair struct {
	left  PacketData
	right PacketData
}

type PacketData struct {
	value      int
	isNested   bool
	nestedList []PacketData
	isTracer   bool
}

type PacketSizeComparison int

const (
	NonDeterminate PacketSizeComparison = iota
	Smaller
	Larger
)

func (p PacketPair) isRightOrder() bool {
	return p.left.compareSize(p.right) == Smaller
}

func (p1 PacketData) compareSize(p2 PacketData) PacketSizeComparison {
	if p1.isNested && p2.isNested {
		p1len := len(p1.nestedList)
		p2len := len(p2.nestedList)

		for i := 0; i < p2len; i++ {
			if i > p1len-1 {
				return Smaller
			}
			isP1Smaller := p1.nestedList[i].compareSize(p2.nestedList[i])
			if isP1Smaller != NonDeterminate {
				return isP1Smaller
			}

		}
		if p1len > p2len {
			return Larger
		}
		return NonDeterminate
	}

	if p1.isNested && !p2.isNested {
		newP2 := PacketData{
			value:      0,
			isNested:   true,
			nestedList: []PacketData{p2},
		}
		return p1.compareSize(newP2)
	}

	if !p1.isNested && p2.isNested {
		newP1 := PacketData{
			value:      0,
			isNested:   true,
			nestedList: []PacketData{p1},
		}
		return newP1.compareSize(p2)
	}

	if p1.value > p2.value {
		return Larger
	} else if p1.value < p2.value {
		return Smaller
	} else {
		return NonDeterminate
	}
}
