package piscine


func AppendRange(min, max int) []int {
var sl []int
	if max>min {
		for i := min-1; i < max-1; i++ {
			sl=append(sl,i+1)
		 }
		 return sl 
	}else{
		return nil
	}

}