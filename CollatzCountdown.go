package piscine


func CollatzCountdown(start int) int {
var result=0
	for i := start; i > 1; i-- {
	
			if i%2==0{
				i=i/2
				result=result+1
			}else{
				i=3*i+1
				result=result+1
			}
	}

return result
}