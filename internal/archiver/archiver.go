package archiver

import (
	"archiver/internal/structures"
	"fmt"
	"io"
)

func makeCodes(pq *structures.PriorityQueue) map[byte]string {
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

func encode(data []byte, destination io.Writer, codesMap map[byte]string) error {
	var currentByte byte
	var bitCount uint8

	for _, char := range data {
		code := codesMap[char]

		for _, bit := range code {
			currentByte <<= 1
			if bit == '1' {
				currentByte |= 1
			}
			bitCount++

			if bitCount == 8 {
				_, err := destination.Write([]byte{currentByte})
				if err != nil {
					return fmt.Errorf("error while writing data")
				}
				currentByte = 0
				bitCount = 0
			}
		}
	}

	if bitCount > 0 {
		currentByte <<= (8 - bitCount)
		_, err := destination.Write([]byte{currentByte})
		if err != nil {
			return fmt.Errorf("error while writing data")
		}
	}

	return nil
}

func Huffman(dataSource io.Reader, destination io.Writer) error {
	data, err := io.ReadAll(dataSource)
	if err != nil {
		return fmt.Errorf("error while reading dataSource: %w", err)
	}

	frequency := countFrequency(data)
	pq := structures.NewPriorityQueue()

	for symbol, freq := range frequency {
		node := &structures.Node{Frequency: freq, Symbol: symbol}
		pq.Enqueue(node)
	}

	codesMap := makeCodes(pq)

	err = encode(data, destination, codesMap)

	return err
}
