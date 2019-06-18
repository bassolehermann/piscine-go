package piscine

import "strings"

const let ="abcdefghijklmnopqrstuvwxyz"

func IsUpper(str string) bool {
   for _, char := range str {  
      if !strings.Contains(let, string(char)) {
         return true
      }
   }
   return false
}