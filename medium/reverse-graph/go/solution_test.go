package solution

import (
	"reflect"
	"testing"
)

func TestRandomGraph(t *testing.T) {
	graph := makeRandomDirectGraph(5)
	t.Logf("Random graph: %v", graph)
}

func TestReverseGraph(t *testing.T) {
	graph := map[int]map[int]bool{
		0: {
			1: true,
			2: true,
		},
		1: {
			0: true,
		},
		2: {
			1: true,
		},
	}

	reversedGraph := ReverseDirectGraph(graph)
	expectedReversedGraph := map[int]map[int]bool{
		0: {
			1: true,
		},
		1: {
			0: true,
			2: true,
		},
		2: {
			0: true,
		},
	}

	isEqual := reflect.DeepEqual(reversedGraph, expectedReversedGraph)

	if isEqual == false {
		t.Errorf("Invalid output. Expected %v but got %v", expectedReversedGraph, reversedGraph)
	}
}
