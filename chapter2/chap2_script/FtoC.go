
package main

import "fmt"

func main(){
	const freezingF ,boilingF = 32 ,312 
	fmt.Printf("%g °F = %g °C\n",freezingF , fToC(freezingF))  //"32°F = 0 °C"
	fmt.Printf("%g °F = %g °C\n",boilingF ,fToC(boilingF))   //"212 °F = 100 °C"
}
	
func fToC(f float64){
	return(f - 32)*5/9
}