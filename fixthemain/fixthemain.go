package main
import "fmt"


type Door struct{


}

func PrintStr(str string) {
	fmt.Println(str)
}

func CloseDoor(door *Door)  {
	PrintStr("Door Closing...")
	
}
func OpenDoor(door *Door)  {
	PrintStr("Door Closing...")
	
}

func IsDoorOpen(door *Door) bool {
	PrintStr("is the Door opened ?")


	return true
}

func IsDoorClose(door *Door) bool {
	PrintStr("is the Door closed ?")
	
	return true
}

func main() {

	door := &Door{}

	
	if IsDoorClose(door) {
		OpenDoor(door)
	}else if IsDoorOpen(door) {
		CloseDoor(door)
	}
	
}