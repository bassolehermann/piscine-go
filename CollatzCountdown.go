package piscine


func CollatzCountdown(n int) int {


var nb = 1
 for n != 1 {
  n = syracuse(n)
  nb++
 }
 return nb

}

func syracuse(n int) int{

 if n%2 == 0{
 	return n/2
 }else{
 	return 3*n+1
 }
 
}