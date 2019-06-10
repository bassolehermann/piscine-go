package piscine
import "fmt"
func PrintComb() {

var  n1, n2, n3;
  n1 = '0';
  n2 = '0';
  n3 = '0';
 
  while(n1 <= '9') {
    while(n2 <= '9') {
      while(n3 <= '9') {
        if(n1 <= n2 && n2 <= n3) {
        write(0, &n1, 1);
        write(0, &n2, 1);
        write(0, &n3, 1);
    write(0, "\n", 1);
    }
        n3 = n3 + 1;
      }
      n3 = '0';
      n2 = n2 + 1;
    }
    n2 = '0';
    n1 = n1 + 1;
  }

}