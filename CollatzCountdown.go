package piscine


func CollatzCountdown(n int) int {


var nb = 1
 for n != 1 {
  	 if n%2 == 0{
 			n=n/2
 	 }else{
 			n=3*n+1
 	 }
  nb++
 }
 return nb

}

