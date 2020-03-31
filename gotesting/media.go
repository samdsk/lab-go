package esami

func Media(v Voti)(float64){
  if len(v) == 0{
    return 0.0
  }
  somma := 0.0
  for _,x := range v{
    if x <= 30{
      somma += float64(x)
    }else{
      somma += 30.0
    }
  }
  return somma/float64(len(v))
}

func Medie(m Nomi2Voti)(map[string]float64){
  res := make(map[string]float64)

  for nome,voti := range m{
    res[nome] = Media(voti)
  }
  return res
}

func Trentesimi2Centodecimi(media float64) float64{
  return 110.0*media/30.0
}
