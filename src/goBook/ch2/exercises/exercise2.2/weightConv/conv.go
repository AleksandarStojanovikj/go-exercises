package weightConv

// LbToKg converts Pound to Kilgoram
func LbToKG(lb Pound) Kilogram { return Kilogram(lb * 0.45359237) }

// KgToLb converts Kilogram to Pound
func KgToLb(kg Kilogram) Pound { return Pound(kg / 0.45359237) }
