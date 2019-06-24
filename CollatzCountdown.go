package piscine


func CollatzCountdown(start int) int {
var result=0
	for i := start; i > 1; i-- {
	
			if i%2==0{
			
				result=result/2
			}else{
			
				result=result*3+1
			}
	}

return result
}