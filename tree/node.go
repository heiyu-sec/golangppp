package main

import (
	"fmt"
)

type treeNode struct {
	value int
	left,right *treeNode
}

func (node treeNode) print() {
	fmt.Print(node.value," ")
}

func (node *treeNode) setValue(value int){
	if node ==nil{
		fmt.Println("setting value to nil"+" node .Ignored.")
		return
	}
	node.value=value
}

func (node *treeNode) traverse() {
	if node ==nil{
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func  createNode(value int) *treeNode {
	return &treeNode{value:value}
}

func main()  {
	var root treeNode

	root = treeNode{value:3}
	root.left = &treeNode{}
	root.right = &treeNode{5,nil,nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)
	root.print()
	root.right.left.setValue(4)

	root.traverse()
	root.right.left.print()
	fmt.Println()

	root.setValue(100)
	var pRoot *treeNode
	pRoot.setValue(200)
	pRoot = &root
	pRoot.setValue(300)
	pRoot.print()

}