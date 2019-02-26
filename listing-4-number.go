package main

import "fmt"

func main() {
	var (
		a  int
		b        = 10
		c  int64 = 20
		o        = 0777
		h        = 0xAbCd
		h2       = 0XaBcd

		e float64
		j float64 = 56.87
		f         = 23.45
		g         = 1.2e6
	)

	fmt.Println("Integer", a, b, c, o, h, h2)
	fmt.Printf("%d %d %d %o %x %X\n", a, b, c, o, h, h2)

	fmt.Println("Float", e, j, f, g)
	fmt.Printf("%f %.2f %f %f\n", e, j, f, g)
}
