package main

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func listsize(l *List) int {
	return 0 //placeholder
}

func ListPushFront(l *List, data interface{}) {

}

// equivalent of the main test
// func ListSize(l *List) int {
// 	link := &List{}

// 	ListPushFront(link, "Hello")
// 	ListPushFront(link, "2")
// 	ListPushFront(link, "you")
// 	ListPushFront(link, "man")

// 	fmt.Println(listsize(link))
// }
