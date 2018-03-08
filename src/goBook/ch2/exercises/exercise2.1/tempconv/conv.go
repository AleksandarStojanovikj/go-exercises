package tempconv

// CToF converts C to F
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK converts C to K
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// FToC converts F to C
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK converts F to K
func FToK(f Fahrenheit) Kelvin { return Kelvin((f + 459.67) / 1.8) }

// KToC converts K to C
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// KToF converts K to F
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(1.8*k - 459.67) }
