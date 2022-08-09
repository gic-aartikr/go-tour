package main 

import ("fmt"
"math")

func pow(x , m , lim float64) float64{
	if v := math.Pow(x, m); v < lim{

		return v
	}
	return lim
}

func main(){
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}
