package piscine 




type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {

		n:=&NodeL{Data:data}

		if l.Head == nil {
		l.Head=n
		} else {
		
			for i:=l.Head; i!=nil ; i=i.Next {

				if i.Next==nil{
					i.Next=n
					break
				}
			}
				
			}

}
