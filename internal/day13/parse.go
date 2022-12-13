package day13

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func PrepareDataPt1(input []byte) ([]PacketPair, error) {
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

func PrepareData2Pt2(input []byte) ([]Packet, error) {
	data := string(input)
	re := regexp.MustCompile("\n+")
	rawPackets := re.Split(data, -1)
	packets := make([]Packet, len(rawPackets)+2)
	for i, rawPacket := range rawPackets {
		packet, err := parsePacket(rawPacket)
		if err != nil {
			return packets, err
		}
		packets[i] = packet

	}
	packets[len(packets)-2] = newDividerPacket(2)
	packets[len(packets)-1] = newDividerPacket(6)

	return packets, nil
}

func parsePacketPair(rawPacketPair string) (PacketPair, error) {
	packets := strings.Split(rawPacketPair, "\n")
	if len(packets) != 2 {
		return PacketPair{}, errors.New("error parsing packet pair")
	}

	left, err := parsePacket(packets[0])
	if err != nil {
		return PacketPair{}, err
	}

	right, err := parsePacket(packets[1])
	if err != nil {
		return PacketPair{}, err
	}

	return PacketPair{
		left:  left,
		right: right,
	}, nil
}

func parsePacket(rawPacket string) (PacketList, error) {
	packetListHistory := []PacketList{}
	lastStr := ""

	for _, token := range strings.Split(rawPacket, "") {
		switch token {
		case "[":
			// "[" represents the start of a list
			// Add a new PacketList to packetListHistory to keep building the list
			packetListHistory = append(packetListHistory, PacketList{})
		case "]", ",":
			// Check if there was a non-empty value to be added to the list preceding "]" or ","
			if lastStr != "" {
				// convert value to integer
				num, err := strconv.Atoi(lastStr)
				if err != nil {
					return PacketList{}, err
				}
				// add value to the most recent PacketList
				packetListHistory[len(packetListHistory)-1].list = append(packetListHistory[len(packetListHistory)-1].list, PacketValue{
					value: num,
				})
				lastStr = ""
			}
			// "]" represents the end of a list - building of all values complete
			// Add the completed PacketList (last in history) to the PacketList list that it is a part of (2nd last in history)
			// Pop the completed PacketList out of the packetListHistory
			if token == "]" && len(packetListHistory) > 1 {
				packetListHistory[len(packetListHistory)-2].list = append(packetListHistory[len(packetListHistory)-2].list, packetListHistory[len(packetListHistory)-1])
				packetListHistory = packetListHistory[:len(packetListHistory)-1]
			}
		default:
			lastStr += token
		}
	}
	return packetListHistory[0], nil
}

func newDividerPacket(value int) PacketList {
	return PacketList{
		tracer: true,
		list: []Packet{
			PacketList{
				list: []Packet{
					PacketValue{
						value: value,
					},
				},
			},
		},
	}
}
