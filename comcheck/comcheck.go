package main 
import "os"
import "fmt"
import "string"

func main(){


	for i := 0; i < len(os.Args); i++ {

		if Strings.Contains(os.Args[i],"01") || Strings.Contains(os.Args[i],"galaxy") || Strings.Contains(os.Args[i],"galaxy 01")  {
			fmt.Println("Alert!!!")
		}
		
	}

}