//Learning URL https://books.studygolang.com/gopl-zh/
package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

type celsiusFlag struct {Celsius}

func (f *celsiusFlag) Set(s string) error  {
	var u string
	var value float64
	fmt.Sscanf(s, "%f%s",&value, &u)
	switch u {
	case "C","°C":
		f.Celsius = Celsius(value)
		return nil
	case "F","°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
