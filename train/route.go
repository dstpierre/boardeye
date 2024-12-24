package train

import (
	"container/heap"
	"fmt"
	"math"
)

func FindShortesPath(graph Graph, start, end string) ([]Station, error) {
	// Initialize distances and previous nodes
	dist := make(map[string]int)
	prev := make(map[string]string)
	for node := range graph {
		dist[node] = math.MaxInt32
		prev[node] = ""
	}
	dist[start] = 0

	// Priority queue to store nodes and their tentative distances
	pq := make(PriorityQueue, 0)
	heap.Push(&pq, &Item{Value: start, Priority: 0})

	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*Item).Value

		for _, neighbor := range graph[curr] {
			if neighbor.Taken {
				continue
			}

			alt := dist[curr] + neighbor.Length
			if alt < dist[neighbor.End] {
				dist[neighbor.End] = alt
				prev[neighbor.End] = curr
				heap.Push(&pq, &Item{Value: neighbor.End, Priority: alt})
			}
		}

		if curr == end {
			break
		}
	}

	// Reconstruct the path
	path := []Station{}
	curr := end

	isStart := true

	for curr != "" {
		//path = append([]Station{{Start: prev[curr], End: curr}}, path...)
		//curr = prev[curr]

		prevNode := prev[curr]
		if isStart {
			isStart = false
			for _, station := range graph[curr] {
				if station.End == prevNode {
					path = append([]Station{station}, path...)
					break
				}
			}
		}
		for _, station := range graph[prevNode] {
			if station.End == curr {
				path = append([]Station{station}, path...)
				break
			}
		}
		curr = prevNode
	}

	if len(path) == 0 {
		return nil, fmt.Errorf("no path found")
	}

	return path, nil
}
