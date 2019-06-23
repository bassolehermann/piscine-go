package piscine


func MakeRange(min, max int) []int {
var t int
t=min-1
	if max>min {

		res:=make([]int,max-min)

		for i := 0; i < max-min ; i++ {
			t=t+1
			res[i]=t
		 }
		 return res 
	}else{
		return nil
	}

}