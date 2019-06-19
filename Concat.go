package piscine

import "strings"


func Concat(str1 string, str2 string) string {


s := []string{str1, str2}
return strings.Join(s,"")

}