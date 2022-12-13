package day13

import (
	"sort"
)

func Part1(packetPairs []PacketPair) int {
	sumOfIndicesInRightOrder := 0

	for idx, packetPair := range packetPairs {
		if packetPair.isRightOrder() {
			sumOfIndicesInRightOrder += (idx + 1)
		}
	}

	return sumOfIndicesInRightOrder
}

func Part2(packets []Packet) int {
	decoderKey := 1
	clonedPackets := packets
	sort.Slice(clonedPackets, func(i, j int) bool {
		return packets[i].compareSize(packets[j]) == Smaller
	})
	for idx, packet := range clonedPackets {
		if packet.isTracer() {
			decoderKey *= (idx + 1)
		}
	}
	return decoderKey
}
