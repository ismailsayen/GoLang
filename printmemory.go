package piscine

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	slice := []string{}
	for _, ele := range arr {
		slice = append(slice, toHex(int(ele)))
	}
	for index, elem := range slice {
		if index%4 == 0 && index > 0 {
			z01.PrintRune('\n')
		} else if index%4 != 0 {
			z01.PrintRune(' ')
		}
		for _, d := range elem {
			z01.PrintRune(d)
		}

	}
	z01.PrintRune('\n')
}

func toHex(nb int) string {
	if nb == 0 {
		return "00"
	}
	res := ""
	bas := "0123456789abcdef"
	for nb > 0 {
		res = string(bas[nb%16]) + res
		nb /= 16
	}
	return res
}
