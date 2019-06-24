package main 
import "os"
import "fmt"
import "strings"

func main(){
	 if strings.Contains(strings.Fields(os.Args),"galaxy01"){
			fmt.Printf("%v","")
			break
		}

	for i := 0; i < len(os.Args); i++ {

		if strings.Contains(os.Args[i],"01") || strings.Contains(os.Args[i],"galaxy")   {
			fmt.Println("Alert!!!")
			
		break	
		}
		
	}

}
