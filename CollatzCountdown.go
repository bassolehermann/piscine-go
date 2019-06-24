package piscine


func CollatzCountdown(start int) int {

	for i := start; i > 1; i-- {
		
		start=start-1
	}

return start
}