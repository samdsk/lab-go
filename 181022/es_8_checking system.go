    if checkA2{
      var x,y float64

      if xB2 > xA2  {
        if xB1 > xA1{
          x = xB1 - xA2
        }else{
          x = xA1 - xA2
        }
      }else{
        if xB1 > xA1{
          x = xA2 - xB1
        }else{
          x = xA2 - xA1
        }
      }

      if yB2 > yA2  {
        if yB1 > yA1{
          y = yB2 - yA1
        }else{
          y = yB2 - yB1
        }
      }else{
        if yB1 > yA1{
          y = yA2 - yA1
        }else{
          y = yA2 - yB1
        }
      }
	_ = x
	_ = y
    }
