package piscine

func Compact(ptr *[]string, length int) int {

var r []string
    for _, str := range *ptr{
        if str != " " {
            r = append(r, str)
        }
    }
    return len(r)

}


