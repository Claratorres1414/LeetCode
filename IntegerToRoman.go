package main

import "fmt"

func intToRoman(num int) string {
	if num == 1 {
		return "I"
	} else if num == 3999 {
		return "MMMCMXCIX"
	}
	m := num / 1000
	num -= m * 1000
	c := num / 100
	num -= c * 100
	d := num / 10
	num -= d * 10
	u := num

	s := ""

	if m != 0 {
		s += convM(m)
	}
	if c != 0 {
		s += convC(c)
	}
	if d != 0 {
		s += convD(d)
	}
	if u != 0 {
		s += convU(u)
	}

	return s
}

func convM(num int) string {
	switch num {
	case 1:
		return "M"
	case 2:
		return "MM"
	default:
		return "MMM"
	}
}

func convC(num int) string {
	switch num {
	case 1:
		return "C"
	case 2:
		return "CC"
	case 3:
		return "CCC"
	case 4:
		return "CD"
	case 5:
		return "D"
	case 6:
		return "DC"
	case 7:
		return "DCC"
	case 8:
		return "DCCC"
	default:
		return "CM"
	}
}

func convD(num int) string {
	switch num {
	case 1:
		return "X"
	case 2:
		return "XX"
	case 3:
		return "XXX"
	case 4:
		return "XL"
	case 5:
		return "L"
	case 6:
		return "LX"
	case 7:
		return "LXX"
	case 8:
		return "LXXX"
	default:
		return "XC"
	}
}

func convU(num int) string {
	switch num {
	case 1:
		return "I"
	case 2:
		return "II"
	case 3:
		return "III"
	case 4:
		return "IV"
	case 5:
		return "V"
	case 6:
		return "VI"
	case 7:
		return "VII"
	case 8:
		return "VIII"
	default:
		return "IX"
	}
}

func main() {
	fmt.Println(intToRoman(1994))
}
