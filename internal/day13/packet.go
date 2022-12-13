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
		return p1.compareListSize(p2)
	} else if p1.isNested && !p2.isNested {
		return p1.compareSize(p2.covertToNested())
	} else if !p1.isNested && p2.isNested {
		return p1.covertToNested().compareSize(p2)
	} else {
		return p1.compareValueSize(p2)
	}
}

func (p1 PacketData) compareListSize(p2 PacketData) PacketSizeComparison {
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

func (p1 PacketData) compareValueSize(p2 PacketData) PacketSizeComparison {
	if p1.value > p2.value {
		return Larger
	} else if p1.value < p2.value {
		return Smaller
	} else {
		return NonDeterminate
	}
}

func (p PacketData) covertToNested() PacketData {
	return PacketData{
		value:      0,
		isNested:   true,
		nestedList: []PacketData{p},
	}
}
