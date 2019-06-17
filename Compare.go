package piscine

import "strings"

func Compare(s string, toFind string) int {
var booleen bool
var l int 

	if len(s)!= len(toFind){
		 l=len(s)-len(toFind)

		 	for i := 0; i < l-1; i++ {

		 			if s[i] != toFind[i]{
		 					return -1
		 			}else{
							return 1
					}
		 	
			 }

	}else{

		for i := 0; i < len(s); i++ {

			if s[i] == toFind[i]{

				booleen=true

			}else{
				return -1
			}	
	
		}
	}
	
	if booleen{
		return 0
	}else{
	
		return -1
	}

}