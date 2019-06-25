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

func CloseDoor(door *Door) bool {
	PrintStr("Door Closing...")
	door.state = CLOSE
	return true
}
func CloseDoor(door *Door) bool {
	PrintStr("Door Closing...")
	door.state = OPEN
	return false
}

func IsDoorOpen(door Door) {
	PrintStr("is the Door opened ?")
	return door.state = OPEN
}

func IsDoorClose(door *Door) bool {
	PrintStr("is the Door closed ?")
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