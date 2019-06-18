package piscine

import "strings"

const num =strings.ToUpper("abcdefghijklmnopqrstuvwxyz") 

func IsLower(str string) bool {
   for _, char := range str {  
      if !strings.Contains(num, string(char)) {
         return false
      }
   }
   return true
}