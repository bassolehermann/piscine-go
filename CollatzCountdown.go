package piscine


func CollatzCountdown(start int) int {

	for i := start; i > 1; i-- {
	
			if start%2==0{
				start=start/2
		
			}else{
				start=start*3+1
			}

			
	}

return start
}