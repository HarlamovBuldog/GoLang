// Package tempconv performs calculations
// for Celsius and for Fahrenheit
// alt + 0176 = °
package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func main() {
	//< Arifmetic operations showcase
	fmt.Printf("%g\n", BoilingC+FreezingC)
	BoilingF := CToF(BoilingC)
	fmt.Printf("%g\n", BoilingF-CToF(FreezingC))
	//Error:Mismatched types
	//fmt.Printf("%g\n", BoilingF+FreezingC)
	fmt.Println("=========")
	//>

	//< Comparing operations show case
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) //"true"
	fmt.Println(f >= 0) //"true"
	//fmt.Println(c == f)  //"Error: mismatched types"
	fmt.Println(c == Celsius(f)) //"true"
	fmt.Println("=========")
	//>

	//< Override String() method for Celsius showcase
	c1 := FToC(212.0)
	fmt.Println(c1.String())
	fmt.Printf("%v\n", c1)
	fmt.Printf("%s\n", c1)
	fmt.Println(c1)
	fmt.Printf("%g\n", c1)   //doesn't call String()
	fmt.Println(float64(c1)) //doesn't call String()
	//>
}
