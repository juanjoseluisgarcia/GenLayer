package services

import (
	"githuhub.com/juanjoseluisgarcia/GenLayer/cmd/03-network-routing/Model"
	"math"
)

type CompressionRunner struct{}

func NewCompressionRunner() *CompressionRunner {
	return &CompressionRunner{}
}

func isCompressionNode(compressionNodes []string, node string) bool {
	for _, n := range compressionNodes {
		if n == node {
			return true
		}
	}
	return false
}

func (*CompressionRunner) FindMinimumLatencyPath(net Model.Network, source string, target string) int {
	distance := make(map[string]int)
	queue := []Model.Node{{Name: source, Latency: 0}}
	net.CompressionAlreadyUsed = false

	for node := range net.Graph {
		distance[node] = math.MaxInt32
	}
	distance[source] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Name == target {
			return current.Latency
		}

		for _, edge := range net.Graph[current.Name] {
			neighbour := edge[0].(string)
			latency := edge[1].(int)

			if isCompressionNode(net.CompressionNodes, current.Name) && !net.CompressionAlreadyUsed {
				compressedLatency := current.Latency + int(math.Ceil(float64(latency)/2))
				if compressedLatency < distance[neighbour] {
					distance[neighbour] = compressedLatency
					net.CompressionAlreadyUsed = true
					queue = append(queue, Model.Node{Name: neighbour, Latency: compressedLatency})
					continue
				}
			}

			newLatency := current.Latency + latency
			if newLatency < distance[neighbour] {
				distance[neighbour] = newLatency
				queue = append(queue, Model.Node{Name: neighbour, Latency: newLatency})
			}
		}
	}

	return distance[target]
}
