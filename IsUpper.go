package piscine

import "strings"

const det ="abcdefghijklmnopqrstuvwxyz"

func IsUpper(str string) bool {
   for _, char := range str {  
      if !strings.Contains(det, string(char)) {
         return true
      }
   }
   return false
}