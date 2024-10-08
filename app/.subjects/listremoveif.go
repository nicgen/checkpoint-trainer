package main

import "fmt"

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func listremoveif(l *List, data_ref interface{}) {
	// write your code here
}

func ListPushBack(l *List, data interface{}) {
	// write your code here
}

func PrintList(l *List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

func ListRemoveIf() {

	link := &List{}
	link2 := &List{}

	fmt.Println("----normal state----")
	ListPushBack(link2, 1)
	PrintList(link2)
	listremoveif(link2, 1)
	fmt.Println("------answer-----")
	PrintList(link2)
	fmt.Println()

	fmt.Println("----normal state----")
	ListPushBack(link, 1)
	ListPushBack(link, "Hello")
	ListPushBack(link, 1)
	ListPushBack(link, "There")
	ListPushBack(link, 1)
	ListPushBack(link, 1)
	ListPushBack(link, 1)
	ListPushBack(link, 1)
	ListPushBack(link, 1)
	ListPushBack(link, 1)
	ListPushBack(link, 1)
	ListPushBack(link, "How")
	ListPushBack(link, 1)
	ListPushBack(link, "are")
	ListPushBack(link, "you")
	ListPushBack(link, 1)
	PrintList(link)

	listremoveif(link, 1)
	fmt.Println("------answer-----")
	PrintList(link)
}
