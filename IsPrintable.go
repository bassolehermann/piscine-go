package piscine

import "strings"

const lettre ="abcdefghijklmnopqrstuvwxyz"

func IsPrintable(str string) bool {
   for _, char := range str {  
      if !strings.Contains(lettre, string(char)) {
         return true
      }
   }
   return false
}