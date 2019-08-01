package main

import (
	"fmt"
	"imooc/learngo/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder()  {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()

}

func main() {
	var root tree.Node
	//fmt.Println(root)

	root = tree.Node{Value: 3}
	//fmt.Println(root)

	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}

	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.Set(4)

	root.Traverse()
	//fmt.Println()
	//mynode := myTreeNode{&root}
	//mynode.postOrder()



	//root.right.left.print()


	//node := []treeNode{
	//	{value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(node)
	//
	//fmt.Println(root)
	//root.print()
	//root.set(100)
	//fmt.Println(root)

	//pRoot := &root
	//fmt.Println(pRoot)
	//
	//
	//pRoot.print()
	//pRoot.set(200)
	//fmt.Println(*pRoot)
	//
	//var nRoot *tree.Node
	//fmt.Println("999999")
	//fmt.Println(nRoot)
	//nRoot.Set(300)
	//fmt.Println("000000")
	//fmt.Println(nRoot)
	//
	//fmt.Println("222222")
	//nRoot = &root
	//fmt.Println(nRoot)
	//fmt.Println("111111")
	//nRoot.Set(400)
	//fmt.Println(nRoot)

	//root.Traverse()

	fmt.Println()
	fmt.Println("----count----")

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
		//fmt.Print(node.Value, " ")
	})

	fmt.Println("Node count: ", nodeCount)

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if maxNode < node.Value {
			maxNode = node.Value
		}
	}

	fmt.Println("max node value: ", maxNode)
}
