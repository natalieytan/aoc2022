package day13

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

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

type Result int

const (
	NonDeterminate Result = iota
	Smaller
	Larger
)

func (p PacketPair) isRightOrder() Result {
	return p.left.isSmaller(p.right)
}

func (p1 PacketData) isSmaller(p2 PacketData) Result {
	if p1.isNested && p2.isNested {
		p1len := len(p1.nestedList)
		p2len := len(p2.nestedList)

		for i := 0; i < p2len; i++ {
			if i > p1len-1 {
				return Smaller
			}
			isP1Smaller := p1.nestedList[i].isSmaller(p2.nestedList[i])
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
		return p1.isSmaller(newP2)
	}

	if !p1.isNested && p2.isNested {
		newP1 := PacketData{
			value:      0,
			isNested:   true,
			nestedList: []PacketData{p1},
		}
		return newP1.isSmaller(p2)
	}

	if p1.value > p2.value {
		return Larger
	} else if p1.value < p2.value {
		return Smaller
	} else {
		return NonDeterminate
	}
}

func PrepareData(input []byte) ([]PacketPair, error) {
	data := string(input)
	rawPacketPairs := strings.Split(data, "\n\n")
	packetPairs := make([]PacketPair, len(rawPacketPairs))
	for i, rawPacketPair := range rawPacketPairs {
		packetPair, err := parsePacketPair(rawPacketPair)
		if err != nil {
			return packetPairs, err
		}
		packetPairs[i] = packetPair
	}

	return packetPairs, nil
}

func PrepareData2(input []byte) ([]PacketData, error) {
	data := string(input)
	rawPackets := strings.Split(data, "\n")
	packets := []PacketData{}
	for _, rawPacket := range rawPackets {
		if rawPacket != "" {
			packet, err := parsePacket(rawPacket)
			if err != nil {
				return packets, err
			}
			packets = append(packets, packet)
		}

	}
	packets = append(packets, newDividerPacket1())

	packets = append(packets, newDividerPacket2())

	return packets, nil
}

func parsePacketPair(rawPacketPair string) (PacketPair, error) {
	packets := strings.Split(rawPacketPair, "\n")
	if len(packets) != 2 {
		return PacketPair{}, errors.New("error parsing packet pair")
	}
	left, _ := parsePacket(packets[0])
	right, _ := parsePacket(packets[1])

	return PacketPair{
		left:  left,
		right: right,
	}, nil
}

func parsePacket(rawPacket string) (PacketData, error) {
	packetHistory := []PacketData{}
	lastStr := ""

	for _, token := range strings.Split(rawPacket, "") {
		if token == "[" {
			packetHistory = append(packetHistory, PacketData{
				value:      0,
				isNested:   true,
				nestedList: []PacketData{},
			})
		} else if token == "]" {
			if lastStr != "" {
				num, err := strconv.Atoi(lastStr)
				if err == nil {
					latestPacket := packetHistory[len(packetHistory)-1]
					latestPacket.isNested = true
					latestPacket.nestedList = append(latestPacket.nestedList, PacketData{
						value:    num,
						isNested: false,
					})
					packetHistory[len(packetHistory)-1] = latestPacket
					lastStr = ""
				}
			}
			if len(packetHistory) > 1 {
				latestPacket := packetHistory[len(packetHistory)-1]
				secondLastPacket := packetHistory[len(packetHistory)-2]
				secondLastPacket.isNested = true
				secondLastPacket.nestedList = append(secondLastPacket.nestedList, latestPacket)
				packetHistory[len(packetHistory)-2] = secondLastPacket
				packetHistory = packetHistory[:len(packetHistory)-1]
			}
		} else if token == "," {
			num, err := strconv.Atoi(lastStr)
			if err == nil {
				latestPacket := packetHistory[len(packetHistory)-1]
				latestPacket.isNested = true
				latestPacket.nestedList = append(latestPacket.nestedList, PacketData{
					value:    num,
					isNested: false,
				})
				packetHistory[len(packetHistory)-1] = latestPacket
				lastStr = ""
			}
		} else {
			lastStr += token
		}
	}
	return packetHistory[0], nil
}

func Part1(packetPairs []PacketPair) int {
	total := 0

	for idx, packetPair := range packetPairs {
		rightorder := packetPair.isRightOrder()
		if rightorder == Smaller {
			total += (idx + 1)
		}
	}

	return total
}

func Part2(packetPairs []PacketData) int {

	answer := 1
	clonedPackets := packetPairs
	sort.Slice(clonedPackets, func(i, j int) bool {
		return packetPairs[i].isSmaller(packetPairs[j]) == Smaller
	})
	for idx, packet := range clonedPackets {
		if packet.isTracer {
			answer *= (idx + 1)
		}
	}
	return answer
}

func newDividerPacket1() PacketData {
	return PacketData{
		isTracer: true,
		isNested: true,
		nestedList: []PacketData{
			{
				isNested: true,
				nestedList: []PacketData{
					{
						value: 2,
					},
				},
			},
		},
	}
}

func newDividerPacket2() PacketData {
	return PacketData{
		isTracer: true,
		isNested: true,
		nestedList: []PacketData{
			{
				isNested: true,
				nestedList: []PacketData{
					{
						value: 6,
					},
				},
			},
		},
	}
}
