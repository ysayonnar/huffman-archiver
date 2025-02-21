package main

import "archiver/internal/archiver"

var text = "aaabbbbvsdfsdfgjdfslgjdsf" // пока для тестирования будет текст, потом файлы

func main() {
	archiver.Huffman([]byte(text))
}
