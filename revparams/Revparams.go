package main


import "fmt"
import "os"


func main (){

	for i:=len(os.Args) ; i>=1;i--{

		fmt.Printf("%v",os.Args)
		fmt.Printf("\n")

	}
	


}