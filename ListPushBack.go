package piscine 

import "container/list"


type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
alist := list.New()
alist.PushBack(data)
for e := alist.Front(); e != nil; e = e.Next() {
     fmt.Println(e.Value) // print out the elements
  }
}