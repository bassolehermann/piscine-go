package main


import "fmt"
import "os"


func main (){

	for i:= len(os.Args) ; i<0;i--{

		fmt.Printf("%v",os.Args[i])
		fmt.Printf("\n")

	}


}