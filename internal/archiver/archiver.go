package archiver

import (
	"archiver/internal/structures"
)

func countFrequency(data []byte) map[byte]int {
	frequency := make(map[byte]int)

	for _, s := range data {
		frequency[s]++
	}

	return frequency
}

func Huffman(data []byte) error {
	frequency := countFrequency(data)
	pq := structures.NewPriorityQueue()

	for symbol, freq := range frequency {
		node := &structures.Node{Frequency: freq, Symbol: symbol}
		pq.Enqueue(node)
	}

	return nil
}
