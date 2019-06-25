package main
import "fmt"


type door struct{
	state bool

}

func PrintStr(str string) {
	arrayRune := []rune(str)
	for _, s := range arrayRune {
		fmt.Println(s)
	}
}
func OpenDoor(ptrDoor *Door){
	PrintStr("Door Opening...")
	door.state = OPEN

}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	door.state = CLOSE

}

func IsDoorOpen(Door *Door) bool{
	PrintStr("is the Door opened ?")
	
	return true
}

func IsDoorClose(Door *Door) bool {
	PrintStr("is the Door closed ?")
	return false
}


func main() {
	door := &Door{}

	
	if IsDoorClose(door) {
		CloseDoor(door)
	}else if IsDoorOpen(door) {
		OpenDoor(door)
	}else if door.state == OPEN {
		OpenDoor(door)
	}else if door.state == CLOSE {
		CloseDoor(door)
	}
}