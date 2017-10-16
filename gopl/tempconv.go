package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FrezzingC Celsius = 0
	BoilingC Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit((c*9/5+32))
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f-32)*5/9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g摄氏度", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g华氏度", f)
}

func main() {
	fmt.Println(CToF(AbsoluteZeroC))
	fmt.Println(CToF(FrezzingC))
	fmt.Println(CToF(BoilingC))
	fmt.Println(FToC(CToF(AbsoluteZeroC)))
	fmt.Println(FToC(CToF(FrezzingC)))
	fmt.Println(FToC(CToF(BoilingC)))
}