package piscine

import "strings"

const let ="abcdefghijklmnopqrstuvwxyz"

func IsLower(str string) bool {
   for _, char := range str {  
      if !strings.Contains(let, string(char)) {
         return false
      }
   }
   return true
}