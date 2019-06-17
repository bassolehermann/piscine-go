package piscine

import "strings"

const alpha = "0123456789"

func IsNumeric(str string) bool {
   for _, char := range str {  
      if !strings.Contains(alpha, strings.ToLower(string(char))) {
         return false
      }
   }
   return true
}