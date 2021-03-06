package main

import (
	"fmt"
	"time"
)

const (
	red   = false
	black = true
)

func isBlack(node *Node) bool {
	if node == nil {
		return true
	}
	return node.color
}

type Comparator func(key1, key2 interface{}) int8

type RbTree struct {
	root       *Node
	size       uint32
	comparator Comparator
}

type Node struct {
	Key    interface{}
	Value  interface{}
	color  bool
	left   *Node
	right  *Node
	parent *Node
}

func NewRbTree(comparator Comparator) *RbTree {
	return &RbTree{
		comparator: comparator,
	}
}

func (rb *RbTree) Set(key, value interface{}) {
	if rb.root == nil {
		//根节点是黑色
		rb.root = &Node{
			Key:   key,
			Value: value,
			color: black,
		}
		rb.size++
		return
	}

	var (
		node    *Node
		current = rb.root
	)

	for {
		switch rb.comparator(key, current.Key) {
		case 0:
			current.Key = key
			current.Value = value
			return
		case 1:
			if current.right == nil {
				current.right = &Node{parent: current, Key: key, Value: value, color: red}
				node = current.right
				rb.size++
				rb.insert(node)
				return
			}
			current = current.right
		case -1:
			if current.left == nil {
				current.left = &Node{parent: current, Key: key, Value: value, color: red}
				node = current.left

				rb.size++
				rb.insert(node)
				return
			}
			current = current.left
		}
	}
}

func (rb *RbTree) insert(node *Node) {
	if isBlack(node.parent) {
		return
	}

	//父节点为红色 -> 祖父节点一定为黑色，且一定存在(红色节点的孩子必须为黑色，根节点必须为黑色)
	var (
		uncle *Node
	)

	//循环判断父节点颜色，知道父节点为黑色的时候退出
	for !isBlack(node.parent) {
		//如果父节点为左节点
		if node.parent == node.parent.parent.left {
			//叔叔节点
			uncle = node.parent.parent.right
			if !isBlack(uncle) { //叔叔节点为红色
				node.parent.color = black
				uncle.color = black
				node.parent.parent.color = red
				node = node.parent.parent
			} else { //叔叔节点为黑色
				//当前节点在右，需要将当前节点变到左边
				if node == node.parent.right {
					node = node.parent
					rb.leftRotate(node)
				}

				//当前节点在左
				node.parent.color = black
				node.parent.parent.color = red
				rb.rightRotate(node.parent.parent)
			}
		} else { //父节点为右节点
			uncle = node.parent.parent.left
			if !isBlack(uncle) { //叔叔节点为红色
				node.parent.color = black
				uncle.color = black
				node.parent.parent.color = red
				node = node.parent.parent
			} else { //叔叔节点为黑色
				//当前节点在左，需要将当前节点变到右面
				if node == node.parent.left {
					node = node.parent
					rb.rightRotate(node)
				}
				//当前节点在右
				node.parent.color = black
				node.parent.parent.color = red
				rb.leftRotate(node.parent.parent)
			}
		}
	}

	rb.root.color = black
}

func (rb *RbTree) leftRotate(node *Node) {
	rightNode := node.right
	//旋转节点的右孩子变为右节点左孩子
	node.right = rightNode.left
	if rightNode.left != nil {
		//右节点左孩子的父亲变为旋转节点
		rightNode.left.parent = node
	}

	//右节点父亲变为旋转节点的父亲
	rightNode.parent = node.parent

	//根节点左旋
	if node.parent == nil {
		rb.root = rightNode
	} else if node == node.parent.left { //左节点左旋
		//旋转节点父亲的左孩子变为右节点
		node.parent.left = rightNode
	} else { //右节点左旋
		//旋转节点父亲的右孩子变为右节点
		node.parent.right = rightNode
	}
	//右节点的左孩子变为旋转节点
	rightNode.left = node
	//旋转节点的父亲变为右节点
	node.parent = rightNode
}

func (rb *RbTree) rightRotate(node *Node) {
	leftNode := node.left
	//旋转节点左孩子变为左节点的右孩子
	node.left = leftNode.right
	if leftNode.right != nil {
		//左节点右孩子的父亲变为旋转节点
		leftNode.right.parent = node
	}

	//左节点父亲变为旋转节点父亲
	leftNode.parent = node.parent

	//根节点右旋
	if node.parent == nil {
		rb.root = leftNode
	} else if node == node.parent.right { //右节点右旋
		//旋转节点父亲的右节点变为左节点
		node.parent.right = leftNode
	} else { //左节点右旋
		//旋转节点父亲的左孩子变为左节点
		node.parent.left = leftNode
	}
	//左节点右孩子变为旋转节点
	leftNode.right = node
	//旋转节点父亲变为左节点
	node.parent = leftNode
}

func (rb *RbTree) Get(key interface{}) (value interface{}, exists bool) {
	if node, ok := rb.GetNode(key); ok {
		return node.Value, ok
	}
	return
}

func (rb *RbTree) GetNode(key interface{}) (node *Node, exists bool) {
	var (
		current = rb.root
	)

	for {
		if current == nil {
			return nil, false
		}
		switch rb.comparator(key, current.Key) {
		case 0:
			return current, true
		case -1:
			current = current.left
		case 1:
			current = current.right
		}
	}
}

//后继节点
func (rb *RbTree) successor(node *Node) (successor *Node) {
	//如果右节点存在
	if node.right != nil {
		//先假设右节点为后继节点
		successor = node.right
		//循环查找右节点的左节点，直到不存在左节点为止
		for successor.left != nil {
			successor = successor.left
		}
		return successor
	}

	var (
		current = node
	)
	//右节点不存在，先假设父节点为后继节点
	successor = current.parent
	//循环查找当前节点的父节点，直到当前节点不为父节点的右节点为止
	for successor != nil && current == successor.right {
		current = successor
		successor = current.parent
	}
	return successor
}

//前驱节点
func (rb *RbTree) presuccessor(node *Node) (presuccessor *Node) {
	//如果左节点存在
	if node.left != nil {
		//先假设左节点为前驱节点
		presuccessor = node.left
		//循环查找左节点的右节点，直到不存在右节点为止
		for presuccessor.right != nil {
			presuccessor = presuccessor.right
		}
		return presuccessor
	}

	//左节点不存在
	//父节点存在
	if node.parent != nil {
		//当前节点为父节点的右节点
		if node.parent.right == node {
			return node.parent
		}

		//当前节点为父节点左节点
		var (
			current = node
		)

		//假设前驱节点为当前节点父节点
		presuccessor = current.parent
		//循环查找，直到当前节点不为父节点的左节点为止
		for presuccessor != nil && presuccessor.left == current {
			current = presuccessor
			presuccessor = current.parent
		}
		return presuccessor
	}

	return presuccessor
}

func (rb *RbTree) Del(key interface{}) {
	_, exists := rb.GetNode(key)
	if !exists {
		return
	}

	rb.size--

}

func (rb *RbTree) Len() (length uint32) {
	return rb.size
}

//非递归中序遍历
func (rb *RbTree) Range(handler func(key, value interface{}) (handled bool)) {
	if rb.root == nil {
		return
	}

	var (
		min = rb.root
	)

	for min.left != nil {
		min = min.left
	}

	if handler(min.Key, min.Value) {
		return
	}

	for {
		node := rb.successor(min)
		min = node
		if node == nil {
			return
		}

		if handler(node.Key, node.Value) {
			return
		}
	}
}

//非递归倒序遍历
func (rb *RbTree) Desc(handler func(key, value interface{}) (handled bool)) {
	if rb.root == nil {
		return
	}

	var (
		max = rb.root
	)

	for max.right != nil {
		max = max.right
	}

	if handler(max.Key, max.Value) {
		return
	}

	for {
		node := rb.presuccessor(max)
		max = node
		if node == nil {
			return
		}

		if handler(node.Key, node.Value) {
			return
		}
	}
}

func main() {
	rbTree := NewRbTree(func(key1, key2 interface{}) int8 {
		diff := key1.(int64) - key2.(int64)
		switch {
		case diff > 0:
			return 1
		case diff < 0:
			return -1
		}
		return 0
	})

	var (
		index int64 = 1
		max   int64 = 1000
	)

	for ; index < max; index++ {
		rbTree.Set(index, time.Now().UnixNano())
	}

	rbTree.Desc(func(key, value interface{}) (handled bool) {
		fmt.Println(key, value)
		return
	})
}
