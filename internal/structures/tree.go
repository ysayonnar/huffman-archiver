package structures

type Node struct {
	Symbol    byte
	Frequency int
	Left      *Node
	Right     *Node
}
