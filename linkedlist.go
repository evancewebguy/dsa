package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (list *LinkedList) AddToHead(property int) {
	var node = Node{}
	node.property = property
	if node.nextNode != nil {
		node.nextNode = list.headNode
	}
	list.headNode = &node
}

// IterateList iterates over LinkedList
func (list *LinkedList) IterateList() {
	var node *Node
	for node = list.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

// LastNode

func (list *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node

	for node = list.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// AddToEnd

func (list *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var lastNode *Node
	lastNode = list.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	}
}

// NodeWithValue

func (list *LinkedList) NodeWithValue(property int) *Node {
	var node *Node
	var nodeWithValue *Node

	for node = list.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWithValue = node
			break
		}
	}
	return nodeWithValue
}

// AddAfter

func (list *LinkedList) AddAfter(property int, value int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	var nodeWithValue *Node
	nodeWithValue = list.NodeWithValue(property)
	if nodeWithValue != nil {
		node.nextNode = nodeWithValue.nextNode
		nodeWithValue.nextNode = node
	}
}

func main() {
	list := LinkedList{}
	list.AddToHead(1)
	list.AddToHead(2)
	list.AddToEnd(5)
	list.AddAfter(1, 7)
	list.IterateList()

}
