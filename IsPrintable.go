package piscine

import "strings"

const lettre ="ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func IsPrintable(str string) bool {
   for _, char := range str {  
      if !strings.Contains(lettre, string(char)) {
         return false
      }
   }
   return true
}