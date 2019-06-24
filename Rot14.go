

func Rot14(str string) string {

	x:=[]byte(string)

    input := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
    output := []byte("OPQRSTUVWXYZABCDEFGHIJKLMNopqrstuvwxyzabcdefghijklmn")
    match := bytes.Index(input, []byte{x})
    if match == -1 {
        return x
    }
    
    return string(output[match])
}