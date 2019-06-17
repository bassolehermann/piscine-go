package piscine

func Compare(s string, toFind string) int {

	var index int
	var booleen bool

	for i := 0; i < len(toFind); i++ {

		if string(s[i])== string(toFind[i]){

			for j := 0; j < len(s); j++ {
				
				if string(s[i])==string(toFind[i]){

					booleen=true
					index=1
					
				
				}else{
					booleen=false
					break 
				}
				if string(s[i])!=string(toFind[0]){
		
				return 0
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