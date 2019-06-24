package piscine

func Rot14(str string) string {

	x:=[]byte(str)

    input := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
    output := []byte("OPQRSTUVWXYZABCDEFGHIJKLMNopqrstuvwxyzabcdefghijklmn")
    match := bytes.Index(input, x)
    if match == -1 {
        return x
    }
    
    return string(output[match])
}