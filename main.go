package main

import "fmt"

//Hi
func main() {
	var helloMsg = "Hello World"
	fmt.Println("Hello" + helloMsg)
}

func returnTrue() bool {
	return true
}

// type Value struct {
// 	value string
// }

type Node struct {
	value string
	next  *Node
}

func (n *Node) Array() []string {
	result := []string{}
	for n != nil {
		result = append(result, n.value)
		n = n.next
	}
	return result
}

func (n *Node) Add(s string) *Node {
	newNode := &Node{
		value: s,
		next:  n,
	}
	return newNode
}

func (n *Node) Remove(toRemove string) *Node {
	var prev *Node

	var head = n

	for n != nil {
		if n.value == toRemove {
			if prev != nil {
				prev.next = n.next
				return head
			}
			return n.next
		}
		prev = n
		n = n.next
	}
	return head
}

// singly-linked list of nodes
// metadata
// e.g. link to next node
// a string payload

// Metadata can contain anything that you need to make the list work
// data contains only the string payload

// A nod can be added to the end

// ADD
// REMOVED by payload
// RETURN ARRAY (slice)
// Searched by payload
