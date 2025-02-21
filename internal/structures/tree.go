package structures

type Tree struct {
	Root *Node
}

type Node struct {
	Value byte
	Left  *Node
	Right *Node
}
