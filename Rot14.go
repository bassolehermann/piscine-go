package piscine

import"unicode"

func Rot14(str string) string {
var result []rune
         rot5map := map[rune]rune{'0': '0', '1': '1', '2': '2', '3': '3', '4': '4', '5': '5', '6': '6', '7': '7', '8': '8', '9': '9'}

         for _, i := range str {
                 switch {
                 case !unicode.IsLetter(i) && !unicode.IsNumber(i):
                         result = append(result, i)
                 case i >= 'A' && i <= 'Z':
                         result = append(result, 'A'+(i-'A'+14)%26)
                 case i >= 'a' && i <= 'z':
                         result = append(result, 'a'+(i-'a'+14)%26)
                 case i >= '0' && i <= '9':
                         result = append(result, rot5map[i])
                 case unicode.IsSpace(i):
                         result = append(result, ' ')
                 }
         }
         return string(result[:])
}