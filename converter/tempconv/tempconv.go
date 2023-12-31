package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilinngC     Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%.2f °C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.2f °F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%.2f °K", k) }
