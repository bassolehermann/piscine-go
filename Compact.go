package piscine
import "strings"
func Compact(ptr *[]string, length int) int {

tab=make([]string, length)	

for i := 0; i < length; i++ {

	if ptr[i] != nil {
		tab[i]=ptr[i]
	}
	
}

return len(tab)
}


