package archiver

func countFrequency(data []byte) map[byte]int {
	frequency := make(map[byte]int)

	for _, s := range data {
		frequency[s]++
	}

	return frequency
}

func Huffman(data []byte) error {
	frequency := countFrequency(data)

	_ = frequency
	return nil
}
