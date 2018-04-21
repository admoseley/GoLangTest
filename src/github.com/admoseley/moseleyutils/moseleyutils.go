package moseleyutils

func Moseley_add(x int, y int) int {
	var g int = x + y
	return int(g)
}

func Moseley_subtract(x int, y int) int {
	var g int
	if x > y {
		g = x - y
		return int(g)
	} else {

		g = y - x
		return int(g)
	}
}
