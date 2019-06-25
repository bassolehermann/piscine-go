package main
import "fmt"


type Door struct{


}

func PrintStr(str string) {
	arrayRune := []rune(str)
	for _, s := range arrayRune {
		fmt.Println(s)
	}
}

func CloseDoor(door *Door)  {
	PrintStr("Door Closing...")
	
}
func OpenDoor(door *Door)  {
	PrintStr("Door Closing...")
	
}

func IsDoorOpen(door Door) bool {
	PrintStr("is the Door opened ?")

	door.state = OPEN
	return false
}

func IsDoorClose(door *Door) bool {
	PrintStr("is the Door closed ?")
	door.state = CLOSE
	return true
}

func main() {

	door := &Door{}

	
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}