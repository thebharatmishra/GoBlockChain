package main

import "fmt"

func main(){
	var intNum int ;
	fmt.Println(intNum)
	var intNum8 int8 ;
	fmt.Println(intNum8)
	var intNum16 int16=32767 ;
	fmt.Println(intNum16)
	var intNum32 int32 ;
	fmt.Println(intNum32)
	var intNum64 int64 ;
	fmt.Println(intNum64)

//Unsigned int only stores positive integer
var intUnNum uint32 = 255;
intNum16 = intNum16+1;
fmt.Println(intUnNum)

var floatNum float64 =12345678.9
	fmt.Println(floatNum)
}