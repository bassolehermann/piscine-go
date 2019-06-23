package piscine

import "strings"

const num = "0123456789"

func IsNumeric(str string) bool {
   for _, char := range str {  
      if !strings.Contains(num, strings.ToLower(string(char))) {
         return false
      }
   }
   return true
}