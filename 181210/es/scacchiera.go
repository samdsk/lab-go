package main

import (
  ."fmt"
  "math"
)

type Colore bool
type Tipo byte

const (
  BIANCO = false
  NERO = true

  NULL = iota
  RE
  REGINA
  TORRE
  ALFIERE
  CAVALLO
  PEDONE
)

type (
	Pezzo struct {
		tipo   Tipo
		colore Colore
	}

	Casella struct {
		riga byte
		col  rune
	}

	Scacchiera map[Casella]Pezzo
)

//stampa nome del pezzo
func NomePezzo(p Pezzo)string{
  re := 0x2654 //valore unicode hex
  t := ""
  if p.tipo >= RE && p.tipo <= PEDONE{ // RE=1,...,PEDONE=6 type byte
    if p.colore == NERO{
      re += 6
    }
    t = string(re - 1 + int(p.tipo))
    
  }

  return t
}

// controlla se le coordinate sono accettabili
func check(riga byte, colonna rune)bool{
  return riga >= 1 && riga <= 8 && colonna >= 'a' && colonna <= 'h'
}

func get(s Scacchiera, r byte, c rune)(p Pezzo, found bool){
  p, found = s[Casella{r,c}]
  return
}

func put(s Scacchiera, p Pezzo, r  byte, c rune){
  if check(r,c){
    s[Casella{r,c}] = p
  }else{
    Println("Error: PUT - Wrong Coordinates")
  }

}

func remove(s Scacchiera, r byte,c rune){
  if check(r,c){
    delete(s, Casella{r,c})
  }else{
    Println("Error: REMOVE - Wrong Coordinates")
  }

}


func Move(s Scacchiera,turno Colore, r1, r2 byte, c1, c2 rune) (src, dest Pezzo, ok bool) {
  if src, dest, ok = validMove(s,turno,r1,r2,c1,c2); ok{
    put(s,src,r2,c2)
    remove(s,r1,c1)
  }
  return
}


func validMove(s Scacchiera, turno Colore, r1, r2 byte, c1, c2 rune) (src, dest Pezzo, ok bool){
  if ok = check(r1,c1) && check(r2,c2); ok{
    src, ok = get(s,r1,c1)
    if ok != false{
      if ok = ok && src.colore == turno; ok{
        dest, found := get(s,r2,c2)
        if found != false{
          if ok = !(found && src.colore == dest.colore); ok{
            switch src.tipo{
              case ALFIERE:
                ok = checkAlfiere(s, r1, r2, c1, c2)
              case TORRE:
                ok = checkTorre(s, r1, r2, c1, c2)
              case REGINA:
                ok = checkAlfiere(s, r1, r2, c1, c2) || checkTorre(s, r1, r2, c1, c2)
              case CAVALLO:
                ok = checkCavallo(src.colore, r1, r2, c1, c2)
              case PEDONE:
                ok = checkPedone(src.colore, found, r1, r2, c1, c2)
              case RE:
                ok = checkRe(r1, r2, c1, c2)
            }
          }
        }
      }
    }
  }
  return
}

func distanzaRige(r1,r2 byte)int{
  temp := int(math.Sqrt(float64(r1-r2)))
  return temp
}

func distanzaColonne(c1,c2 rune)int{
  temp := int(math.Sqrt(float64(c1-c2)))
  return temp
}

func direzione(ok bool)int{
  if ok {
    return 1
  }
  return -1
}

func checkPath(s Scacchiera,r1 byte,c1 rune, dRige,dColonne,passi int) (ok bool){
  for ; passi > 1; passi--{
    c1 += rune(dColonne)
    r1 += byte(dRige)

    if _, found := get(s, r1,c1); found{
      return false
    }
  }
  return true
}

func checkAlfiere(s Scacchiera, r1,r2 byte,c1,c2 rune) (ok bool){
  passi := distanzaRige(r1,r2)
  if ok = passi == distanzaColonne(c1,c2); ok{
    dRige,dColonne := direzione(r1 < r2), direzione(c1 < c2)
    ok = checkPath(s,r1,c1,dRige,dColonne,passi)
  }
  return
}

func checkTorre(s Scacchiera, r1,r2 byte,c1,c2 rune)(ok bool){
  dRige,dColonne :=  distanzaRige(r1,r2),distanzaColonne(c1,c2)

  if ok = dRige == 0 || dColonne == 0; ok{
    passi := dRige + dColonne
    if dRige == 0{
      dColonne = direzione(c1 < c2)
    } else {
      dRige = direzione(r1 < r2)
    }
    ok = checkPath(s, r1,c1,dRige,dColonne,passi)
  }
  return
}


func checkRe(r1,r2 byte,c1,c2 rune)(ok bool){
  return distanzaRige(r1,r2) == 1 || distanzaColonne(c1,c2) == 1
}

func checkCavallo(color Colore, r1,r2 byte, c1,c2 rune)(ok bool){
  if !prima[color]{
    dRige,dColonne :=  distanzaRige(r1,r2),distanzaColonne(c1,c2)
    ok = (dRige == 1 && dColonne == 2) || (dRige == 2 && dColonne == 1 )
  }
  return
}

func checkPedone(colore Colore, piena bool, r1,r2 byte, c1,c2 rune)(bool){
  dColonne :=  distanzaColonne(c1,c2)
  return (dColonne == 1 && piena || dColonne == 0 && !piena) &&
		(colore == BIANCO && r2 == r1+1 || colore == NERO && r2 == r1-1)
}
