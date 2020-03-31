/*
Scrivere un programma Go che legga sei interi
(coordinate dei punti A e B diagonali di un
rettangolo e coordinate di un punto P)
e che calcola se il punto P è fuori
o dentro il rettangolo AB


Esempio:

A (1,1)
B (2,2)

P (5,5)

è fuori

etc.

*/

package main

import (
    "fmt"
)

func main() {
    var x1,x2,y1,y2,xp,yp int

    fmt.Print("Inserisci coordinate del punto A: ")
    fmt.Scan(&x1,&y1)

    fmt.Print("\nInserisci coordinate del punto B: ")
    fmt.Scan(&x2,&y2)

    fmt.Print("\nInserisci coordinate del punto P: ")
    fmt.Scan(&xp,&yp)


    if ((xp >= x1) && (xp <= x2) || (xp <= x1) && (xp >= x2)) && ((yp >= y1) && (yp <= y2) || (yp <= y1) && (yp >= y2)) {
      fmt.Println("Il punto appartiene al rettangolo")
    }else{
      fmt.Println("Il punto non appartiene al rettangolo")
    }



}
