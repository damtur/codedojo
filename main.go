package main

import "fmt"

//Hi
func main() {
	var helloMsg = "Czesc"
	fmt.Println("Hello" + helloMsg)
}

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

func (n *Node) Find(s string) *Node {
	if n == nil {
		return nil
	}
	if n.value == s {
		return n
	} else {
		return n.next.Find(s)
	}
}
