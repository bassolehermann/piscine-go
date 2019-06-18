package piscine

import "strings"

const det ="abcdefghijklmnopqrstuvwxyz"

func IsPrintable(str string) bool {
   for _, char := range str {  
      if !strings.Contains(det, string(char)) {
         return false
      }
   }
   return true
}