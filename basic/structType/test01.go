package main

import "fmt"

type treeNode struct {
	value int
	left,right *treeNode
}

//返回了局部变量的地址
func createTreeNode(value int)  *treeNode{
	return &treeNode{value:value}
}
//接收者方法
func (node treeNode) print() {
	fmt.Println(node.value)
}
func (node *treeNode)setValue(value int)  {
	if node == nil{
		fmt.Println("setting value is ignored")
		return
	}
	node.value = value
}
/**
遍历
 */
func (node *treeNode) traverse(){
	if node == nil{
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}
func main() {
	var root treeNode
	root = treeNode{value:3}
	root.left =&treeNode{}
	root.right = &treeNode{value:5}

	root.print()

	root.setValue(5)
	root.print()

}
