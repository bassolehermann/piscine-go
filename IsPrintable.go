package piscine

import "strings"
import "strconv"

func IsPrintable(str string) bool {
const lettre = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"	
var booleen bool 

	for i := 0; i < len(str); i++ {

		if strings.Contains(lettre, strconv.QuoteToASCII (str)){

			booleen=false
			
		}else if !strings.Contains(lettre, strconv.QuoteToASCII (str)) && !strings.Contains(strconv.QuoteToASCII (str), "\\"){

			booleen=true
		}
		
	}

	if booleen{

		return true
	}else{
		return false
	}

	
}