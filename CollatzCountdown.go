package piscine


func CollatzCountdown(start int) int {

	for i := start; i > 1; i-- {
	
			if i%2==0{
				i=i/2
		
			}else{
				i=i*3+1
			}

			
	}

return i
}