package main
import "fmt"


type Door struct{


}

func PrintStr(str string) {
	fmt.Println(str)
}

func CloseDoor(door *Door) bool {
	PrintStr("Door Closing...")
	return false
}
func OpenDoor(door *Door) bool {
	PrintStr("Door Opening...")
	return true
}

func IsDoorOpen(door *Door)  {
	PrintStr("is the Door opened ?")
	
}

func IsDoorClose(door *Door) {
	PrintStr("is the Door closed ?")	
}

func main() {

	door := &Door{}

	
	//if  OpenDoor(door){
	//	IsDoorClose(door)
	//}else 
	if CloseDoor(door) {
		IsDoorOpen(door)
	}
	
}