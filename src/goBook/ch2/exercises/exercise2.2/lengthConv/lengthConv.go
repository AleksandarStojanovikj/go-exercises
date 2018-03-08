// package lengthConv converts Feet and Meters
package lengthConv

import "fmt"

type Feet float64
type Meter float64

func (f Feet) String() string { return fmt.Sprintf("%g feet", f) }

func (m Meter) String() string { return fmt.Sprintf("%g meters", m) }
