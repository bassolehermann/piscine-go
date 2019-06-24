package piscine


func CollatzCountdown(start int) int {
var result=1
	for start>1 {
	
			if start%2==0{
				start=start/2
				result++
		
			}else{
				start=start*3+1
				result++
			}
	
	}

return result
}