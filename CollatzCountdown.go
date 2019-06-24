package piscine


func CollatzCountdown(start int) int {
var result=0
	for i := start; i > 1; i-- {
	
			if i%2==0{
				i=i/2
			}else{
				i=3*i+1
			}
	}

return i
}