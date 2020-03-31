/*
Problema Scrivere un programma Go che converta in Fahrenheit
una temperatura espressa in gradi centigradi.
*/

package main
import "fmt"

func main() {
	var celsius, farenheit float64
	fmt.Println("Inserire temperatura in gradi celsius.")
	fmt.Scan(& celsius)
	farenheit=(celsius*9/5) + 32
	fmt.Println("Farenheit:", farenheit)
	
}

package main
import "fmt"

func main() {
	var temp	float64  // no variabili non PARLANTI

	fmt.Println("Inserire la temperatura in gradi:")
	fmt.Scan(& temp)
	fmt.Println("La temperatura in F è", temp*9/5+32)
}

package main //Con scelta

import "fmt"

func main() {
        var tempC, tempF float32
        var convert bool
        fmt.Println(" Vuoi convertire in Celsius(0) o in Fahrenheit(1)? ")
        fmt.Scan(&convert)
        if !convert {
                fmt.Println(" Inserisci i gradi Fahrenheit ")
                fmt.Scan(&tempF)
                fmt.Println(tempF, " gradi F corrispondono a ", (tempF-32)*5/9)
                }
        if convert {
                fmt.Println(" Inserisci i gradi Celsius ") 
                fmt.Scan(&tempC)
                fmt.Println(tempC, " gradi C corrispondono a ", (tempC*9/5)+32)
        }
        }


