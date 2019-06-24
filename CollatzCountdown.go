package piscine


func CollatzCountdown(start int) int {
var result=1
	for i := start; i > 1; i-- {
	
			if start%2==0{
				start=start/2
				result++
		
			}else{
				start=start*3+1
				result= result+2
			}

			
	}

return result
}