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

type List interface {
	Array() []string
	Add(string)
}

type list struct {
	value *string
	next  *List
}

func NewList() List {
	return &list{
		value: nil,
		next:  nil,
	}
}

func (n *list) Array() []string {
	if n.value == nil {
		return []string{}
	}
	return []string{*n.value}
}

func (n *list) Add(toAdd string) list {
	n.value = &toAdd
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
