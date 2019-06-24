package main 
import "os"
import "fmt"
import "strings"

func main(){
	
	for i ,val:=range (os.Args) {

		if strings.Contains(val[i],"01") || strings.Contains(val[i],"galaxy") || strings.Contains(val[i],"galaxy 01")  {
			fmt.Println("Alert!!!")
		}
		
	}

}
