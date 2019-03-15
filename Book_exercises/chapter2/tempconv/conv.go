package tempconv

// CToF converts given temperature in Celsius to temperature in Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK converts given temperature in Celsius to temperature in Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// FToC converts given temperature in Fahrenheit to temperature in Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK converts given temperature in Fahrenheit to temperature in Kelvin
func FToK(f Fahrenheit) Kelvin { return Kelvin((f + 459.67) * 5 / 9) }

// KToC converts given temperature in Kelvin to temperature in Celsius
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// KToF converts given temperature in Kelvin to temperature in Fahrenheit
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(k*9/5 - 459.67) }
