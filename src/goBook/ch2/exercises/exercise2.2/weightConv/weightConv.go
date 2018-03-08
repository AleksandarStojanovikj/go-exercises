// package weightConv converts Pound and Kilogram
package weightConv

import "fmt"

type Pound float64
type Kilogram float64

func (lb Pound) String() string { return fmt.Sprintf("%g pounds", lb) }

func (kg Kilogram) String() string { return fmt.Sprintf("%g kilograms", kg) }
