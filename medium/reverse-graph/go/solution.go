package solution

import (
	"math/rand"
)

func makeRandomDirectGraph(size int) map[int]map[int]bool {
	var nodes = make(map[int]map[int]bool)

	for nodeNth := 0; nodeNth < size; nodeNth++ {
		node := map[int]bool{}
		numberOfConnections := int(rand.Float32() * float32(size))

		for connection := 1; connection <= numberOfConnections; connection++ {
			randomConnection := int(rand.Float32() * float32(size))
			node[randomConnection] = true
		}

		nodes[nodeNth] = node
	}

	return nodes
}

// ReverseDirectGraph return a copy of the given graph but inverted
func ReverseDirectGraph(graph map[int]map[int]bool) map[int]map[int]bool {
	var result = make(map[int]map[int]bool) // Naive: O(n) in memory because we copy the whole strucutre

	for node, edges := range graph {
		for targetNode := range edges {
			if result[targetNode] == nil {
				result[targetNode] = make(map[int]bool)
			}
			result[targetNode][node] = true
		}
	}
	return result
}
