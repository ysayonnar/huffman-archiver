package structures

type Tree struct {
	Root *Node
}

type Node struct {
	Symbol    byte
	Frequency int
	Left      *Node
	Right     *Node
}
