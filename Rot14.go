package piscine

import"unicode"

func Rot14(str string) string {
var result []rune
         

         for _, i := range str {
                 switch {
                 case !unicode.IsLetter(i) && !unicode.IsNumber(i):
                         result = append(result, i)
                 case i >= 'A' && i <= 'Z':
                         result = append(result, 'A'+(i-'A'+14)%26)
                 case i >= 'a' && i <= 'z':
                         result = append(result, 'a'+(i-'a'+14)%26)
                 case unicode.IsSpace(i):
                         result = append(result, ' ')
                 }
         }
         return string(result[:])
}