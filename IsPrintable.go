package piscine

import "strings"


const lettre ="ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const lettre1="abcdefghijklmnopqrstuvwxyz"

func IsPrintable(str string) bool {
	var booleen bool


   for _, char := range str {  
      if !strings.Contains(lettre, string(char)) {

      		booleen=false
         
      }else if !strings.Contains(lettre1, string(char)){

      		booleen=false

      }else if !strings.Contains(rune(92), string(char)){

      	booleen=false
      }else{

      		booleen=true
      }
   }

   if booleen{

   	return false
   }else{

   	return true
   }
}