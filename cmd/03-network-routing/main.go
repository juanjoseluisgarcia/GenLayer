package main

import (
	"fmt"
	"githuhub.com/juanjoseluisgarcia/GenLayer/cmd/03-network-routing/Model"
	"githuhub.com/juanjoseluisgarcia/GenLayer/cmd/03-network-routing/services"
)

func main() {

	r := services.NewCompressionRunner()
	net := Model.Network{
		Graph: map[string][][2]interface{}{
			"A": {{"B", 10}, {"C", 20}},
			"B": {{"D", 15}},
			"C": {{"D", 30}},
			"D": {},
		},
		CompressionNodes: []string{"B", "C"},
	}

	source := "A"
	destination := "D"

	minLatency := r.FindMinimumLatencyPath(net, source, destination)
	fmt.Printf("Minimum total latency: %d\n", minLatency)
}
