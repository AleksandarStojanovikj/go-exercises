package lengthConv

// MToF converts M to F
func MToF(m Meter) Feet { return Feet(m / 0.3048) }

// FToM converts F to M
func FToM(f Feet) Meter { return Meter(f * 0.3048) }

