package archiver

import (
	"archiver/internal/structures"
	"io"
)

func MakeCodes(pq *structures.PriorityQueue) map[byte]string {
	for pq.Len() > 1 {
		firstNode := pq.Dequeue()
		secondNode := pq.Dequeue()

		node := &structures.Node{
			Frequency: firstNode.Frequency + secondNode.Frequency,
			Right:     secondNode,
			Left:      firstNode,
		}

		pq.Enqueue(node)
	}

	//построение кодов
	codesMap := make(map[byte]string)
	var dfs func(root *structures.Node, code string)
	dfs = func(root *structures.Node, code string) {
		if root == nil {
			return
		}

		if root.Right == nil && root.Left == nil {
			codesMap[root.Symbol] = code
			return
		}

		dfs(root.Left, code+"0")
		dfs(root.Right, code+"1")
	}

	dfs(pq.Dequeue(), "")

	return codesMap
}

func countFrequency(data []byte) map[byte]int {
	frequency := make(map[byte]int)

	for _, s := range data {
		frequency[s]++
	}

	return frequency
}

func Huffman(data []byte, destination io.Writer) {
	frequency := countFrequency(data)
	pq := structures.NewPriorityQueue()

	for symbol, freq := range frequency {
		node := &structures.Node{Frequency: freq, Symbol: symbol}
		pq.Enqueue(node)
	}

	codesMap := MakeCodes(pq)

	//TODO: дописать кодирование
}
