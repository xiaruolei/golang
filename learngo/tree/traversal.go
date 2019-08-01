package tree

import "fmt"

func (node *Node) Traverse()  {
	//if node == nil {
	//	return
	//}
	//node.Left.Traverse()
	//node.Print()
	//node.Right.Traverse()

	node.TraverseFunc(func(node *Node) {
		node.Print()
	})
}

func (node *Node) TraverseFunc(f func(*Node))  {
	if node == nil {
		return
	}

	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	fmt.Println("func again")
	go func() {

		fmt.Println("=================")
		node.TraverseFunc(func(node *Node) {
			out <- node
			fmt.Println(node)
		})
		close(out)
	}()

	return out
}