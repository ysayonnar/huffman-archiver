package archiver

func countFrequency(text string) map[byte]int {
	frequency := make(map[byte]int)

	for i := 0; i < len(text); i++ {
		frequency[text[i]]++
	}

	return frequency
}

func Huffman(text string) error {
	frequency := countFrequency(text)

	_ = frequency
	return nil
}
