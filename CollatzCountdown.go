package piscine


func CollatzCountdown(start int) int {
var result=1
	for i := start; i > 1; i-- {
	
			if start%2==0{
				start=start/2

			}else{
				start=start*3+1
			}

			result++
	}

return result
}