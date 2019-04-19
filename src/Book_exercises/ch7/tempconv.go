// alt + 0176 = °
// Exercise 7.6 realization from book page 220
package main

import (
	"Book_exercises/chapter2/tempconv"
	"flag"
	"fmt"
)

// *celciusFlag corresponds to flag.Value interface
type celsiusFlag struct{ tempconv.Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // No need to check errors
	switch unit {
	case "C", "°C":
		f.Celsius = tempconv.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
		return nil
	case "K":
		f.Celsius = tempconv.KToC(tempconv.Kelvin(value))
		return nil
	}
	return fmt.Errorf("wrong temperature %q", s)
}

// CelsiusFlag determines flag Celsius with specified name,
// default value and usage instruction in string.
// Also returns address of flag variable.
// Flag argument should contain numeric value and measure unit ("100C")
func CelsiusFlag(name string, value tempconv.Celsius, usage string) *tempconv.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
