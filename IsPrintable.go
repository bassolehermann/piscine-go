package piscine

import "strings"




func IsPrintable(str string) bool {


	byteArray := []byte(str)

	if byteArray>127{

		return false
	}else{
		return true
	}


	
}
