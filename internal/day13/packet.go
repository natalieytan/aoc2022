package day13

type PacketPair struct {
	left  Packet
	right Packet
}

func (p PacketPair) isRightOrder() bool {
	return p.left.compareSize(p.right) == Smaller
}

type Packet interface {
	compareSize(Packet) PacketSizeComparison
	isTracer() bool
}

type PacketValue struct {
	value int
}

type PacketList struct {
	list   []Packet
	tracer bool
}

type PacketSizeComparison int

const (
	NonDeterminate PacketSizeComparison = iota
	Smaller
	Larger
)

func (p1 PacketList) compareSize(p Packet) PacketSizeComparison {
	switch p2 := p.(type) {
	case PacketList:
		p1len := len(p1.list)
		p2len := len(p2.list)

		for i := 0; i < p2len; i++ {
			if i > p1len-1 {
				return Smaller
			}
			isP1Smaller := p1.list[i].compareSize(p2.list[i])
			if isP1Smaller != NonDeterminate {
				return isP1Smaller
			}

		}
		if p1len > p2len {
			return Larger
		}
		return NonDeterminate
	case PacketValue:
		return p1.compareSize(p2.convertToList())
	default:
		return NonDeterminate
	}
}

func (p PacketList) isTracer() bool {
	return p.tracer
}

func (p1 PacketValue) compareSize(p Packet) PacketSizeComparison {
	switch p2 := p.(type) {
	case PacketValue:
		if p1.value > p2.value {
			return Larger
		} else if p1.value < p2.value {
			return Smaller
		} else {
			return NonDeterminate
		}
	case PacketList:
		return p1.convertToList().compareSize(p2)
	default:
		return NonDeterminate

	}
}

func (p PacketValue) convertToList() PacketList {
	return PacketList{
		list: []Packet{p},
	}
}

func (p PacketValue) isTracer() bool {
	return false
}
