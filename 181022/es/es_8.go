/*
Scrivere un programma in Go che, ricevuti in input 8 float
rappresentanti le coordinate di 4 punti nel piano
(ogni coppia di punti identifica i vertici opposti di un rettangolo),
restituisca in output l'area della sovrapposizione tra i due rettangoli.
*/

package main

import (
  "fmt"
  "math"
)

func main() {
  //2 Punti dei diagonali A1 B1 (x,y) del primo rettangolo e A2 B2 (x,y) sono del secondo rettangolo
  var xA1,xB1,xA2,xB2,yA1,yB1,yA2,yB2 float64

  //Acquisisco i dati dall'utente
  fmt.Print("Inserisci le coordinate del punto A1:")
  fmt.Scan(&xA1,&yA1)

  fmt.Print("\nInserisci le coordinate del punto B1:")
  fmt.Scan(&xB1,&yB1)

  fmt.Print("\nInserisci le coordinate del punto A2:")
  fmt.Scan(&xA2,&yA2)

  fmt.Print("\nInserisci le coordinate del punto B2:")
  fmt.Scan(&xB2,&yB2)

  //Calcolo Area dei rettangoli 1 e 2
  area1 := math.Abs(xA1-xB1) * math.Abs(yA1 - yB1)
  area2 := math.Abs(xA2-xB2) * math.Abs(yA2 - yB2)

  //Confronto le aree
  var min float64
  if area1 < area2 {
    fmt.Println("Area 2 > Area 1.")
    min = area1
  }else if area1 > area2 {
    fmt.Println("Area 1 > Area 2.")
    min = area2
  }else{
    fmt.Println("Area 2 = Area 1.")
    min = area1
  }

  //Controllo e assegno ad un variabile bool due punti A2 B2 se sono interni o esterni al primo rettangolo
  var checkA2, checkB2 bool
  checkA2 = ((xA2 >= xA1) && (xA2 <= xB1) || (xA2 <= xA1) && (xA2 >= xB1)) && ((yA2 >= yA1) && (yA2 <= yB1) || (yA2 <= yA1) && (yA2 >= yB1))
  checkB2 = ((xB2 >= xA1) && (xB2 <= xB1) || (xB2 <= xA1) && (xB2 >= xB1)) && ((yB2 >= yA1) && (yB2 <= yB1) || (yB2 <= yA1) && (yB2 >= yB1))

  //Controllo tipo di sovrapposizione dei rettangoli
  if checkA2 && checkB2 {
    //Se entrambi punti del diagonale sono interni
    fmt.Println("Sovrapposizione propria, l'area è ", min)
  }else if !checkA2 && !checkB2 {
    //Se entrambi punti del diagonale NON sono interni
    fmt.Println("Non sovrapposto, l'area è 0")
  }else{
    //Se solo un punto del diagonale è interno allora calcolo l'area della sovrapposizione
    area := (math.Min(xB1, xB2) - math.Max(xA1,xA2))*(math.Min(yB1, yB2) - math.Max(yA1,yA2))

    fmt.Println("Sovrapposizione impropria, l'area della svorapposizione è ", area)
  }


}
