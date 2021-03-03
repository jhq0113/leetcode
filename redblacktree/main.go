package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
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
				rb.insert(node, false)
				return
			}
			current = current.right
		case -1:
			if current.left == nil {
				current.left = &Node{parent: current, Key: key, Value: value, color: red}
				node = current.left

				rb.size++
				rb.insert(node, true)
				return
			}
			current = current.left
		}
	}
}

func (rb *RbTree) insert(node *Node, isLeft bool) {
	if isBlack(node.parent) {
		return
	}

	//父节点为红色 -> 祖父节点一定为黑色，且一定存在(红色节点的孩子必须为黑色，根节点必须为黑色)
	var (
		reference *Node
	)

	if isLeft { //插入左节点
		reference = node.parent.parent.right
		if !isBlack(reference) {
			node.parent.color = black
			reference.color = black
			node.parent.parent.color = red
			node = node.parent.parent
		} else {
			if node == node.parent.right {
				node = node.parent
				rb.leftRotate(node)
			}

			node.parent.color = black
			node.parent.parent.color = red
			rb.rightRotate(node.parent.parent)
		}
	} else { //插入右节点
		reference = node.parent.parent.left
		if !isBlack(reference) {
			node.parent.color = black
			reference.color = black
			node.parent.parent.color = red
			node = node.parent.parent
		} else {
			if node == node.parent.left {
				node = node.parent
				rb.rightRotate(node)
			}
			node.parent.color = black
			node.parent.parent.color = red
			rb.leftRotate(node.parent.parent)
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
	var (
		node = rb.root
	)

	for {
		if node == nil {
			return nil, false
		}
		switch rb.comparator(key, node.Key) {
		case 0:
			return node.Value, true
		case -1:
			node = node.left
		case 1:
			node = node.right
		}
	}
}

func (rb *RbTree) Del(key interface{}) {

}

func (rb *RbTree) Len() (length uint32) {
	return rb.size
}

func (rb *RbTree) Range(func(key, value interface{}) (handled bool)) {

}

func main() {
	rbTree := NewRbTree(func(key1, key2 interface{}) int8 {
		s1 := key1.(string)
		s2 := key2.(string)
		min := len(s2)
		if len(s1) < len(s2) {
			min = len(s1)
		}

		diff := 0
		for i := 0; i < min && diff == 0; i++ {
			diff = int(s1[i]) - int(s2[i])
		}

		if diff == 0 {
			diff = len(s1) - len(s2)
		}

		if diff < 0 {
			return -1
		}

		if diff > 0 {
			return 1
		}
		return 0
	})

	rand.Seed(time.Now().UnixNano())

	var (
		max = 1000000
	)

	start := time.Now()
	keys := make([]string, max, max)
	for index := 0; index < max; index++ {
		current := time.Now()
		keys[index] = fmt.Sprintf("%x", md5.Sum([]byte(current.String())))
		rbTree.Set(keys[index], current.UnixNano())
	}

	fmt.Println("set cost:", time.Now().Sub(start))
	start = time.Now()
	for _, key := range keys {
		if _, exist := rbTree.Get(key); !exist {
			fmt.Println("err key:", key)
		}
	}
	fmt.Println("get cost:", time.Now().Sub(start))
}
