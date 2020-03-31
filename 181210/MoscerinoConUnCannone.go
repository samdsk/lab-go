package main

import (
	"fmt"
	"sort"
)

// trovare min e max di una slice
func main() {
	s := []float64{0, -2.3, -14, 4, 10 , -3, 7}
	min, max := minmax(s)
	fmt.Println(min, max)
}

//trova min e max in una slice usando la ricorsione
func minmax(slice []float64) (min, max float64) {
	if ( len(slice) == 1 ) {
		return slice[0], slice[0]
	}
	mmin, mmax := minmax(slice[1:])
	slc := []float64{slice[0], mmin, mmax}
	sort.Float64s(slc)//Ma questo non mette già in ordine tutta la slice? 
	//Cioè intendevo la funzione. Non serve fare tutto questo
	
	
	/*
	if slice[0] < mmin {
		mmin = slice[0]
	}
	
	if slice[0] > mmax {
		mmax = slice[0]
	}
	return mmin, mmax
	*/
	
	/*
	return math.Min(slice[0], mmin), math.Max(slice[0], mmax)
	*/
	return slc[0], slc[2]
}





