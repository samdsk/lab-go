/* hello_world.go
package main
import "fmt"
func main() {
	fmt.Println(" Hello World !")
}*/

//esegue le quattro operazioni su due numeri float64 in ingresso

package main

import "fmt"

func main() {
	var num1, num2 float64
	fmt.Println("inserisci due numeri (float)")
	fmt.Scan(&num1, &num2) //Input 
	//fmt.Scanf("%f,%f\n", &num1, &num2)	//per usare "," come separatore
	fmt.Println("num1 =", num1) //Output
	fmt.Println("num2 =", num2) //Output
		
	fmt.Println("somma =", num1+num2) //Output
	fmt.Println("differenza =", num1-num2)
	fmt.Println("prodotto =", num1*num2)
	fmt.Println("quoziente =", num1/num2)
}
