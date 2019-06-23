package piscine

import "strings"

const alpha = "abcdefghijklmnopqrstuvwxyz0123456789"

func IsAlpha(str string) bool {
   for _, char := range str {  
      if !strings.Contains(alpha, strings.ToLower(string(char))) {
         return false
      }
   }
   return true
}