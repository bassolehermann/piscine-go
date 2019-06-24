package piscine

import"bytes"

func Rot14(str string) string {
var retour string
	x:=[]byte(str)

    input := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
    output := []byte("OPQRSTUVWXYZABCDEFGHIJKLMNopqrstuvwxyzabcdefghijklmn")
    match := bytes.Index(input, x)
    if match == -1 {
        return x
    }
    retour=string(output[match])
    return retour
}