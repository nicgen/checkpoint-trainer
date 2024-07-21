package main

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func listremoveif(l *List, data_ref interface{}) {

}

func ListPushBack(l *List, data interface{}) {

}

// testing area

// func PrintList(l *piscine.List) {
// 	it := l.Head
// 	for it != nil {
// 		fmt.Print(it.Data, " -> ")
// 		it = it.Next
// 	}

// 	fmt.Print(nil, "\n")
// }

// // equivalent of the main test
// func ListRemoveIf() {
// 	link := &List{}
// 	link2 := &List{}

// 	fmt.Println("----normal state----")
// 	ListPushBack(link2, 1)
// 	PrintList(link2)
// 	ListRemoveIf(link2, 1)
// 	fmt.Println("------answer-----")
// 	PrintList(link2)
// 	fmt.Println()

// 	fmt.Println("----normal state----")
// 	ListPushBack(link, 1)
// 	ListPushBack(link, "Hello")
// 	ListPushBack(link, 1)
// 	ListPushBack(link, "There")
// 	ListPushBack(link, 1)
// 	ListPushBack(link, 1)
// 	ListPushBack(link, "How")
// 	ListPushBack(link, 1)
// 	ListPushBack(link, "are")
// 	ListPushBack(link, "you")
// 	ListPushBack(link, 1)
// 	PrintList(link)

// 	ListRemoveIf(link, 1)
// 	fmt.Println("------answer-----")
// 	PrintList(link)
// }
