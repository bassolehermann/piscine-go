package piscine

func Index(s string, toFind string) int {

	var index int
	var booleen bool

	for i := 0; i < len(s); i++ {

		if string(s[i])== string(toFind[0]){

			for j := 0; j < len(toFind); j++ {
				
				if string(s[i])==string(toFind[0]){

					booleen=true
					index=i
					
				
				}else{
					booleen=false
					break 
				}
				
			}
			i++
		}
		
	}
	if booleen {
		return index
	}else{
		return -1
	}

}
