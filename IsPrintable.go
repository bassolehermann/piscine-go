package piscine

import "strings"
import "strconv"

const lettre = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"


func IsPrintable(str string) bool {
	
var booleen bool 

	for i := 0; i < len(str); i++ {

		if !strings.Contains(lettre, strconv.QuoteToASCII (str)){

			booleen=true

		}else{

			booleen=false
		}
		
	}

	if booleen{

		return true
	}else{
		return false
	}

	
}